package stations

import (
	"io/fs"
	"os"
	"testing"
	"testing/fstest"

	"tobloggan/code/contracts"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestSourceScannerFixture(t *testing.T) {
	gunit.Run(new(SourceScannerFixture), t)
}

type SourceScannerFixture struct {
	StationFixture
	fs fstest.MapFS
}

func (this *SourceScannerFixture) Setup() {
	this.fs = make(fstest.MapFS)
	this.fs["src"] = &fstest.MapFile{Mode: fs.ModeDir}
	this.fs["src/article-1.md"] = &fstest.MapFile{Data: []byte("article 1 source")}
	this.fs["src/article-2.txt"] = &fstest.MapFile{Data: []byte("article 2 source")}
	this.fs["src/article-3.md"] = &fstest.MapFile{Data: []byte("article 3 source")}
	this.fs["src/inner"] = &fstest.MapFile{Mode: fs.ModeDir}
	this.fs["src/inner/article-4.md"] = &fstest.MapFile{Data: []byte("article 4 source")}
	this.fs["src/dir.md"] = &fstest.MapFile{Mode: fs.ModeDir} // a directory that looks like a markdown file
	this.station = NewSourceScanner(this.fs)
}

func (this *SourceScannerFixture) TestGivenASourceDirectoryThatDoesNotExist_EmitError() {
	clear(this.fs)
	this.station.Do(contracts.SourceDirectory("NOT-THERE"), this.output)
	if this.So(this.outputs, should.HaveLength, 1) {
		this.So(this.outputs[0], should.Wrap, os.ErrNotExist)
	}
}

func (this *SourceScannerFixture) TestGivenASourceDirectoryWithBlogSourceFiles_EmitAllBlogSourceFilePaths() {
	this.station.Do(contracts.SourceDirectory("src"), this.output)
	this.assertOutputs(
		contracts.SourceFilePath("src/article-1.md"),
		contracts.SourceFilePath("src/article-3.md"),
		contracts.SourceFilePath("src/inner/article-4.md"),
	)
}
