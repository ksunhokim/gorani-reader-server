package util

type DummyWriter struct {
}

func (d DummyWriter) Write(b []byte) (int, error) {
	return len(b), nil
}
