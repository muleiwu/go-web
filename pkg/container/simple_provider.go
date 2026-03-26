package container

import "reflect"

// SimpleProvider 即时注册已有实例的 Provider
type SimpleProvider struct {
	typ      reflect.Type
	instance any
	priority int
	deps     []reflect.Type
}

// NewSimpleProvider 创建一个包装已有实例的 Provider
func NewSimpleProvider(typ reflect.Type, instance any) *SimpleProvider {
	return &SimpleProvider{
		typ:      typ,
		instance: instance,
		priority: 0,
	}
}

// NewSimpleProviderWithPriority 创建带优先级的 SimpleProvider
func NewSimpleProviderWithPriority(typ reflect.Type, instance any, priority int) *SimpleProvider {
	return &SimpleProvider{
		typ:      typ,
		instance: instance,
		priority: priority,
	}
}

func (p *SimpleProvider) Type() reflect.Type        { return p.typ }
func (p *SimpleProvider) Build() any                { return p.instance }
func (p *SimpleProvider) Priority() int             { return p.priority }
func (p *SimpleProvider) DependsOn() []reflect.Type { return p.deps }

// NewSimpleProviderWithDeps 创建声明依赖的 SimpleProvider
func NewSimpleProviderWithDeps(typ reflect.Type, instance any, deps ...reflect.Type) *SimpleProvider {
	return &SimpleProvider{
		typ:      typ,
		instance: instance,
		deps:     deps,
	}
}

// LazyProvider 延迟创建实例的 Provider
type LazyProvider struct {
	typ      reflect.Type
	factory  func() any
	priority int
	deps     []reflect.Type
}

// NewLazyProvider 创建延迟初始化的 Provider
func NewLazyProvider(typ reflect.Type, factory func() any) *LazyProvider {
	return &LazyProvider{
		typ:     typ,
		factory: factory,
	}
}

// NewLazyProviderWithPriority 创建带优先级的 LazyProvider
func NewLazyProviderWithPriority(typ reflect.Type, factory func() any, priority int) *LazyProvider {
	return &LazyProvider{
		typ:      typ,
		factory:  factory,
		priority: priority,
	}
}

func (p *LazyProvider) Type() reflect.Type        { return p.typ }
func (p *LazyProvider) Build() any                { return p.factory() }
func (p *LazyProvider) Priority() int             { return p.priority }
func (p *LazyProvider) DependsOn() []reflect.Type { return p.deps }

// NewLazyProviderWithDeps 创建声明依赖的 LazyProvider
func NewLazyProviderWithDeps(typ reflect.Type, factory func() any, deps ...reflect.Type) *LazyProvider {
	return &LazyProvider{
		typ:     typ,
		factory: factory,
		deps:    deps,
	}
}
