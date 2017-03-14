package main

import (
	"flag"
	"fmt"
	"log"

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
	processAccountActivities(account)
}

type key struct {
	class    string
	weapon   string
	activity string
}

type value struct {
	kills     int64
	precision int64
}

var killCounts = map[key]*value{}

func addKills(k key, kills, precision int64) {
	v := killCounts[k]
	if v == nil {
		v = &value{}
		killCounts[k] = v
	}
	v.kills += kills
	v.precision += precision
}

var b *bungie.API

func processActivities(activities []*bungie.ActivityRecord, mode bungie.Mode) {
	for _, activity := range activities {
		pgcr, err := b.LookupPostGameCarnageReport(activity)
		if err != nil {
			log.Fatal(err)
		}
		for _, entry := range pgcr.Entries {
			if entry.Player.DestinyUserInfo.DisplayName != *gamertag {
				continue
			}
			for _, weapon := range entry.Extended.Weapons {
				item, err := b.ManifestInventoryItem(weapon.ReferenceID)
				if err != nil {
					log.Fatal(err)
				}

				killCount := int64(weapon.Values.UniqueWeaponKills.Basic.Value)
				pKillCount := int64(weapon.Values.UniqueWeaponPrecisionKills.Basic.Value)

				k := key{weapon: item.ItemName}
				for _, k.activity = range []string{mode.String(), "All"} {
					for _, k.class = range []string{entry.Player.CharacterClass, "All"} {
						addKills(k, killCount, pKillCount)
					}
				}
			}
		}
	}
}

func processAccountActivities(account *bungie.Account) {
	for _, mode := range []bungie.Mode{bungie.AllPvE, bungie.AllPvP} {
		for _, c := range account.Characters {
			page := 0
			for {
				activities, err := b.LookupActivities(c, mode, 100, page)
				if err != nil {
					log.Fatal(err)
				}
				if len(activities) > 0 {
					processActivities(activities, mode)
				}

				if len(activities) < 100 {
					break
				}
				page++
			}
		}
	}

	for k, count := range killCounts {
		fmt.Printf("%q,%q,%q,%d,%d\n", k.class, k.weapon, k.activity, count.kills, count.precision)
	}
}
