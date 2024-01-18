package generalSets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/sets"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ GeneralSet = new(UnderReviewGeneralSet)

type UnderReviewGeneralSet struct {
	LastSet    sets.Set
	AtLeastOne bool
	Nums       int
	Gaps       *gaps.Gaps
}

func NewGeneralSet_UnderReview(lastSet sets.Set) *UnderReviewGeneralSet {
	return &UnderReviewGeneralSet{lastSet, false, 0, gaps.NewGaps(time.NewTime(23, 54))}
}

func (s *UnderReviewGeneralSet) GetName() string {
	return "Under-Review"
}

func (s *UnderReviewGeneralSet) GetLastSet() sets.Set {
	return s.LastSet
}

func (s *UnderReviewGeneralSet) SetLastSet(set sets.Set) {
	s.LastSet = set
	s.Gaps = gaps.NewGaps(set.GetLastTime())
}

func (s *UnderReviewGeneralSet) GetAtLeastOne() bool {
	return s.AtLeastOne
}

func (s *UnderReviewGeneralSet) SetAtLeastOne(b bool) {
	s.AtLeastOne = b
}

func (s *UnderReviewGeneralSet) GetNums() int {
	return s.Nums
}

func (s *UnderReviewGeneralSet) IncreaseNums() {
	s.Nums++
}

func (s *UnderReviewGeneralSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}
