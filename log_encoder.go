package refererparser

import (
	"github.com/rs/zerolog"
)

func (r *RefererResult) MarshalZerologObject(e *zerolog.Event) {
	e.
		Bool("known", r.Known).
		Str("referer", r.Referer).
		Str("medium", r.Medium).
		Str("search_parameter", r.SearchParameter).
		Str("search_term", r.SearchTerm).
		Str("url", r.URI.String())
}
