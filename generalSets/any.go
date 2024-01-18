package generalSets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/sets"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ GeneralSet = new(AnyGeneralSet)

type AnyGeneralSet struct {
	LastSet    sets.Set
	AtLeastOne bool
	Nums       int
	Gaps       *gaps.Gaps
}

func NewGeneralSet_Any(lastSet sets.Set) *AnyGeneralSet {
	return &AnyGeneralSet{lastSet, false, 0, gaps.NewGaps(time.NewTime(00, 00))}
}

func (s *AnyGeneralSet) GetName() string {
	return "Any"
}

func (s *AnyGeneralSet) GetLastSet() sets.Set {
	return s.LastSet
}

func (s *AnyGeneralSet) SetLastSet(set sets.Set) {
	s.LastSet = set
	s.Gaps = gaps.NewGaps(set.GetLastTime())
}

func (s *AnyGeneralSet) GetAtLeastOne() bool {
	return s.AtLeastOne
}

func (s *AnyGeneralSet) SetAtLeastOne(b bool) {
	s.AtLeastOne = b
}

func (s *AnyGeneralSet) GetNums() int {
	return s.Nums
}

func (s *AnyGeneralSet) IncreaseNums() {
	s.Nums++
}

func (s *AnyGeneralSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}
