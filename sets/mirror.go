package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

// Check if the set implements the interface
var _ Set = new(MirrorSet)

type MirrorSet struct {
	LastTime time.Time
	Name     string
	Label    string
	Nums     int
	Gaps     *gaps.Gaps
}

func NewSet_Mirror(name, label string) *MirrorSet {
	lstTime := time.NewTime(23, 56)
	return &MirrorSet{lstTime, name, label, 0, gaps.NewGaps(lstTime)}
}

func (s *MirrorSet) GetLastTime() time.Time {
	return s.LastTime
}

func (s *MirrorSet) GetName() string {
	return s.Name
}

func (s *MirrorSet) GetLabel() string {
	return s.Label
}

func (s *MirrorSet) GetNums() int {
	return s.Nums
}

func (s *MirrorSet) IncreaseNums() {
	s.Nums++
}

func (s *MirrorSet) GetGaps() *gaps.Gaps {
	return s.Gaps
}

func (s *MirrorSet) Verify(h1, h2, m1, m2 int) bool {
	return h1 == m2 && h2 == m1 && h1 != h2
}
