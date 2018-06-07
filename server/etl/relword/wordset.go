package relword

type wordSet map[int]struct{}

func (a wordSet) Sub(b wordSet) wordSet {
	out := make(wordSet, len(a))
	for i := range a {
		if _, ok := b[i]; !ok {
			out[i] = struct{}{}
		}
	}
	return out
}

func (a wordSet) Intersect(b wordSet) wordSet {
	out := make(wordSet, len(a))
	for i := range a {
		if _, ok := b[i]; ok {
			out[i] = struct{}{}
		}
	}
	return out
}
