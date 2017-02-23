package main

import (
	"strings"

	"github.com/mscherme/destiny/bungie"
	"github.com/mscherme/destiny/gear"
)

// engram
// year
// tier
// gear set
// class
// item type

var itemTypeOrder = map[string]int{
	"Auto Rifle":      1,
	"Hand Cannon":     2,
	"Pulse Rifle":     3,
	"Scout Rifle":     4,
	"Fusion Rifle":    5,
	"Shotgun":         6,
	"Sidearm":         7,
	"Sniper Rifle":    8,
	"Machine Gun":     9,
	"Rocket Launcher": 10,
	"Sword":           11,

	"Helmet":       12,
	"Gauntlets":    13,
	"Chest Armor":  14,
	"Leg Armor":    15,
	"Hunter Cloak": 16,
	"Titan Mark":   17,
	"Warlock Bond": 18,
}

var gearSetOrder = map[gear.Set]int{
	gear.VaultOfGlass:   1,
	gear.CrotasEnd:      2,
	gear.PrisonOfElders: 3,

	gear.KingsFall:         4,
	gear.IronBannerY2:      5,
	gear.ChallengeOfElders: 6,

	gear.WrathOfTheMachine:    7,
	gear.RiseOfIronRecordBook: 8,

	gear.TrialsOfOsirisY3: 9,
	gear.IronBannerY3:     10,
}

type sortableItemList []*bungie.Item

func (l sortableItemList) Less(i, j int) bool {
	// Get Info objects
	iInfo := lookup(l[i].ItemHash)
	jInfo := lookup(l[j].ItemHash)

	// Sort engrams last
	iEngram := strings.Contains(iInfo.ItemName, "Engram")
	jEngram := strings.Contains(jInfo.ItemName, "Engram")
	if iEngram != jEngram {
		return jEngram
	}

	// Sort by Year
	iLL := l[i].PrimaryStat.Value > 170
	jLL := l[j].PrimaryStat.Value > 170
	if iLL != jLL {
		return jLL
	}

	// Sort by Tier
	if iInfo.TierType != jInfo.TierType {
		return iInfo.TierType < jInfo.TierType
	}

	// Sort by gear.Set
	iGS := gearSetOrder[gear.SetForItem(iInfo)]
	jGS := gearSetOrder[gear.SetForItem(jInfo)]
	if iGS == 0 {
		iGS = 10
	}
	if jGS == 0 {
		jGS = 10
	}
	if iGS != jGS {
		return iGS < jGS
	}
	if iGS == 10 {
		return strings.Compare(iInfo.ItemName, jInfo.ItemName) < 0
	}

	// Sort by Class
	if iInfo.ClassType != jInfo.ClassType {
		return iInfo.ClassType < jInfo.ClassType
	}

	// Sort by item type
	iType := itemTypeOrder[iInfo.ItemTypeName]
	jType := itemTypeOrder[jInfo.ItemTypeName]
	return iType < jType
}

func (l sortableItemList) Len() int {
	return len(l)
}

func (l sortableItemList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
