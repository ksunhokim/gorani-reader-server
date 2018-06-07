package book

type (
	InfoProvider interface {
		Name() string
	}

	ReviewProvider interface {
		InfoProvider
		Rating(isbn string) (float32, error)
		Reviews(isbn string, maxresult int) ([]string, error)
	}

	GenreProvider interface {
		InfoProvider
		Genre(isbn string) ([]string, error)
	}
)

var (
	ReviewProviders []ReviewProvider
	GenreProviders  []GenreProvider
)

func Rating(isbn string) (providers []string, ratings []float32) {
	for _, p := range ReviewProviders {
		r, err := p.Rating(isbn)
		if err == nil {
			ratings = append(ratings, r)
			providers = append(providers, p.Name())
		}
	}
	return
}

func Reviews(isbn string, maxresult int) (providers []string, reviews []string) {
	for _, p := range ReviewProviders {
		rarr, err := p.Reviews(isbn, maxresult)
		if err == nil {
			reviews = append(reviews, rarr...)
			for range rarr {
				providers = append(providers, p.Name())
			}
		}
	}
	return
}

func Genre(isbn string) (providers []string, genres []string) {
	for _, p := range GenreProviders {
		garr, err := p.Genre(isbn)
		if err == nil {
			genres = append(genres, garr...)
			for range garr {
				providers = append(providers, p.Name())
			}
		}
	}
	return
}
