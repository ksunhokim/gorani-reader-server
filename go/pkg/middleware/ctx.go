package middleware

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "sunho/gorani-reader/middleware context value " + k.name
}
