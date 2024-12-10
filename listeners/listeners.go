package listeners

import (
	"MiniIMServe/clients"
	"sync"

	"log/slog"
)

// EstablishFn 新的客户端建立成功的回调函数
type EstablishFn func(id string, c clients.Client) error

// Listener 每个监听器都需要实现的方法
type Listener interface {
	Init(*slog.Logger) error // 初始化
	Serve(EstablishFn)       // 开始服务
	ID() string              // 返回监听器ID
	Address() string         // 监听器监听的地址
	Protocol() string        // 协议
	Close()                  // 关闭
}

// Listeners 管理所有监听器
type Listeners struct {
	internal  map[string]Listener // 包含所有监听器
	ClientsWg sync.WaitGroup      // 等待所有侦听器中的所有客户端完成的 WaitGroup。
	sync.RWMutex
}

// New 新建并返回监听管理器
func New() *Listeners {
	return &Listeners{
		internal: map[string]Listener{},
	}
}

// Add 新增一个监听器
func (l *Listeners) Add(val Listener) {
	l.Lock()
	defer l.Unlock()
	l.internal[val.ID()] = val
}

// Get 拿到指定ID的监听器
func (l *Listeners) Get(id string) (Listener, bool) {
	l.RLock()
	defer l.RUnlock()
	val, ok := l.internal[id]
	return val, ok
}

// Len 监听器总数量
func (l *Listeners) Len() int {
	l.RLock()
	defer l.RUnlock()
	return len(l.internal)
}

// Delete 删除指定监听器
func (l *Listeners) Delete(id string) {
	l.Lock()
	defer l.Unlock()
	delete(l.internal, id)
}

// Serve 指定监听器开始服务
func (l *Listeners) Serve(id string, establisher EstablishFn) {
	l.RLock()
	defer l.RUnlock()
	listener := l.internal[id]

	go func(e EstablishFn) {
		listener.Serve(e)
	}(establisher)
}

// ServeAll 所有监听器开始服务
func (l *Listeners) ServeAll(establisher EstablishFn) {
	l.RLock()
	i := 0
	ids := make([]string, len(l.internal))
	for id := range l.internal {
		ids[i] = id
		i++
	}
	l.RUnlock()

	for _, id := range ids {
		l.Serve(id, establisher)
	}
}

// Close 关闭指定监听器
func (l *Listeners) Close(id string) {
	l.RLock()
	defer l.RUnlock()
	if listener, ok := l.internal[id]; ok {
		listener.Close()
	}
}

// CloseAll 关闭所有监听器
func (l *Listeners) CloseAll() {
	l.RLock()
	i := 0
	ids := make([]string, len(l.internal))
	for id := range l.internal {
		ids[i] = id
		i++
	}
	l.RUnlock()

	for _, id := range ids {
		l.Close(id)
	}
	l.ClientsWg.Wait()
}
