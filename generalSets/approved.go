package generalSets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/sets"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ GeneralSet = new(ApprovedGeneralSet)

type ApprovedGeneralSet struct {
	LastSet    sets.Set
	AtLeastOne bool
	Nums       int
	Gaps       *gaps.Gaps
}

func NewGeneralSet_Approved(lastSet sets.Set) *ApprovedGeneralSet {
	return &ApprovedGeneralSet{lastSet, false, 0, gaps.NewGaps(time.NewTime(23, 54))}
}

func (s *ApprovedGeneralSet) GetName() string {
	return "Approved"
}

func (s *ApprovedGeneralSet) GetLastSet() sets.Set {
	return s.LastSet
}

func (s *ApprovedGeneralSet) SetLastSet(set sets.Set) {
	s.LastSet = set
	s.Gaps = gaps.NewGaps(set.GetLastTime())
}

func (s *ApprovedGeneralSet) GetAtLeastOne() bool {
	return s.AtLeastOne
}

func (s *ApprovedGeneralSet) SetAtLeastOne(b bool) {
	s.AtLeastOne = b
}

func (s *ApprovedGeneralSet) GetNums() int {
	return s.Nums
}

func (s *ApprovedGeneralSet) IncreaseNums() {
	s.Nums++
}

func (s *ApprovedGeneralSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}
