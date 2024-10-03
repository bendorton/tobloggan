package stations

import "tobloggan/code/contracts"

type DraftRemoval struct{}

func NewDraftRemoval() contracts.Station {
	return &DraftRemoval{}
}

func (this *DraftRemoval) Do(input any, output func(any)) {
	//TODO: given a contracts.Article, only output it if !input.Draft.
	switch input := input.(type) {
	case contracts.Article:
		if !input.Draft {
			output(input)
		}
	default:
		output(input)
	}
}
