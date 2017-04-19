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

	for _, raid := range []string{"All", "Vault of Glass", "Crota's End", "King's Fall", "Wrath of the Machine"} {
		var l fireteamList
		for _, e := range stats[raid] {
			l = append(l, e)
		}
		sort.Sort(sort.Reverse(l))
		fmt.Println(raid)
		fmt.Println("gamertag,total,normal,heroic,aot,cm")
		for i, e := range l {
			if e.total() < 3 && i >= 20 {
				break
			}
			fmt.Printf("%q,%d,%d,%d,%d,%d\n", e.gamertag, e.total(), e.normal, e.heroic, e.aot, e.cm)
		}
		fmt.Println()
	}
}

type fireteamList []*fireteamEntry

type fireteamEntry struct {
	gamertag string
	normal   int
	heroic   int
	aot      int
	cm       int
}

func (f *fireteamEntry) total() int {
	return f.heroic + f.normal + f.cm + f.aot
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

type raidFireteamStats map[string]*fireteamEntry

var (
	b     *bungie.API
	stats = map[string]raidFireteamStats{}
)

func processActivities(activities []*bungie.ActivityRecord) {
	for _, activity := range activities {
		raidDetails, ok := bungie.Raids[activity.ActivityDetails.ReferenceID]
		raidActivityDetails, err := b.ManifestActivity(activity.ActivityDetails.ReferenceID)
		if err != nil {
			log.Fatal(err)
		}
		if !ok {
			log.Printf("Unknown Raid: %d", activity.ActivityDetails.ReferenceID)
			log.Print(raidActivityDetails)
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
			for _, raidName := range []string{"All", raidActivityDetails.ActivityName} {
				fireteam, ok := stats[raidName]
				if !ok {
					fireteam = raidFireteamStats{}
					stats[raidName] = fireteam
				}

				f := fireteam[gamertag]
				if f == nil {
					f = &fireteamEntry{gamertag: gamertag}
					fireteam[gamertag] = f
				}
				switch raidDetails.Difficulty {
				case bungie.NormalRaid:
					f.normal++
				case bungie.HeroicRaid:
					f.heroic++
				case bungie.AgeOfTriumph:
					f.aot++
				case bungie.AgeOfTriumphCM:
					f.cm++
				}
			}
		}
	}
}
