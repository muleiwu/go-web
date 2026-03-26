package interfaces

type AssemblyInterface interface {
	// Name 返回此 Assembly 的唯一标识（即它注册到容器中的服务名）
	Name() string
	// DependsOn 声明此 Assembly 依赖的其他 Assembly 名称
	DependsOn() []string
	// Assembly 构建服务实例并返回，由框架统一注册到容器
	Assembly() (any, error)
}
