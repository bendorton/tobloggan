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
	switch input := input.(type) {
	case contracts.Article:
		var err error = nil
		// Validate Title
		if input.Title == "" {
			err = errInvalidTitle
		}
		// Validate Slug
		for i := range input.Slug {
			if !validSlugCharacters.Contains(rune(input.Slug[i])) {
				err = errInvalidSlug
				break
			}
		}
		// Validate Unique Slug (finally?)
		if this.slugs.Contains(input.Slug) {
			err = errDuplicateSlug
		} else {
			this.slugs.Add(input.Slug)
		}

		if err != nil {
			output(err)
		} else {
			output(input)
		}
	default:
		output(input)
	}
}

var validSlugCharacters = set.New([]rune("abcdefghijklmnopqrstuvwxyz0123456789-/")...)

var errInvalidTitle = errors.New("invalid title")
var errInvalidSlug = errors.New("invalid slug")
var errDuplicateSlug = errors.New("duplicate slug")
