package stations

import (
	"testing"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestDraftRemovalFixture(t *testing.T) {
	gunit.Run(new(DraftRemovalFixture), t)
}

type DraftRemovalFixture struct {
	StationFixture
}

func (this *DraftRemovalFixture) Setup() {
	this.station = NewDraftRemoval()
}

func (this *DraftRemovalFixture) TestDraftDropped() {
	article := contracts.Article{Draft: true}
	this.do(article)
	this.So(this.outputs, should.HaveLength, 0)
}
func (this *DraftRemovalFixture) TestNonDraftRetained() {
	article := contracts.Article{Draft: false}
	this.do(article)
	this.So(this.outputs, should.HaveLength, 1)
	this.assertOutputs(article)
}
