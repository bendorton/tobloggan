package stations

import (
	"errors"
	"fmt"
	"testing"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestMarkdownConverterFixture(t *testing.T) {
	gunit.Run(new(MarkdownConverterFixture), t)
}

type MarkdownConverterFixture struct {
	StationFixture
}

func (this *MarkdownConverterFixture) Setup() {
	this.station = NewMarkdownConverter(this)
}

func (this *MarkdownConverterFixture) TestBodyConverted() {
	article := contracts.Article{Body: "markdown body"}
	this.do(article)
	this.So(this.outputs, should.HaveLength, 1)
	this.assertOutputs(contracts.Article{Body: "converted markdown body"})
}
func (this *MarkdownConverterFixture) TestInvalidMarkdown_InvalidBody() {
	article := contracts.Article{Body: "invalid"}
	this.do(article)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, markdownConverterErr)
	}
}

func (this *MarkdownConverterFixture) Convert(content string) (string, error) {
	if content == "invalid" {
		return "", errors.New("invalid content")
	}
	return fmt.Sprintf("converted %s", content), nil
}
