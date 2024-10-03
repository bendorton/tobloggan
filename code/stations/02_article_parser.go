package stations

import (
	"encoding/json"
	"errors"
	"strings"

	"tobloggan/code/contracts"
)

var (
	separatorMissing = errors.New("invalid separator")
)

type ArticleParser struct{}

func NewArticleParser() contracts.Station {
	return &ArticleParser{}
}

func (this *ArticleParser) Do(input any, output func(any)) {
	switch fileContents := input.(type) {
	case contracts.SourceFile:
		var article contracts.Article
		parts := strings.Split(string(fileContents), "+++")
		if len(parts) != 2 {
			output(contracts.Errorf("%w: %w", errMalformedContent, separatorMissing))
			return
		}
		header := []byte(strings.TrimSpace(parts[0]))
		body := parts[1]
		err := json.Unmarshal(header, &article)
		if err != nil {
			output(contracts.Errorf("%w: %w", errMalformedContent, err))
			return
		}
		article.Body = body
		output(article)
	default:
		output(input)
	}
}

var errMalformedContent = errors.New("malformed content")
