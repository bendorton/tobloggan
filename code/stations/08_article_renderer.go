package stations

import (
	"strings"

	"tobloggan/code/contracts"
)

type ArticleRenderer struct {
	template string
}

func NewArticleRenderer(template string) contracts.Station {
	return &ArticleRenderer{template: template}
}

const (
	titlePlaceholder = "{{Title}}"
	slugPlaceholder  = "{{Slug}}"
	datePlaceholder  = "{{Date}}"
	bodyPlaceholder  = "{{Body}}"
)

func (this *ArticleRenderer) Do(input any, output func(any)) {
	//TODO: combine the fields of the incoming contracts.Article with the article template (provided via the constructor),
	//replace: {{Title}} with contracts.Article.Title
	//         {{Slug}} with contracts.Article.Slug
	//         {{Date}} with contracts.Article.Date.Format("January 2, 2006")
	//         {{Body}} with contracts.Article.Body
	//output(contracts.Page{
	//    Path:    input.Slug,
	//    Content: replacedTemplate,
	//})
	switch article := input.(type) {
	case contracts.Article:
		this.template = strings.ReplaceAll(this.template, titlePlaceholder, article.Title)
		this.template = strings.ReplaceAll(this.template, slugPlaceholder, article.Slug)
		this.template = strings.ReplaceAll(this.template, datePlaceholder, article.Date.Format("January 2, 2006"))
		this.template = strings.ReplaceAll(this.template, bodyPlaceholder, article.Body)
		page := contracts.Page{
			Path:    article.Slug,
			Content: this.template,
		}
		output(page)
	default:
		output(input)
	}
}
