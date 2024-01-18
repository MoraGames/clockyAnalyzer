package generalSets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/sets"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ GeneralSet = new(ApprovedGeneralSet)

type ApprovedGeneralSet struct {
	Name       string
	LastSet    sets.Set
	AtLeastOne bool
	Nums       int
	Gaps       *gaps.Gaps
}

func NewGeneralSet_Approved(name string, lastSet sets.Set) *ApprovedGeneralSet {
	return &ApprovedGeneralSet{name, lastSet, false, 0, gaps.NewGaps(time.NewTime(23, 54))}
}

func (gs *ApprovedGeneralSet) GetName() string {
	return gs.Name
}

func (gs *ApprovedGeneralSet) GetLastSet() sets.Set {
	return gs.LastSet
}

func (gs *ApprovedGeneralSet) SetLastSet(set sets.Set) {
	gs.LastSet = set
	gs.Gaps = gaps.NewGaps(set.GetLastTime())
}

func (gs *ApprovedGeneralSet) GetAtLeastOne() bool {
	return gs.AtLeastOne
}

func (gs *ApprovedGeneralSet) SetAtLeastOne(b bool) {
	gs.AtLeastOne = b
}

func (gs *ApprovedGeneralSet) GetNums() int {
	return gs.Nums
}

func (gs *ApprovedGeneralSet) IncreaseNums() {
	gs.Nums++
}

func (gs *ApprovedGeneralSet) GetGaps() *gaps.Gaps {
	return gs.Gaps
}
