package generalSets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/sets"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ GeneralSet = new(UnderReviewGeneralSet)

type UnderReviewGeneralSet struct {
	Name       string
	LastSet    sets.Set
	AtLeastOne bool
	Nums       int
	Gaps       *gaps.Gaps
}

func NewGeneralSet_UnderReview(name string, lastSet sets.Set) *UnderReviewGeneralSet {
	return &UnderReviewGeneralSet{name, lastSet, false, 0, gaps.NewGaps(time.NewTime(23, 54))}
}

func (gs *UnderReviewGeneralSet) GetName() string {
	return gs.Name
}

func (gs *UnderReviewGeneralSet) GetLastSet() sets.Set {
	return gs.LastSet
}

func (gs *UnderReviewGeneralSet) SetLastSet(set sets.Set) {
	gs.LastSet = set
	gs.Gaps = gaps.NewGaps(set.GetLastTime())
}

func (gs *UnderReviewGeneralSet) GetAtLeastOne() bool {
	return gs.AtLeastOne
}

func (gs *UnderReviewGeneralSet) SetAtLeastOne(b bool) {
	gs.AtLeastOne = b
}

func (gs *UnderReviewGeneralSet) GetNums() int {
	return gs.Nums
}

func (gs *UnderReviewGeneralSet) IncreaseNums() {
	gs.Nums++
}

func (gs *UnderReviewGeneralSet) GetGaps() *gaps.Gaps {
	return gs.Gaps
}
