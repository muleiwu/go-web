package interfaces

type AppProvider interface {
	Assemblies() []AssemblyInterface
	Servers() []ServerInterface
}
