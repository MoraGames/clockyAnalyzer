package groups

import (
	"github.com/MoraGames/clockyAnalyzer/gaps"
	"github.com/MoraGames/clockyAnalyzer/sets"
	"github.com/MoraGames/clockyAnalyzer/time"
)

type Group struct {
	Name       string
	EventNums  int
	SetNums    int
	AtLeastOne bool
	Gaps       *gaps.Gaps
	LastSet    *sets.Set
}

func NewGroup(name string) *Group {
	return &Group{Name: name, EventNums: 0, SetNums: 0, AtLeastOne: false, Gaps: gaps.NewGaps(time.NewTime(00, 00)), LastSet: nil}
}

func (g *Group) IsReady() bool {
	return g.LastSet != nil
}

func (g *Group) SetSetNums(setNums int) {
	if setNums < 0 {
		panic("groups: SetSetNums: setNums is negative")
	}
	g.SetNums = setNums
}

func (g *Group) SetLastSet(set *sets.Set) {
	if set == nil {
		panic("groups: SetLastSet: set is nil")
	}
	if g.Name != "Any" && set.Label != g.Name {
		panic("groups: SetLastSet: the set cannot belong to this group with this label")
	}
	g.LastSet = set
	g.Gaps = gaps.NewGaps(set.LastTime)
}
