package bungie

import (
	"fmt"
	"time"
)

type Player struct {
	DestinyUserInfo struct {
		IconPath       string `json:"iconPath"`
		MembershipType int64  `json:"membershipType"`
		MembershipID   string `json:"membershipId"`
		DisplayName    string `json:"displayName"`
	} `json:"destinyUserInfo"`
	CharacterClass    string `json:"characterClass"`
	CharacterLevel    int64  `json:"characterLevel"`
	LightLevel        int64  `json:"lightLevel"`
	BungieNetUserInfo struct {
		IconPath       string `json:"iconPath"`
		MembershipType int64  `json:"membershipType"`
		MembershipID   string `json:"membershipId"`
		DisplayName    string `json:"displayName"`
	} `json:"bungieNetUserInfo"`
}

type PostGameCarnageReport struct {
	Period          time.Time        `json:"period"`
	ActivityDetails *ActivityDetails `json:"activityDetails"`
	Entries         []struct {
		Standing int64 `json:"standing"`
		Score    struct {
			Basic struct {
				Value        float64 `json:"value"`
				DisplayValue string  `json:"displayValue"`
			} `json:"basic"`
		} `json:"score"`
		Player      *Player `json:"player"`
		CharacterID string  `json:"characterId"`
		Values      *Values `json:"values"`
		Extended    struct {
			Weapons []struct {
				ReferenceID int64 `json:"referenceId"`
				Values      struct {
					UniqueWeaponKills               Stat `json:"uniqueWeaponKills"`
					UniqueWeaponPrecisionKills      Stat `json:"uniqueWeaponPrecisionKills"`
					UniqueWeaponKillsPrecisionKills Stat `json:"uniqueWeaponKillsPrecisionKills"`
				} `json:"values"`
			} `json:"weapons"`
			Values map[string]Stat `json:"values"`
		} `json:"extended"`
	} `json:"entries"`
	Teams []struct {
		TeamID   int64  `json:"teamId"`
		Standing Stat   `json:"standing"`
		Score    Stat   `json:"score"`
		TeamName string `json:"teamName"`
	} `json:"teams"`
}

type pgcrJSON struct {
	jsonStatusFields
	Response struct {
		Data *PostGameCarnageReport `json:"data"`
	} `json:"Response"`
}

func (b *API) LookupPostGameCarnageReport(activity *ActivityRecord) (*PostGameCarnageReport, error) {
	url := fmt.Sprintf("Stats/PostGameCarnageReport/%s/?definitions=false", activity.ActivityDetails.InstanceID)
	var x pgcrJSON
	err := b.get(url, &x, true)
	if err != nil {
		return nil, err
	}
	return x.Response.Data, nil
}
