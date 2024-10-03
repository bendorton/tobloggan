package stations

import (
	"testing"
	"time"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestFutureRemovalFixture(t *testing.T) {
	gunit.Run(new(FutureRemovalFixture), t)
}

type FutureRemovalFixture struct {
	StationFixture
	now time.Time
}

func (this *FutureRemovalFixture) Setup() {
	this.now = time.Now()
	this.station = NewFutureRemoval(this.now)
}

func (this *FutureRemovalFixture) TestPastArticleKept() {
	article := articleAtTime(this.now.Add(-time.Second))
	this.do(article)
	this.So(this.outputs, should.HaveLength, 1)
}
func (this *FutureRemovalFixture) TestCurrentArticleKept() {
	article := articleAtTime(this.now)
	this.do(article)
	this.So(this.outputs, should.HaveLength, 1)
}
func (this *FutureRemovalFixture) TestFutureArticleDropped() {
	article := articleAtTime(this.now.Add(time.Second))
	this.do(article)
	this.So(this.outputs, should.HaveLength, 0)
}

func articleAtTime(t time.Time) contracts.Article {
	return contracts.Article{
		Date: t,
	}
}
