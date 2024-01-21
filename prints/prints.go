package prints

import (
	"fmt"

	"github.com/MoraGames/clockyAnalyzer/groups"
	"github.com/MoraGames/clockyAnalyzer/sets"
	"github.com/MoraGames/clockyAnalyzer/time"
	"github.com/MoraGames/clockyAnalyzer/utils"
	"github.com/fatih/color"
)

const (
	PrintTimeStats       = false
	PrintFinalStats      = true
	PrintOnlyUnderReview = false
	PrintAllGaps         = false
)

var (
	SymbolsColor = color.New(color.FgRed, color.Bold).SprintFunc()
	TitlesColor  = color.New(color.FgRed, color.Bold, color.Underline).SprintFunc()

	GroupsNameColor       = color.New(color.FgHiYellow, color.Bold).SprintFunc()
	GroupsPropertiesColor = color.New(color.FgYellow).SprintFunc()
	GroupsValuesColor     = color.New(color.FgYellow).SprintFunc()

	SetsNameColor       = color.New(color.FgHiGreen, color.Bold).SprintFunc()
	SetsPropertiesColor = color.New(color.FgGreen).SprintFunc()
	SetsValuesColor     = color.New(color.FgGreen).SprintFunc()
)

func PrintTime(allSets []*sets.Set, currentTime time.Time, totalTimes int) {
	if PrintTimeStats {
		fmt.Printf("%v %v (%v):\n", SymbolsColor(">>>"), TitlesColor("TIME STATS"), TitlesColor(currentTime.Stringify()))
		for _, set := range allSets {
			if !PrintOnlyUnderReview || set.Label == "Under-Review" {
				fmt.Printf(" |- %v = %v/%v (Gap = %v)\n", SetsNameColor(set.Name), SetsValuesColor(set.Nums), SetsValuesColor(totalTimes), SetsValuesColor(set.Gaps.Cur))
			}
		}
		fmt.Printf("\n")
	}
}

func PrintFinal(allSets []*sets.Set, allGroups map[string]*groups.Group, totalTimes int) {
	if PrintFinalStats {
		fmt.Printf("%s %s:\n", SymbolsColor(">>>"), TitlesColor("FINAL STATS"))
		for _, set := range allSets {
			if !PrintOnlyUnderReview || set.Label == "Under-Review" {
				fmt.Printf(" |- %s = %v/%v\n", SetsNameColor(set.Name), SetsValuesColor(set.Nums), SetsValuesColor(totalTimes))
				fmt.Printf(" |   %v:\n", SetsPropertiesColor("Gaps"))
				if PrintAllGaps {
					fmt.Printf(" |    |- %v: %v\n", SetsPropertiesColor("All"), SetsValuesColor(set.Gaps.All))
				}
				fmt.Printf(" |    |- %v: %vm (%v - %v)\n", SetsPropertiesColor("Min"), SetsValuesColor(set.Gaps.Min.Val), SetsValuesColor(set.Gaps.Min.Prv.Stringify()), SetsValuesColor(set.Gaps.Min.Cur.Stringify()))
				fmt.Printf(" |    |- %v: %vm (%v - %v)\n", SetsPropertiesColor("Max"), SetsValuesColor(set.Gaps.Max.Val), SetsValuesColor(set.Gaps.Max.Prv.Stringify()), SetsValuesColor(set.Gaps.Max.Cur.Stringify()))
				gapsAvgMinutes, gapsAvgSeconds := utils.DecimalToSexagesimal(set.Gaps.Avg)
				fmt.Printf(" |    |- %v: %vm%vs (%sm)\n\n", SetsPropertiesColor("Avg"), SetsValuesColor(gapsAvgMinutes), SetsValuesColor(gapsAvgSeconds), SetsValuesColor(fmt.Sprintf("%f", set.Gaps.Avg)[:utils.FloatWidth(set.Gaps.Avg, 3)]))
			}
		}
		for _, group := range allGroups {
			if !PrintOnlyUnderReview || group.Name == "Under-Review" {
				fmt.Printf(" |- %s = %v/%v (%v sets)\n", GroupsNameColor(group.Name), GroupsValuesColor(group.EventNums), GroupsValuesColor(totalTimes), GroupsValuesColor(group.SetNums))
				fmt.Printf(" |   %v:\n", GroupsPropertiesColor("Gaps"))
				if PrintAllGaps {
					fmt.Printf(" |    |- %v: %v\n", GroupsPropertiesColor("All"), GroupsValuesColor(group.Gaps.All))
				}
				fmt.Printf(" |    |- %v: %vm (%v - %v)\n", GroupsPropertiesColor("Min"), GroupsValuesColor(group.Gaps.Min.Val), GroupsValuesColor(group.Gaps.Min.Prv.Stringify()), GroupsValuesColor(group.Gaps.Min.Cur.Stringify()))
				fmt.Printf(" |    |- %v: %vm (%v - %v)\n", GroupsPropertiesColor("Max"), GroupsValuesColor(group.Gaps.Max.Val), GroupsValuesColor(group.Gaps.Max.Prv.Stringify()), GroupsValuesColor(group.Gaps.Max.Cur.Stringify()))
				gapsAvgMinutes, gapsAvgSeconds := utils.DecimalToSexagesimal(group.Gaps.Avg)
				fmt.Printf(" |    |- %v: %vm%vs (%sm)\n\n", GroupsPropertiesColor("Avg"), GroupsValuesColor(gapsAvgMinutes), GroupsValuesColor(gapsAvgSeconds), GroupsValuesColor(fmt.Sprintf("%f", group.Gaps.Avg)[:utils.FloatWidth(group.Gaps.Avg, 3)]))
			}
		}
	}
}
