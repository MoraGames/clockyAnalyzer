package generalSets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/sets"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ GeneralSet = new(AnyGeneralSet)

type AnyGeneralSet struct {
	Name       string
	LastSet    sets.Set
	AtLeastOne bool
	Nums       int
	Gaps       *gaps.Gaps
}

func NewGeneralSet_Any(name string, lastSet sets.Set) *AnyGeneralSet {
	return &AnyGeneralSet{name, lastSet, false, 0, gaps.NewGaps(time.NewTime(00, 00))}
}

func (gs *AnyGeneralSet) GetName() string {
	return gs.Name
}

func (gs *AnyGeneralSet) GetLastSet() sets.Set {
	return gs.LastSet
}

func (gs *AnyGeneralSet) SetLastSet(set sets.Set) {
	gs.LastSet = set
	gs.Gaps = gaps.NewGaps(set.GetLastTime())
}

func (gs *AnyGeneralSet) GetAtLeastOne() bool {
	return gs.AtLeastOne
}

func (gs *AnyGeneralSet) SetAtLeastOne(b bool) {
	gs.AtLeastOne = b
}

func (gs *AnyGeneralSet) GetNums() int {
	return gs.Nums
}

func (gs *AnyGeneralSet) IncreaseNums() {
	gs.Nums++
}

func (gs *AnyGeneralSet) GetGaps() *gaps.Gaps {
	return gs.Gaps
}
