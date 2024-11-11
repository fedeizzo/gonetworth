package loader

import "github.com/fedeizzo/gonetworth/internal/db"

type Loader interface {
	LoadAll()
}

type loader struct {
	queries db.Queries
}

func New(queries db.Queries) loader {
	return loader{
		queries: queries,
	}
}

// func (l loader) LoadAll() (accounts []accounts)
