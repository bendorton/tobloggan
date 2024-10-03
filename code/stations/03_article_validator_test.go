package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestArticleValidatorFixture(t *testing.T) {
	gunit.Run(new(ArticleValidatorFixture), t)
}

type ArticleValidatorFixture struct {
	StationFixture
	article contracts.Article
}

func (this *ArticleValidatorFixture) Setup() {
	this.station = NewArticleValidator()
	this.article = contracts.Article{
		Draft: false,
		Slug:  "a/valid/url/here",
		Title: "This is a Valid Title",
		Date:  time.Time{}, // Don't care
		Body:  "Body",      // Don't care
	}
}

func (this *ArticleValidatorFixture) TestValidArticle() {
	this.do(this.article)
	this.assertOutputs(this.article)
}

func (this *ArticleValidatorFixture) TestInvalidSlugs() {
	this.article.Slug = "an/INVALID/<URL>/..."
	this.do(this.article)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errInvalidSlug)
	}
}

func (this *ArticleValidatorFixture) TestInvalidTitles() {
	this.article.Title = ""
	this.do(this.article)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, errInvalidTitle)
	}
}

func (this *ArticleValidatorFixture) TestSlugsMustBeUnique() {
	this.do(this.article)
	this.do(this.article) // Process article with same slug twice
	if this.So(this.outputs, should.HaveLength, 2) {
		this.So(this.outputs[1], should.Wrap, errDuplicateSlug)
	}
}
