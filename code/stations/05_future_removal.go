package stations

import (
	"time"

	"tobloggan/code/contracts"

	"github.com/mdwhatcott/pipelines"
)

type FutureRemoval struct {
	started time.Time
}

func NewFutureRemoval(started time.Time) pipelines.Station {
	return &FutureRemoval{
		started: started,
	}
}

func (this *FutureRemoval) Do(input any, output func(any)) {
	switch article := input.(type) {
	case contracts.Article:
		if this.started.Before(article.Date) {
			return
		}
		output(article)
	default:
		output(input)
	}
}
