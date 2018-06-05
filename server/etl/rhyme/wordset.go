package rhyme

type WordSet map[int32]struct{}

func (a WordSet) Sub(b WordSet) WordSet {
	out := make(WordSet, len(a))
	for i := range a {
		if _, ok := b[i]; !ok {
			out[i] = struct{}{}
		}
	}
	return out
}

func (a WordSet) Intersect(b WordSet) WordSet {
	out := make(WordSet, len(a))
	for i := range a {
		if _, ok := b[i]; ok {
			out[i] = struct{}{}
		}
	}
	return out
}
