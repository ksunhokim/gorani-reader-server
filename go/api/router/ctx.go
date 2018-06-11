package router

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "sunho/gorani-reader/api context value " + k.name
}
