// websocket 监听器需要实现Listener和Client两个接口

package listeners

import (
	"context"
	"crypto/tls"
	"errors"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"log/slog"

	"github.com/gorilla/websocket"
)

const TypeWS = "ws"

var (
	// ErrInvalidMessage 消息类型不是binary
	ErrInvalidMessage = errors.New("message type not binary")
)

// WebsocketConfig  Websocket初始化配置
type WebsocketConfig struct {
	ID        string
	Address   string
	TLSConfig *tls.Config
}

// Websocket 实现 Listener 接口
type Websocket struct {
	sync.RWMutex
	id        string       // the internal id of the listener
	address   string       // the network address to bind to
	listen    *http.Server // a http server for serving websocket connections
	log       *slog.Logger // server logger
	config    *WebsocketConfig
	establish EstablishFn         // the server's establish connection handler
	upgrader  *websocket.Upgrader //  upgrade the incoming http/tcp connection to a websocket compliant connection.
	end       uint32              // ensure the close methods are only called once
}

// NewWebsocket 创建一个ws监听器
func NewWebsocket(config *WebsocketConfig) *Websocket {
	return &Websocket{
		id:      config.ID,
		address: config.Address,
		config:  config,
		upgrader: &websocket.Upgrader{
			Subprotocols: []string{},
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

// ID 返回当前监听器的ID
func (l *Websocket) ID() string {
	return l.id
}

// Address 返回当前监听器的地址
func (l *Websocket) Address() string {
	return l.address
}

// Protocol 返回当前监听器的协议
func (l *Websocket) Protocol() string {
	if l.config.TLSConfig != nil {
		return "wss"
	}

	return "ws"
}

// Init 监听器初始化
func (l *Websocket) Init(log *slog.Logger) error {
	l.log = log

	mux := http.NewServeMux()
	mux.HandleFunc("/", l.handler)
	l.listen = &http.Server{
		Addr:         l.address,
		Handler:      mux,
		TLSConfig:    l.config.TLSConfig,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	return nil
}

// handler http升级到ws的处理函数
func (l *Websocket) handler(w http.ResponseWriter, r *http.Request) {
	c, err := l.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()

	var wg sync.WaitGroup
	conn := &wsConn{listener: l.id, c: c, chMsg: make(chan []byte, 1000), wgClose: &wg, Log: l.log}

	wg.Add(1)
	go conn.LoopTask()

	err = l.establish(l.id, conn)
	if err != nil {
		conn.Log.Warn("[websocket][handler]", "error", err)
	}
}

// Serve 监听器开始服务
func (l *Websocket) Serve(establish EstablishFn) {
	var err error
	l.establish = establish

	if l.listen.TLSConfig != nil {
		err = l.listen.ListenAndServeTLS("", "")
	} else {
		err = l.listen.ListenAndServe()
	}

	// After the listener has been shutdown, no need to print the http.ErrServerClosed error.
	if err != nil && atomic.LoadUint32(&l.end) == 0 {
		l.log.Error("[websocket] failed to serve.", "error", err, "listener", l.id)
	}
}

// Close 关闭监听器
func (l *Websocket) Close() {
	l.Lock()
	defer l.Unlock()

	if atomic.CompareAndSwapUint32(&l.end, 0, 1) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = l.listen.Shutdown(ctx)
	}

}

// wsConn 实现 Client 接口
type wsConn struct {
	listener   string
	id         uint32
	clientType uint32
	c          *websocket.Conn
	Log        *slog.Logger
	chMsg      chan []byte     // 消息写入列表
	wgClose    *sync.WaitGroup // 关闭锁
}

// Init 初始化
func (ws *wsConn) Init(clientId uint32, clientType uint32, log *slog.Logger) error {

	ws.id = clientId
	ws.clientType = clientType
	ws.Log = log

	return nil
}

// ListenerID 返回监听器ID
func (ws *wsConn) ListenerID() string {
	return ws.listener
}

// ID 返回客户端ID
func (ws *wsConn) ID() uint32 {
	return ws.id
}

// ID 返回客户端类型
func (ws *wsConn) ClientType() uint32 {
	return ws.clientType
}

// ReadMessage 读取一条消息
func (ws *wsConn) ReadMessage() ([]byte, error) {
	op, r, err := ws.c.ReadMessage()
	if err != nil {
		return nil, err
	}

	if op != websocket.BinaryMessage {
		err = ErrInvalidMessage
		return nil, err
	}

	return r, err
}

// WriteMessage 写入一条消息
func (ws *wsConn) WriteMessage(data []byte) error {

	select {
	case ws.chMsg <- data:
		return nil
	default:
		ws.Log.Error("[websocket] The message failed to be sent")
		return errors.New("channel close")
	}

}

// Close 客户端关闭
func (ws *wsConn) Close() error {

	ws.Log.Info("[websocket][Close]")

	// 安全的关闭
	ws.wgClose.Wait()
	close(ws.chMsg)

	return ws.c.Close()
}

// LoopTask 循环任务
func (ws *wsConn) LoopTask() {
	defer ws.wgClose.Done()

	ws.Log.Info("[websocket][LoopTask]")

	for {
		val, ok := <-ws.chMsg
		if !ok {
			ws.Log.Info("[websocket]LoopTask] Channel closed, exiting consumer.")
			break
		}

		ws.Log.Info("[websocket][LoopTask] Send message", "Len", strconv.Itoa(len(val)))

		err := ws.c.WriteMessage(websocket.BinaryMessage, val)
		if err != nil {
			ws.Log.Error("[websocket][LoopTask] Send message", "err", err.Error())
			break
		}
	}
}
