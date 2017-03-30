package main

import (
	"flag"
	"fmt"
	"log"
	"sort"

	"github.com/mscherme/destiny/bungie"
)

var gamertag = flag.String("gamertag", "mscherme", "")

func main() {
	flag.Parse()

	var err error
	b, err = bungie.New()
	if err != nil {
		log.Fatal(err)
	}
	account, err := b.LookupAccount(bungie.XBOX, *gamertag)
	if err != nil {
		log.Fatal(err)
	}

	for _, c := range account.Characters {
		page := 0
		for {
			activities, err := b.LookupActivities(c, bungie.Raid, 100, page)
			if err != nil {
				log.Fatal(err)
			}
			if len(activities) > 0 {
				processActivities(activities)
			}

			if len(activities) < 100 {
				break
			}
			page++
		}
	}

	var l fireteamList
	for _, e := range fireteam {
		l = append(l, e)
	}
	sort.Sort(sort.Reverse(l))
	fmt.Println("gamertag,total,heroic,normal")
	for _, e := range l {
		fmt.Printf("%q,%d,%d,%d\n", e.gamertag, e.total(), e.heroic, e.normal)
	}
}

type fireteamList []*fireteamEntry

type fireteamEntry struct {
	gamertag string
	heroic   int
	normal   int
}

func (f *fireteamEntry) total() int {
	return f.heroic + f.normal
}

func (l fireteamList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l fireteamList) Len() int {
	return len(l)
}

func (l fireteamList) Less(i, j int) bool {
	return l[i].total() < l[j].total()
}

var (
	b        *bungie.API
	fireteam = map[string]*fireteamEntry{}
)

var raids = map[int64]bool{
	bungie.WrathOfTheMachineNormal: false,
	bungie.WrathOfTheMachineHeroic: true,
	bungie.CrotasEndNormal:         false,
	bungie.CrotasEndHeroic:         true,
	bungie.CrotasEndAgeOfTriumph:   true,
	bungie.VaultOfGlassNormal:      false,
	bungie.VaultOfGlassHeroic:      true,
	bungie.KingsFallNormal:         false,
	bungie.KingsFallHeroic:         true,
}

func processActivities(activities []*bungie.ActivityRecord) {
	for _, activity := range activities {
		heroic, ok := raids[activity.ActivityDetails.ReferenceID]
		if !ok {
			log.Printf("Unknown Raid: %d", activity.ActivityDetails.ReferenceID)
			continue
		}
		if activity.Values.Completed.Basic.Value != 1 {
			continue
		}
		pgcr, err := b.LookupPostGameCarnageReport(activity)
		if err != nil {
			log.Fatal(err)
		}
		for _, entry := range pgcr.Entries {
			gamertag := entry.Player.DestinyUserInfo.DisplayName
			if entry.Values.Completed.Basic.Value == 0 {
				continue
			}
			f := fireteam[gamertag]
			if f == nil {
				f = &fireteamEntry{gamertag: gamertag}
				fireteam[gamertag] = f
			}
			if heroic {
				f.heroic++
			} else {
				f.normal++
			}
		}
	}
}
