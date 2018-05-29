package router

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "sunho/gorani-reader/etl context value " + k.name
}
