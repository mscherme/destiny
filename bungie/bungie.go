package bungie

import (
	"errors"
	"fmt"
	"time"
)

const (
	XBOX int64 = 1
	PSN  int64 = 2
)

type accountJSON struct {
	jsonStatusFields
	Response string `json:"Response"`
}

func (b *API) LookupAccount(membershipType int64, displayName string) (*Account, error) {
	url := fmt.Sprintf("%d/Stats/GetMembershipIdByDisplayName/%s/",
		membershipType, displayName)
	var x accountJSON
	err := b.get(url, &x, true)
	if err != nil {
		return nil, err
	}
	account := &Account{
		MembershipType: membershipType,
		MembershipID:   x.Response,
	}
	if err = b.Refresh(account); err != nil {
		return nil, err
	}
	return account, nil
}

type Item struct {
	ItemHash       int64  `json:"itemHash"`
	ItemID         string `json:"itemId"`
	Quantity       int64  `json:"quantity"`
	DamageType     int64  `json:"damageType"`
	DamageTypeHash int64  `json:"damageTypeHash"`
	PrimaryStat    struct {
		StatHash     int64 `json:"statHash"`
		Value        int64 `json:"value"`
		MaximumValue int64 `json:"maximumValue"`
	} `json:"primaryStat"`
	IsGridComplete bool  `json:"isGridComplete"`
	TransferStatus int64 `json:"transferStatus"`
	State          int64 `json:"state"`
	CharacterIndex int64 `json:"characterIndex"`
	BucketHash     int64 `json:"bucketHash"`
}

type Inventory struct {
	Items      []*Item `json:"items"`
	Currencies []struct {
		ItemHash int64 `json:"itemHash"`
		Value    int64 `json:"value"`
	} `json:"currencies"`
}

type Account struct {
	MembershipID   string       `json:"membershipId"`
	MembershipType int64        `json:"membershipType"`
	Characters     []*Character `json:"characters"`
	Inventory      *Inventory   `json:"inventory"`
	GrimoireScore  int64        `json:"grimoireScore"`
	Versions       int64        `json:"versions"`
}

type CharacterStat struct {
	StatHash     int64 `json:"statHash"`
	Value        int64 `json:"value"`
	MaximumValue int64 `json:"maximumValue"`
}

type Character struct {
	CharacterBase struct {
		MembershipID             string    `json:"membershipId"`
		MembershipType           int64     `json:"membershipType"`
		CharacterID              string    `json:"characterId"`
		DateLastPlayed           time.Time `json:"dateLastPlayed"`
		MinutesPlayedThisSession string    `json:"minutesPlayedThisSession"`
		MinutesPlayedTotal       string    `json:"minutesPlayedTotal"`
		PowerLevel               int64     `json:"powerLevel"`
		RaceHash                 int64     `json:"raceHash"`
		GenderHash               int64     `json:"genderHash"`
		ClassHash                int64     `json:"classHash"`
		CurrentActivityHash      int64     `json:"currentActivityHash"`
		LastCompletedStoryHash   int64     `json:"lastCompletedStoryHash"`
		Stats                    struct {
			Defense    CharacterStat `json:"STAT_DEFENSE"`
			Intellect  CharacterStat `json:"STAT_INTELLECT"`
			Discipline CharacterStat `json:"STAT_DISCIPLINE"`
			Strength   CharacterStat `json:"STAT_STRENGTH"`
			Light      CharacterStat `json:"STAT_LIGHT"`
			Armor      CharacterStat `json:"STAT_ARMOR"`
			Agility    CharacterStat `json:"STAT_AGILITY"`
			Recovery   CharacterStat `json:"STAT_RECOVERY"`
			Optics     CharacterStat `json:"STAT_OPTICS"`
		} `json:"stats"`
		Customization struct {
			Personality  int64 `json:"personality"`
			Face         int64 `json:"face"`
			SkinColor    int64 `json:"skinColor"`
			LipColor     int64 `json:"lipColor"`
			EyeColor     int64 `json:"eyeColor"`
			HairColor    int64 `json:"hairColor"`
			FeatureColor int64 `json:"featureColor"`
			DecalColor   int64 `json:"decalColor"`
			WearHelmet   bool  `json:"wearHelmet"`
			HairIndex    int64 `json:"hairIndex"`
			FeatureIndex int64 `json:"featureIndex"`
			DecalIndex   int64 `json:"decalIndex"`
		} `json:"customization"`
		GrimoireScore int64 `json:"grimoireScore"`
		PeerView      struct {
			Equipment []struct {
				ItemHash int64 `json:"itemHash"`
				//Dyes     []interface{} `json:"dyes"`
			} `json:"equipment"`
		} `json:"peerView"`
		GenderType         int64 `json:"genderType"`
		ClassType          int64 `json:"classType"`
		BuildStatGroupHash int64 `json:"buildStatGroupHash"`
	} `json:"characterBase"`
	LevelProgression struct {
		DailyProgress       int64 `json:"dailyProgress"`
		WeeklyProgress      int64 `json:"weeklyProgress"`
		CurrentProgress     int64 `json:"currentProgress"`
		Level               int64 `json:"level"`
		Step                int64 `json:"step"`
		ProgressToNextLevel int64 `json:"progressToNextLevel"`
		NextLevelAt         int64 `json:"nextLevelAt"`
		ProgressionHash     int64 `json:"progressionHash"`
	} `json:"levelProgression"`
	EmblemPath         string  `json:"emblemPath"`
	BackgroundPath     string  `json:"backgroundPath"`
	EmblemHash         int64   `json:"emblemHash"`
	CharacterLevel     int64   `json:"characterLevel"`
	BaseCharacterLevel int64   `json:"baseCharacterLevel"`
	IsPrestigeLevel    bool    `json:"isPrestigeLevel"`
	PercentToNextLevel float64 `json:"percentToNextLevel"`
}

type accountSummaryJSON struct {
	jsonStatusFields
	Response struct {
		Data Account `json:"data"`
	} `json:"Response"`
}

func (b *API) Refresh(a *Account) error {
	url := fmt.Sprintf("%d/Account/%s/Summary/", a.MembershipType, a.MembershipID)
	var x accountSummaryJSON
	err := b.get(url, &x, false)
	if err != nil {
		return err
	}
	*a = x.Response.Data
	return nil
}

type characterInventoryJSON struct {
	jsonStatusFields
	Response struct {
		Data *Inventory `json:"data"`
	} `json:"Response"`
}

func (b *API) CharacterInventory(c *Character) (*Inventory, error) {
	url := fmt.Sprintf("%d/Account/%s/Character/%s/Inventory/Summary/",
		c.CharacterBase.MembershipType, c.CharacterBase.MembershipID,
		c.CharacterBase.CharacterID)
	var x characterInventoryJSON
	err := b.get(url, &x, false)
	if err != nil {
		return nil, err
	}
	return x.Response.Data, nil
}

type vaultSummaryJSON struct {
	jsonStatusFields
	Response struct {
		Data *Inventory `json:"data"`
	} `json:"Response"`
}

func (b *API) GetVaultContents(account *Account) (*Inventory, error) {
	if b.xcsrf == "" || b.cookie == "" {
		return nil, errors.New("Can't get vault contents without cookie and xcsrf")
	}
	url := fmt.Sprintf("%d/MyAccount/Vault/Summary/", account.MembershipType)
	var x vaultSummaryJSON
	err := b.get(url, &x, false)
	if err != nil {
		return nil, err
	}
	return x.Response.Data, nil
}
