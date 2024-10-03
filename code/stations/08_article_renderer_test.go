package stations

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleRendererFixture(t *testing.T) {
	gunit.Run(new(ArticleRendererFixture), t)
}

type ArticleRendererFixture struct {
	StationFixture
}

func (this *ArticleRendererFixture) Setup() {
	this.station = NewArticleRenderer(fmt.Sprintf("%s, %s, %s, %s, %s, %s, %s, %s",
		titlePlaceholder,
		bodyPlaceholder,
		slugPlaceholder,
		titlePlaceholder,
		datePlaceholder,
		titlePlaceholder,
		datePlaceholder,
		slugPlaceholder),
	)
}

func (this *ArticleRendererFixture) TestRendering() {
	title := "howdy"
	body := "Cody wants a better title"
	now := time.Now()
	nowStr := now.Format("January 2, 2006")
	slug := "http"
	this.do(contracts.Article{
		Body:  body,
		Title: title,
		Date:  now,
		Slug:  slug,
	})

	this.So(this.outputs, should.HaveLength, 1)
	page, ok := this.outputs[0].(contracts.Page)
	this.So(ok, should.BeTrue)
	this.So(strings.Contains(page.Content, titlePlaceholder), should.BeFalse)
	this.So(strings.Contains(page.Content, datePlaceholder), should.BeFalse)
	this.So(strings.Contains(page.Content, slugPlaceholder), should.BeFalse)
	this.So(strings.Contains(page.Content, bodyPlaceholder), should.BeFalse)

	this.So(strings.Contains(page.Content, body), should.BeTrue)
	this.So(strings.Contains(page.Content, title), should.BeTrue)
	this.So(strings.Contains(page.Content, nowStr), should.BeTrue)
	this.So(strings.Contains(page.Content, slug), should.BeTrue)
}
