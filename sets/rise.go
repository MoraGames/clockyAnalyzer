package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(RiseSet)

type RiseSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_Rise(name, label string) *RiseSet {
	var set = RiseSet{Name: name, Label: label, Nums: 0}

	lstTime := time.NewTime(23, 45)
	if !set.Verify(lstTime.SplitTime()) {
		panic("sets: NewSet_Mirror: lstTime is not valid")
	}
	set.LastTime = lstTime
	set.Gaps = gaps.NewGaps(lstTime)

	return &set
}

func (s *RiseSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *RiseSet) GetName() string {
	return s.Name
}

func (s *RiseSet) GetLabel() string {
	return s.Label
}

func (s *RiseSet) GetNums() int {
	return s.Nums
}

func (s *RiseSet) IncreaseNums() {
	s.Nums++
}

func (s *RiseSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *RiseSet) Verify(h1, h2, m1, m2 int) bool {
	return h2 == h1+1 && m1 == h2+1 && m2 == m1+1
}
