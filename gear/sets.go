package gear

import (
	"strings"

	"github.com/mscherme/destiny/bungie"
)

type Set int64

const (
	VaultOfGlass         Set = 686593720
	CrotasEnd            Set = 3107502809
	PrisonOfElders       Set = 36493462
	KingsFall            Set = 3551688287
	ChallengeOfElders    Set = 3739898362
	WrathOfTheMachine    Set = 3147905712
	RiseOfIronRecordBook Set = 346792680
	TrialsOfOsirisY3     Set = 3413298620
	IronBannerY3         Set = 478645002
	IronBannerY2         Set = 1
)

var setNames = map[Set]string{
	VaultOfGlass:   "Vault of Glass",
	CrotasEnd:      "Crota's End",
	PrisonOfElders: "Prison of Elders",

	KingsFall:         "King's Fall",
	ChallengeOfElders: "Challenge of the Elders",

	WrathOfTheMachine:    "Wrath of the Machine",
	RiseOfIronRecordBook: "Rise of Iron Record Book",

	TrialsOfOsirisY3: "Trials of Osiris Year 3",
	IronBannerY3:     "Iron Banner Year 3",
	IronBannerY2:     "Iron Banner Year 2",
}

var gearSetOverrides = map[int64]Set{
	3497087277: IronBannerY2,
	2897238917: IronBannerY2,
	904536202:  IronBannerY2,
	3452625744: IronBannerY2,
	2443083323: IronBannerY2,
	3068424913: IronBannerY2,
	1272989272: IronBannerY2,
	1264422556: IronBannerY2,
	3536592559: IronBannerY2,
	2121113047: IronBannerY2,
	2747259661: IronBannerY2,
	3399822332: IronBannerY2,
	554244708:  IronBannerY2,
	2402750214: IronBannerY2,
	3907475891: IronBannerY2,
	1866413378: IronBannerY2,
	82944038:   IronBannerY2,
	3910559228: IronBannerY2,
	870077908:  IronBannerY2,
	3995873645: IronBannerY2,
	1212068371: IronBannerY2,
	4242215215: IronBannerY2,
	825392783:  IronBannerY2,
	1722532281: IronBannerY2,
	3841951818: IronBannerY2,
}

func (g Set) String() string {
	return setNames[g]
}

func SetForItem(info *bungie.InventoryItem) Set {
	if strings.Contains(info.ItemName, "Engram") {
		return 0
	}
	if s := gearSetOverrides[info.ItemHash]; s != 0 {
		return s
	}
	if info.TierType == 5 {
		for _, source := range info.SourceHashes {
			if _, ok := setNames[Set(source)]; ok {
				return Set(source)
			}
		}
	}
	return 0
}
