package stations

import (
	"os"
	"testing"
	"testing/fstest"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSourceReaderFixture(t *testing.T) {
	gunit.Run(new(SourceReaderFixture), t)
}

type SourceReaderFixture struct {
	StationFixture
	files fstest.MapFS
}

func (this *SourceReaderFixture) Setup() {
	this.files = fstest.MapFS{}
	this.station = NewSourceReader(this.files)
}

func (this *SourceReaderFixture) TestSourceFileContentReadFromDiskAndEmitted() {
	this.files["src/article.md"] = &fstest.MapFile{Data: []byte("article content")}
	this.do(contracts.SourceFilePath("src/article.md"))
	this.assertOutputs(contracts.SourceFile("article content"))
}

func (this *SourceReaderFixture) TestIOError() {
	clear(this.files)
	this.do(contracts.SourceFilePath("src/article.md"))
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, os.ErrNotExist)
	}
}
