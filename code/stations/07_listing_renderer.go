package stations

import (
	"fmt"
	"sort"
	"strings"

	"tobloggan/code/contracts"
)

type ListingRenderer struct {
	articles []contracts.Article
	base     string
}

func NewListingRenderer(base string) *ListingRenderer {
	return &ListingRenderer{base: base}
}

func (this *ListingRenderer) Do(input any, output func(any)) {
	//    TODO: given a contracts.Article, append it to a slice and send it on
	switch input := input.(type) {
	case contracts.Article:
		this.articles = append(this.articles, input)
	}

	output(input) // we always pass along the input
}
func (this *ListingRenderer) Finalize(output func(any)) {
	//    TODO: sort the slice (by Date), generate a <li> for each article in a big string,
	sort.Sort(ByDate(this.articles)) // This should sort it by date (see below)

	//Generate the <li> guy
	var builder strings.Builder
	for _, article := range this.articles {
		builder.WriteString("\t\t\t")
		_, _ = fmt.Fprintf(&builder, `<li><a href="%s">%s</a></li>`, article.Slug, article.Title)
		builder.WriteString("\n")
	}

	pageContent := strings.Replace(this.base, "{{Listing}}", builder.String(), 1)
	output(contracts.Page{Path: "/", Content: pageContent})
}

// Custom sorting to sort byDate using the sort package (which is cool)
type ByDate []contracts.Article

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].Date.Before(a[j].Date) }
