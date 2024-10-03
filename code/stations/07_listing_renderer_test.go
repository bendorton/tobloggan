package stations

import (
	"sort"
	"testing"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestListingRenderer(t *testing.T) {
	gunit.Run(new(ListingRendererFixture), t)
}

type ListingRendererFixture struct {
	StationFixture
	articles []contracts.Article
}

func (this *ListingRendererFixture) Setup() {
	this.station = NewListingRenderer("{{Listing}}") // This is so we can make sure that it gets replaced
}

func (this *ListingRendererFixture) TestArticlesWrittenToListing() {
	article1 := contracts.Article{Title: "Article1", Slug: "t1", Date: date("2024-10-01")}
	article2 := contracts.Article{Title: "Article2", Slug: "t2", Date: date("2024-02-01")}
	article3 := contracts.Article{Title: "Article3", Slug: "t3", Date: date("2024-04-01")}

	this.do(article1)
	this.do(article2)
	this.do(article3)

	this.finalize()

	page := this.outputs[3].(contracts.Page)
	content := page.Content

	this.So(page.Path, should.Equal, "/")
	this.So(content, should.ContainSubstring, `href="t2"`)
	this.So(content, should.ContainSubstring, `href="t3"`)
	this.So(content, should.ContainSubstring, `href="t1"`)
	this.So(content, should.ContainSubstring, `t1`)
	this.So(content, should.ContainSubstring, `t2`)
	this.So(content, should.ContainSubstring, `t3`)
}

func (this *ListingRendererFixture) TestArticlesAdded() {}

func (this *ListingRendererFixture) TestSortedByDate() {
	article1 := contracts.Article{Date: date("2024-10-01")}
	article2 := contracts.Article{Date: date("2024-02-01")}
	article3 := contracts.Article{Date: date("2024-01-01")}

	articles := []contracts.Article{article1, article2, article3}
	this.So(articles[0].Date, should.Equal, date("2024-10-01"))

	sort.Sort(ByDate(articles))

	this.So(articles[0].Date, should.Equal, date("2024-01-01"))
	this.So(articles[1].Date, should.Equal, date("2024-02-01"))
	this.So(articles[2].Date, should.Equal, date("2024-10-01"))
}
