package stations

import (
	"errors"
	"fmt"

	"tobloggan/code/contracts"
)

type Markdown interface {
	Convert(content string) (string, error)
}

type MarkdownConverter struct {
	Markdown
}

func NewMarkdownConverter(markdown Markdown) contracts.Station {
	return &MarkdownConverter{Markdown: markdown}
}

func (this *MarkdownConverter) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.Article:
		mdBody, err := this.Convert(input.Body)
		if err != nil {
			output(fmt.Errorf("%w: %w", markdownConverterErr, err))
		} else {
			input.Body = mdBody
			output(input)
		}
	default:
		output(input)
	}
}

var (
	markdownConverterErr = errors.New("MarkdownConverter Error")
)
