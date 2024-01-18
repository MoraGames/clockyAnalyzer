package sets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/time"
)

type Set interface {
	GetLastTime() time.Time
	GetName() string
	GetLabel() string
	GetNums() int
	IncreaseNums()
	GetGaps() *gaps.Gaps
	Verify(h1, h2, m1, m2 int) bool
}
