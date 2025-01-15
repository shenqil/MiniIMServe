package server

import (
	"MiniIMServe/clients"
	"MiniIMServe/listeners"
	"MiniIMServe/protocol"
	pb "MiniIMServe/protocol/protobuf"
	"errors"
	"log/slog"
	"strconv"
	"time"
)

var (
	ErrListenerIDExists       = errors.New("listener id already exists")                               // a listener with the same id already exists
	ErrConnectionClosed       = errors.New("connection not open")                                      // connection is closed
	ErrInlineClientNotEnabled = errors.New("please set Options.InlineClient=true to use this feature") // inline client is not enabled by default
	ErrOptionsUnreadable      = errors.New("unable to read options from bytes")
)

// IM服务所有配置
type Options struct {
	// 内部所有日志
	Logger *slog.Logger `yaml:"-" json:"-"`
}

// IM 服务
type Server struct {
	Listeners *listeners.Listeners // 所有监听器
	Clients   *clients.Clients     // 所有客户端
	Log       *slog.Logger         // 内部日志
	done      chan bool            // 当前服务是否已经结束
}

// 创建一个IM服务
func New(log *slog.Logger) *Server {

	s := &Server{
		Listeners: listeners.New(),
		Clients:   clients.NewClients(),
		Log:       log,
		done:      make(chan bool),
	}

	return s
}

// 添加一个监听器
func (s *Server) AddListener(l listeners.Listener) error {
	if _, ok := s.Listeners.Get(l.ID()); ok {
		return ErrListenerIDExists
	}

	nl := s.Log.With(slog.String("listener", l.ID()))
	err := l.Init(nl)
	if err != nil {
		return err
	}

	s.Listeners.Add(l)

	s.Log.Info("[server][AddListener] attached listener", "id", l.ID(), "protocol", l.Protocol(), "address", l.Address())
	return nil
}

// Serve 服务开始运行
func (s *Server) Serve() error {
	s.Log.Info("[server][Serve] im serve starting")
	defer s.Log.Info("[server][Serve] im server started")

	s.Listeners.ServeAll(s.EstablishConnection) // start listening on all listeners.

	return nil
}

// Close 关闭所有监听器
func (s *Server) Close() error {
	close(s.done)
	s.Log.Info("[server][Close] gracefully stopping server")
	s.Listeners.CloseAll()

	s.Log.Info("[server][Close] im server stopped")
	return nil
}

// EstablishConnection 客户端建立连接之后，会回调这个函数
func (s *Server) EstablishConnection(listener string, c clients.Client) error {

	var loginInfo *pb.LoginPack

	// 1.第一包数据，必定为登录数据
	err := func() error {
		message, err := c.ReadMessage()
		if err != nil {
			return err
		}
		loginInfo, err = protocol.VerifyLogin(message)
		loginResponse := &pb.ResponsePack{
			Code:    0,
			Payload: "登录成功",
		}
		if err != nil {
			loginResponse.Code = 1
			loginResponse.Payload = err.Error()
		}

		// 回复客户端登录成功或者失败
		loginResponseData, err2 := protocol.Response(strconv.Itoa(int(pb.PackType_LOGIN)), loginResponse)
		if err2 == nil {
			c.WriteMessage(loginResponseData)
		} else {
			s.Log.Error("[server][EstablishConnection] 登录消息组装失败", "rrror", err2.Error())
		}

		return err
	}()

	if err != nil {
		s.Log.Error("[server][EstablishConnection] 登录失败", "listener", listener, "err", err)
		return err
	}

	// 2.登录成功初始化数据
	nl := s.Log.With(slog.String("listener", listener), slog.String("clientId", strconv.Itoa(int(loginInfo.Uid))), slog.String("clientType", strconv.Itoa(int(loginInfo.ClientType))))
	c.Init(loginInfo.Uid, loginInfo.ClientType, nl)

	nl.Info("[server][EstablishConnection] 登录成功")

	// 保存当前客户端
	s.Clients.Add(c)
	defer s.Clients.Delete(loginInfo.Uid, loginInfo.ClientType)

	for {

		// 解析每一包数据
		message, err := c.ReadMessage()
		if err != nil {
			break
		}

		nl.Info("[server][EstablishConnection][ReadMessage]", "len", len(message))
		packData, err := protocol.Decode(message)

		if err != nil {
			nl.Error(err.Error())
			continue
		}

		// 消息分发
		switch packData.Type {

		// 1.消息转发
		case pb.PackType_MESSAGE:
			clients, ok := s.Clients.Get(packData.To)

			if ok {
				// 修改为服务端时间戳
				packData.Timestamp = uint64(time.Now().UnixNano())
				// 重新编码
				data, err := protocol.Encode(packData)

				if err == nil {
					for _, client := range clients {
						client.WriteMessage(data)
					}
				}

			}

		// 2.退出登录
		case pb.PackType_LOGOUT:
			nl.Info("[server][EstablishConnection][logout]")
			logoutResponse := &pb.ResponsePack{
				Code:    0,
				Payload: "",
			}

			logoutResponseData, err := protocol.Response(strconv.Itoa(int(pb.PackType_LOGOUT)), logoutResponse)
			if err != nil {
				nl.Error("[server][EstablishConnection][logout]", "error", err)
			} else {
				c.WriteMessage(logoutResponseData)
			}

		}

	}

	return nil
}
