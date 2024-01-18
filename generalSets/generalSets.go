package generalSets

import (
	"other/bp/cuwu_setstest/gaps"
	"other/bp/cuwu_setstest/sets"
)

type GeneralSet interface {
	GetName() string
	GetLastSet() sets.Set
	SetLastSet(sets.Set)
	GetAtLeastOne() bool
	SetAtLeastOne(bool)
	GetNums() int
	IncreaseNums()
	GetGaps() *gaps.Gaps
}
