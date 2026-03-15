package container

// Provider 所有可装配服务必须实现此接口
type Provider interface {
	Name() string  // 服务唯一标识
	Build() any    // 构建服务实例
	Priority() int // 优先级（同类型多个实现时生效）
}

// Initializable 服务启动后的初始化钩子
type Initializable interface {
	Init() error
}

// Destroyable 服务销毁钩子
type Destroyable interface {
	Destroy() error
}
