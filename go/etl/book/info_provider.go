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
	MaxReviewsResult = 10
	ReviewProviders  []ReviewProvider
	GenreProviders   []GenreProvider
)

func Rating(isbn string) (providers []string, ratings []float32, err error) {
	for _, p := range ReviewProviders {
		r, err := p.Rating(isbn)
		if err != nil {
			return nil, nil, err
		}

		ratings = append(ratings, r)
		providers = append(providers, p.Name())
	}
	return
}

func Reviews(isbn string, maxresult int) (providers []string, reviews []string, err error) {
	for _, p := range ReviewProviders {
		rarr, err := p.Reviews(isbn, maxresult)
		if err != nil {
			return nil, nil, err
		}

		reviews = append(reviews, rarr...)
		for range rarr {
			providers = append(providers, p.Name())
		}
	}
	return
}

func Genre(isbn string) (providers []string, genres []string, err error) {
	for _, p := range GenreProviders {
		garr, err := p.Genre(isbn)
		if err != nil {
			return nil, nil, err
		}

		genres = append(genres, garr...)
		for range garr {
			providers = append(providers, p.Name())
		}
	}
	return
}

func (b *Book) fetchInfos() error {
	_, genres, err := Genre(b.Isbn)
	if err != nil {
		return err
	}
	b.Genre = genres

	providers, reviews, err := Reviews(b.Isbn, MaxReviewsResult)
	if err != nil {
		return err
	}
	b.Reviews = make([]BookReview, len(reviews))
	for i := range reviews {
		b.Reviews[i] = BookReview{
			Provider: providers[i],
			Comment:  reviews[i],
		}
	}

	providers, ratings, err := Rating(b.Isbn)
	if err != nil {
		return err
	}
	b.Ratings = make([]BookRating, len(providers))
	for i := range ratings {
		b.Ratings[i] = BookRating{
			Provider: providers[i],
			Rating:   ratings[i],
		}
	}

	return nil
}
