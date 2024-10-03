package stations

import (
	"errors"

	"tobloggan/code/contracts"
	"tobloggan/code/set"
)

type ArticleValidator struct {
	slugs set.Set[string]
}

func NewArticleValidator() contracts.Station {
	return &ArticleValidator{
		slugs: set.New[string](),
	}
}

func (this *ArticleValidator) Do(input any, output func(any)) {
	switch article := input.(type) {
	case contracts.Article:
		// Validate Title
		if article.Title == "" {
			output(errInvalidTitle)
			return
		}
		// Validate Slug
		for i := range article.Slug {
			if !validSlugCharacters.Contains(rune(article.Slug[i])) {
				output(errInvalidSlug)
				return

			}
		}
		// Validate Unique Slug
		if this.slugs.Contains(article.Slug) {
			output(errDuplicateSlug)
			return
		} else {
			this.slugs.Add(article.Slug)
		}
		output(article)
	default:
		output(input)
	}
}

var validSlugCharacters = set.New([]rune("abcdefghijklmnopqrstuvwxyz0123456789-/")...)

var errInvalidTitle = errors.New("invalid title")
var errInvalidSlug = errors.New("invalid slug")
var errDuplicateSlug = errors.New("duplicate slug")
