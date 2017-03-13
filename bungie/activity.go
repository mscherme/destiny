package bungie

import (
	"fmt"
	"time"
)

type ActivityFilter string

const (
	None              ActivityFilter = "None"
	Story             ActivityFilter = "Story"
	Strike            ActivityFilter = "Strike"
	Raid              ActivityFilter = "Raid"
	AllPvP            ActivityFilter = "AllPvP"
	Patrol            ActivityFilter = "Patrol"
	AllPvE            ActivityFilter = "AllPvE"
	PvPIntroduction   ActivityFilter = "PvPIntroduction"
	ThreeVsThree      ActivityFilter = "ThreeVsThree"
	Control           ActivityFilter = "Control"
	Lockdown          ActivityFilter = "Lockdown"
	Team              ActivityFilter = "Team"
	FreeForAll        ActivityFilter = "FreeForAll"
	Nightfall         ActivityFilter = "Nightfall"
	Heroic            ActivityFilter = "Heroic"
	AllStrikes        ActivityFilter = "AllStrikes"
	IronBanner        ActivityFilter = "IronBanner"
	AllArena          ActivityFilter = "AllArena"
	Arena             ActivityFilter = "Arena"
	ArenaChallenge    ActivityFilter = "ArenaChallenge"
	TrialsOfOsiris    ActivityFilter = "TrialsOfOsiris"
	Elimination       ActivityFilter = "Elimination"
	Rift              ActivityFilter = "Rift"
	Mayhem            ActivityFilter = "Mayhem"
	ZoneControl       ActivityFilter = "ZoneControl"
	Racing            ActivityFilter = "Racing"
	Supremacy         ActivityFilter = "Supremacy"
	PrivateMatchesAll ActivityFilter = "PrivateMatchesAll"
)

type ActivityID int64

const (
	WrathOfTheMachineNM ActivityID = 260765522
	WrathOfTheMachineHM ActivityID = 1387993552
)

type ActivityDetails struct {
	ReferenceID              int64  `json:"referenceId"`
	InstanceID               string `json:"instanceId"`
	Mode                     int    `json:"mode"`
	ActivityTypeHashOverride int    `json:"activityTypeHashOverride"`
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

func (b *API) LookupActivities(c *Character, filter ActivityFilter, count, page int) ([]*ActivityRecord, error) {
	url := fmt.Sprintf("Stats/ActivityHistory/%d/%s/%s/?mode=%s&count=%d&page=%d",
		c.CharacterBase.MembershipType,
		c.CharacterBase.MembershipID,
		c.CharacterBase.CharacterID,
		filter, count, page)
	var x activityHistoryJSON
	err := b.get(url, &x, false)
	if err != nil {
		return nil, err
	}
	return x.Response.Data.Activities, nil
}
