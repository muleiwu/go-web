package interfaces

type ConfigInterface interface {
	Get(key string) (any, error)
	Set(key string, value any)
}

type InitConfig interface {
	InitConfig() map[string]any
}
