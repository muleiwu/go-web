package interfaces

import "reflect"

type AssemblyInterface interface {
	// Type 返回此 Assembly 提供的服务类型标识
	Type() reflect.Type
	// DependsOn 声明此 Assembly 依赖的其他 Assembly 类型
	DependsOn() []reflect.Type
	// Assembly 构建服务实例并返回，由框架统一注册到容器
	Assembly() (any, error)
}
