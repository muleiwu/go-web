package interfaces

type AppProvider interface {
	Assemblies(helper HelperInterface) []AssemblyInterface
	Servers(helper HelperInterface) []ServerInterface
}
