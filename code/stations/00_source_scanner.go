package stations

import (
	"fmt"
	"io/fs"
	"strings"

	"tobloggan/code/contracts"
)

type SourceScanner struct {
	fs fs.FS
}

func NewSourceScanner(fs fs.FS) contracts.Station {
	return &SourceScanner{fs: fs}
}

func (this *SourceScanner) Do(input any, output func(any)) {
	switch input := input.(type) {
	case contracts.SourceDirectory:
		err := fs.WalkDir(this.fs, string(input), func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(d.Name(), ".md") {
				output(contracts.SourceFilePath(path))
			}
			return nil
		})
		if err != nil {
			output(fmt.Errorf("error reading from source directory [%s]: %w", input, err))
		}
	default:
		output(input)
	}
}
