package stations

import (
	"io/fs"

	"tobloggan/code/contracts"
)

type SourceReader struct {
	files fs.FS
}

func NewSourceReader(files fs.FS) contracts.Station {
	return &SourceReader{files: files}
}

func (this *SourceReader) Do(input any, output func(v any)) {
	switch input := input.(type) {
	case contracts.SourceFilePath:
		content, err := fs.ReadFile(this.files, string(input))
		if err != nil {
			output(err)
		} else {
			output(contracts.SourceFile(content))
		}
	default:
		output(input)
	}
}
