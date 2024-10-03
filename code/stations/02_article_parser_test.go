package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

//func (this *ArticleParserFixture) TestArticleMetaAndContentReadFromDiskAndEmitted() {}
//func (this *ArticleParserFixture) TestMissingDivider() {}
//func (this *ArticleParserFixture) TestMalformedMetadata() {}

const article1Content = `{
	"date": "2024-09-04T00:00:00Z",
	"slug": "/article/1",
	"title": "Article 1"
}

+++

The contents of article 1.`

func TestArticleParserFixture(t *testing.T) {
	gunit.Run(new(ArticleParserFixture), t)
}

type ArticleParserFixture struct {
	StationFixture
}

func (this *ArticleParserFixture) Setup() {
	this.station = NewArticleParser()
}

func (this *ArticleParserFixture) TestCorrectlyParsingArticle() {
	t, _ := time.Parse(time.RFC3339, "2024-09-04T00:00:00Z")
	want := contracts.Article{
		Draft: false,
		Slug:  "/article/1",
		Title: "Article 1",
		Date:  t,
		Body:  "\n\nThe contents of article 1.",
	}
	this.do(contracts.SourceFile(article1Content))
	this.assertOutputs(want)
}

func (this *ArticleParserFixture) TestBadJson() {
	badJson := "{\"hi\":}+++"
	this.do(contracts.SourceFile(badJson))
	this.So(this.outputs, should.HaveLength, 1)
	this.So(this.outputs[0], should.Wrap, errMalformedContent)
}

func (this *ArticleParserFixture) TestNoSeparator() {
	badArticle := "{} ++"
	this.do(contracts.SourceFile(badArticle))
	this.So(this.outputs, should.HaveLength, 1)
	this.So(this.outputs[0], should.Wrap, separatorMissing)
}
