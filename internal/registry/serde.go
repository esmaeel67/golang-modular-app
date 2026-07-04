package registry

type Serde interface {
	Register(v Registrable, options ...BuildOption) error
	RegisterKey(key string, b interface{}, options ...BuildOption) error
	RegisterFactory(key string, fn func() interface{}, options ...BuildOption) error
}
