package bungie

import (
	"fmt"
	"time"
)

type Mode int64

const (
	None                Mode = 0
	Story               Mode = 2
	Strike              Mode = 3
	Raid                Mode = 4
	AllPvP              Mode = 5
	Patrol              Mode = 6
	AllPvE              Mode = 7
	PvPIntroduction     Mode = 8
	ThreeVsThree        Mode = 9
	Control             Mode = 10
	Lockdown            Mode = 11
	Team                Mode = 12
	FreeForAll          Mode = 13
	TrialsOfOsiris      Mode = 14
	Doubles             Mode = 15
	Nightfall           Mode = 16
	Heroic              Mode = 17
	AllStrikes          Mode = 18
	IronBanner          Mode = 19
	AllArena            Mode = 20
	Arena               Mode = 21
	ArenaChallenge      Mode = 22
	Elimination         Mode = 23
	Rift                Mode = 24
	AllMayhem           Mode = 25
	MayhemClash         Mode = 26
	MayhemRumble        Mode = 27
	ZoneControl         Mode = 28
	Racing              Mode = 29
	ArenaElderChallenge Mode = 30
	Supremacy           Mode = 31
	PrivateMatchesAll   Mode = 32
	SupremacyRumble     Mode = 33
	SupremacyClash      Mode = 34
	SupremacyInferno    Mode = 35
	SupremacyMayhem     Mode = 36
)

func (a Mode) String() string {
	switch a {
	case None:
		return "None"
	case Story:
		return "Story"
	case Strike:
		return "Strike"
	case Raid:
		return "Raid"
	case AllPvP:
		return "AllPvP"
	case Patrol:
		return "Patrol"
	case AllPvE:
		return "AllPvE"
	case PvPIntroduction:
		return "PvPIntroduction"
	case ThreeVsThree:
		return "Skirmish"
	case Control:
		return "Control"
	case Lockdown:
		return "Salvage"
	case Team:
		return "Clash"
	case FreeForAll:
		return "Rumble"
	case TrialsOfOsiris:
		return "TrialsOfOsiris"
	case Doubles:
		return "Doubles"
	case Nightfall:
		return "Nightfall"
	case Heroic:
		return "Heroic"
	case AllStrikes:
		return "AllStrikes"
	case IronBanner:
		return "IronBanner"
	case AllArena:
		return "AllArena"
	case Arena:
		return "Arena"
	case ArenaChallenge:
		return "ArenaChallenge"
	case Elimination:
		return "Elimination"
	case Rift:
		return "Rift"
	case AllMayhem:
		return "AllMayhem"
	case MayhemClash:
		return "MayhemClash"
	case MayhemRumble:
		return "MayhemRumble"
	case ZoneControl:
		return "ZoneControl"
	case Racing:
		return "Racing"
	case ArenaElderChallenge:
		return "ArenaElderChallenge"
	case Supremacy:
		return "Supremacy"
	case PrivateMatchesAll:
		return "PrivateMatchesAll"
	case SupremacyRumble:
		return "SupremacyRumble"
	case SupremacyClash:
		return "SupremacyClash"
	case SupremacyInferno:
		return "SupremacyInferno"
	case SupremacyMayhem:
		return "SupremacyMayhem"
	}
	return "Unknown"
}

type ActivityDetails struct {
	ReferenceID              int64  `json:"referenceId"`
	InstanceID               string `json:"instanceId"`
	Mode                     Mode   `json:"mode"`
	ActivityTypeHashOverride int64  `json:"activityTypeHashOverride"`
	IsPrivate                bool   `json:"isPrivate"`
}

type ActivityRecord struct {
	Period          time.Time        `json:"period"`
	ActivityDetails *ActivityDetails `json:"activityDetails"`
	Values          *Values          `json:"values"`
}

type activityHistoryJSON struct {
	jsonStatusFields
	Response struct {
		Data struct {
			Activities []*ActivityRecord `json:"activities"`
		} `json:"data"`
	} `json:"Response"`
}

func (b *API) LookupActivities(c *Character, mode Mode, count, page int) ([]*ActivityRecord, error) {
	url := fmt.Sprintf("Stats/ActivityHistory/%d/%s/%s/?mode=%d&count=%d&page=%d",
		c.CharacterBase.MembershipType,
		c.CharacterBase.MembershipID,
		c.CharacterBase.CharacterID,
		mode, count, page)
	var x activityHistoryJSON
	err := b.get(url, &x, false)
	if err != nil {
		return nil, err
	}
	return x.Response.Data.Activities, nil
}
