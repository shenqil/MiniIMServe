package clients

import (
	"errors"
	"log/slog"
	"sync"
)

var (
	ErrMinimumKeepalive = errors.New("client keepalive is below minimum recommended value and may exhibit connection instability")
)

// Client 每个客户端都要实现的方法
type Client interface {
	Init(clientId uint32, clientType uint32, log *slog.Logger) error // 初始化
	ListenerID() string                                              // 对应监听器ID
	ID() uint32                                                      // 客户端唯一ID
	ClientType() uint32                                              // 客户端类型
	ReadMessage() ([]byte, error)                                    // 消息读取函数
	WriteMessage([]byte) error                                       // 消息写入函数
	Close() error                                                    // 客户端关闭
}

// 管理所有客户端
type Clients struct {
	internal map[uint32][]Client // 通过map保存所有客户端
	sync.RWMutex
}

// NewClients 返回客户端管理实例
func NewClients() *Clients {
	return &Clients{
		internal: make(map[uint32][]Client),
	}
}

// Add 新增一个客户端
func (cl *Clients) Add(val Client) {
	cl.Lock()
	defer cl.Unlock()
	id := val.ID()
	cl.internal[id] = append(cl.internal[id], val)
}

// GetAll 返回所有客户端
func (cl *Clients) GetAll() map[uint32][]Client {
	cl.RLock()
	defer cl.RUnlock()
	m := map[uint32][]Client{}
	for k, v := range cl.internal {

		newClients := make([]Client, len(v))
		copy(newClients, v)

		m[k] = newClients
	}
	return m
}

// Get 返回指定ID的客户端
func (cl *Clients) Get(id uint32) ([]Client, bool) {
	cl.RLock()
	defer cl.RUnlock()
	val, ok := cl.internal[id]

	newClients := make([]Client, len(val))
	copy(newClients, val)

	return newClients, ok
}

// Len 返回客户端总数量
func (cl *Clients) Len() int {
	cl.RLock()
	defer cl.RUnlock()
	val := len(cl.internal)
	return val
}

// Delete 删除指定id的客户端
func (cl *Clients) Delete(id uint32, clientType uint32) {
	cl.Lock()
	defer cl.Unlock()

	for i, client := range cl.internal[id] {
		if client.ClientType() == clientType {
			// 删除该 client
			cl.internal[id] = append(cl.internal[id][:i], cl.internal[id][i+1:]...)
			break
		}
	}

	if len(cl.internal[id]) == 0 {
		delete(cl.internal, id)
	}
}

// GetByListener 获取对应监听器的所有客户端
func (cl *Clients) GetByListener(listener string) []Client {
	cl.RLock()
	defer cl.RUnlock()
	newClients := make([]Client, 0, cl.Len())
	for _, clients := range cl.internal {
		for _, client := range clients {
			if client.ListenerID() == listener {
				newClients = append(newClients, client)
			}
		}
	}
	return newClients
}
