package bungie

import "fmt"

type manifestJSON struct {
	jsonStatusFields
	Response struct {
		Data struct {
			RequestID     int64
			Activity      *Activity `json:"activity"`
			InventoryItem *InventoryItem
		}
	}
}

type InventoryItem struct {
	ItemHash            int64  `json:"itemHash"`
	ItemName            string `json:"itemName"`
	ItemDescription     string `json:"itemDescription"`
	Icon                string `json:"icon"`
	HasIcon             bool   `json:"hasIcon"`
	SecondaryIcon       string `json:"secondaryIcon"`
	ActionName          string `json:"actionName"`
	HasAction           bool   `json:"hasAction"`
	DeleteOnAction      bool   `json:"deleteOnAction"`
	TierTypeName        string `json:"tierTypeName"`
	TierType            int64  `json:"tierType"`
	ItemTypeName        string `json:"itemTypeName"`
	BucketTypeHash      int64  `json:"bucketTypeHash"`
	PrimaryBaseStatHash int64  `json:"primaryBaseStatHash"`
	Stats               map[string]struct {
		StatHash int64 `json:"statHash"`
		Value    int64 `json:"value"`
		Minimum  int64 `json:"minimum"`
		Maximum  int64 `json:"maximum"`
	} `json:"stats"`
	PerkHashes      []int64 `json:"perkHashes"`
	SpecialItemType int64   `json:"specialItemType"`
	TalentGridHash  int64   `json:"talentGridHash"`
	EquippingBlock  struct {
		WeaponSandboxPatternIndex int64 `json:"weaponSandboxPatternIndex"`
		GearArtArrangementIndex   int64 `json:"gearArtArrangementIndex"`
		//DefaultDyes               []interface{} `json:"defaultDyes"`
		LockedDyes []struct {
			ChannelHash int64 `json:"channelHash"`
			DyeHash     int64 `json:"dyeHash"`
		} `json:"lockedDyes"`
		//CustomDyes          []interface{} `json:"customDyes"`
		CustomDyeExpression struct {
			//Steps []interface{} `json:"steps"`
		} `json:"customDyeExpression"`
		WeaponPatternHash int64 `json:"weaponPatternHash"`
		Arrangements      []struct {
			ClassHash               int64 `json:"classHash"`
			GearArtArrangementIndex int64 `json:"gearArtArrangementIndex"`
		} `json:"arrangements"`
		EquipmentSlotHash int64 `json:"equipmentSlotHash"`
	} `json:"equippingBlock"`
	HasGeometry    bool    `json:"hasGeometry"`
	StatGroupHash  int64   `json:"statGroupHash"`
	ItemLevels     []int64 `json:"itemLevels"`
	QualityLevel   int64   `json:"qualityLevel"`
	Equippable     bool    `json:"equippable"`
	Instanced      bool    `json:"instanced"`
	RewardItemHash int64   `json:"rewardItemHash"`
	// TODO
	//Values         struct {
	//} `json:"values"`
	ItemType    int64 `json:"itemType"`
	ItemSubType int64 `json:"itemSubType"`
	ClassType   int64 `json:"classType"`
	Sources     []struct {
		ExpansionIndex   int64 `json:"expansionIndex"`
		Level            int64 `json:"level"`
		MinQuality       int64 `json:"minQuality"`
		MaxQuality       int64 `json:"maxQuality"`
		MinLevelRequired int64 `json:"minLevelRequired"`
		MaxLevelRequired int64 `json:"maxLevelRequired"`
		Exclusivity      int64 `json:"exclusivity"`
		ComputedStats    map[string]struct {
			StatHash int64 `json:"statHash"`
			Value    int64 `json:"value"`
			Minimum  int64 `json:"minimum"`
			Maximum  int64 `json:"maximum"`
		} `json:"computedStats"`
		SourceHashes []int64 `json:"sourceHashes"`
	} `json:"sources"`
	ItemCategoryHashes []int64 `json:"itemCategoryHashes"`
	SourceHashes       []int64 `json:"sourceHashes"`
	NonTransferrable   bool    `json:"nonTransferrable"`
	Exclusive          int64   `json:"exclusive"`
	MaxStackSize       int64   `json:"maxStackSize"`
	ItemIndex          int64   `json:"itemIndex"`
	//SetItemHashes                []interface{} `json:"setItemHashes"`
	TooltipStyle        string `json:"tooltipStyle"`
	QuestlineItemHash   int64  `json:"questlineItemHash"`
	NeedsFullCompletion bool   `json:"needsFullCompletion"`
	//ObjectiveHashes              []interface{} `json:"objectiveHashes"`
	AllowActions                 bool    `json:"allowActions"`
	QuestTrackingUnlockValueHash int64   `json:"questTrackingUnlockValueHash"`
	BountyResetUnlockHash        int64   `json:"bountyResetUnlockHash"`
	UniquenessHash               int64   `json:"uniquenessHash"`
	ShowActiveNodesInTooltip     bool    `json:"showActiveNodesInTooltip"`
	DamageTypes                  []int64 `json:"damageTypes"`
	Hash                         int64   `json:"hash"`
	Index                        int64   `json:"index"`
	Redacted                     bool    `json:"redacted"`
}

type Activity struct {
	ActivityHash        int64   `json:"activityHash"`
	ActivityName        string  `json:"activityName"`
	ActivityDescription string  `json:"activityDescription"`
	Icon                string  `json:"icon"`
	ReleaseIcon         string  `json:"releaseIcon"`
	ReleaseTime         int64   `json:"releaseTime"`
	ActivityLevel       int64   `json:"activityLevel"`
	CompletionFlagHash  int64   `json:"completionFlagHash"`
	ActivityPower       float64 `json:"activityPower"`
	MinParty            int64   `json:"minParty"`
	MaxParty            int64   `json:"maxParty"`
	MaxPlayers          int64   `json:"maxPlayers"`
	DestinationHash     int64   `json:"destinationHash"`
	PlaceHash           int64   `json:"placeHash"`
	ActivityTypeHash    int64   `json:"activityTypeHash"`
	Tier                int64   `json:"tier"`
	PgcrImage           string  `json:"pgcrImage"`
	//Rewards             []interface{} `json:"rewards"`
	//Skulls              []interface{} `json:"skulls"`
	IsPlaylist  bool  `json:"isPlaylist"`
	IsMatchmade bool  `json:"isMatchmade"`
	Hash        int64 `json:"hash"`
	Index       int64 `json:"index"`
	Redacted    bool  `json:"redacted"`
}

func (b *API) ManifestInventoryItem(rID int64) (*InventoryItem, error) {
	url := fmt.Sprintf("Manifest/InventoryItem/%d/", rID)
	var x manifestJSON
	err := b.get(url, &x, true)
	if err != nil {
		return nil, err
	}
	return x.Response.Data.InventoryItem, nil
}

func (b *API) ManifestActivity(rID int64) (*Activity, error) {
	url := fmt.Sprintf("Manifest/Activity/%d/", rID)
	var x manifestJSON
	err := b.get(url, &x, true)
	if err != nil {
		return nil, err
	}
	return x.Response.Data.Activity, nil
}
