package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(ShortRapidFallSet)

type ShortRapidFallSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_ShortRapidFall(name, label string) *ShortRapidFallSet {
	var set = ShortRapidFallSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(15, 31)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Mirror: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *ShortRapidFallSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *ShortRapidFallSet) GetName() string {
	return s.Name
}

func (s *ShortRapidFallSet) GetLabel() string {
	return s.Label
}

func (s *ShortRapidFallSet) GetNums() int {
	return s.Nums
}

func (s *ShortRapidFallSet) IncreaseNums() {
	s.Nums++
}

func (s *ShortRapidFallSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *ShortRapidFallSet) Verify(_, h2, m1, m2 int) bool {
	return m1 == h2-2 && m2 == m1-2
}
