package driver

import (
	"fmt"
	"sync"
)

// DriverFactory 驱动工厂函数，接收配置并创建驱动实例
type DriverFactory[T any] func(config any) (T, error)

// Manager 泛型驱动管理器，管理同一类型服务的多个驱动实现
type Manager[T any] struct {
	mu          sync.RWMutex
	factories   map[string]DriverFactory[T]
	defaultName string
}

// NewManager 创建新的驱动管理器
func NewManager[T any]() *Manager[T] {
	return &Manager[T]{
		factories: make(map[string]DriverFactory[T]),
	}
}

// Extend 注册驱动工厂
func (m *Manager[T]) Extend(name string, factory DriverFactory[T]) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.factories[name] = factory
}

// SetDefault 设置默认驱动名称
func (m *Manager[T]) SetDefault(name string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.defaultName = name
}

// Make 按名称创建驱动实例
func (m *Manager[T]) Make(name string, config any) (T, error) {
	m.mu.RLock()
	factory, ok := m.factories[name]
	m.mu.RUnlock()

	var zero T
	if !ok {
		return zero, fmt.Errorf("driver %q not registered", name)
	}
	return factory(config)
}

// MakeDefault 使用默认驱动名称创建实例
func (m *Manager[T]) MakeDefault(config any) (T, error) {
	m.mu.RLock()
	name := m.defaultName
	m.mu.RUnlock()

	if name == "" {
		var zero T
		return zero, fmt.Errorf("no default driver set")
	}
	return m.Make(name, config)
}

// Has 检查驱动是否已注册
func (m *Manager[T]) Has(name string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	_, ok := m.factories[name]
	return ok
}
