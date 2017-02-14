package gear

import (
	"strings"

	"destiny/bungie"
)

type Set int64

const (
	VaultOfGlass         Set = 686593720
	CrotasEnd                = 3107502809
	PrisonOfElders           = 36493462
	KingsFall                = 3551688287
	ChallengeOfElders        = 3739898362
	WrathOfTheMachine        = 3147905712
	RiseOfIronRecordBook     = 346792680
	TrialsOfOsiris           = 3413298620
	IronBanner               = 478645002
)

var setNames = map[Set]string{
	VaultOfGlass:   "Vault of Glass",
	CrotasEnd:      "Crota's End",
	PrisonOfElders: "Prison of Elders",

	KingsFall:         "King's Fall",
	ChallengeOfElders: "Challenge of the Elders",

	WrathOfTheMachine:    "Wrath of the Machine",
	RiseOfIronRecordBook: "Rise of Iron Record Book",

	TrialsOfOsiris: "Trials of Osiris",
	IronBanner:     "Iron Banner",
}

func (g Set) String() string {
	return setNames[g]
}

func SetForItem(info *bungie.InventoryItem) Set {
	if strings.Contains(info.ItemName, "Engram") {
		return 0
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
