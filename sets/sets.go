package sets

import (
	"github.com/MoraGames/clockyAnalyzer/gaps"
	"github.com/MoraGames/clockyAnalyzer/time"
)

type Set struct {
	Name     string
	Label    string
	Nums     int
	LastTime time.Time
	Verify   func(h1, h2, m1, m2 int) bool
	Gaps     *gaps.Gaps
}

func NewSet(name, label string, verify func(h1, h2, m1, m2 int) bool, lastTime time.Time) *Set {
	var set = Set{Name: name, Label: label, Nums: 0, Verify: verify}

	if !set.Verify(lastTime.SplitTime()) {
		panic("sets: NewSet: lastTime is not valid")
	}
	set.LastTime = lastTime
	set.Gaps = gaps.NewGaps(lastTime)

	return &set
}
