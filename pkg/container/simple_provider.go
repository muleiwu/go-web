package container

// SimpleProvider 即时注册已有实例的 Provider
type SimpleProvider struct {
	name     string
	instance any
	priority int
}

// NewSimpleProvider 创建一个包装已有实例的 Provider
func NewSimpleProvider(name string, instance any) *SimpleProvider {
	return &SimpleProvider{
		name:     name,
		instance: instance,
		priority: 0,
	}
}

// NewSimpleProviderWithPriority 创建带优先级的 SimpleProvider
func NewSimpleProviderWithPriority(name string, instance any, priority int) *SimpleProvider {
	return &SimpleProvider{
		name:     name,
		instance: instance,
		priority: priority,
	}
}

func (p *SimpleProvider) Name() string  { return p.name }
func (p *SimpleProvider) Build() any    { return p.instance }
func (p *SimpleProvider) Priority() int { return p.priority }

// LazyProvider 延迟创建实例的 Provider
type LazyProvider struct {
	name     string
	factory  func() any
	priority int
}

// NewLazyProvider 创建延迟初始化的 Provider
func NewLazyProvider(name string, factory func() any) *LazyProvider {
	return &LazyProvider{
		name:    name,
		factory: factory,
	}
}

// NewLazyProviderWithPriority 创建带优先级的 LazyProvider
func NewLazyProviderWithPriority(name string, factory func() any, priority int) *LazyProvider {
	return &LazyProvider{
		name:     name,
		factory:  factory,
		priority: priority,
	}
}

func (p *LazyProvider) Name() string  { return p.name }
func (p *LazyProvider) Build() any    { return p.factory() }
func (p *LazyProvider) Priority() int { return p.priority }
