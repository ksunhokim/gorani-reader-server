package middleware

type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "sunho/gorani-reader context value " + k.name
}
