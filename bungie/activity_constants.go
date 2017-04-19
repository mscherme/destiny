package bungie

const (
	VaultOfGlassNormal         int64 = 2659248071
	VaultOfGlassHeroic         int64 = 2659248068
	VaultOfGlassAgeOfTriumphCM int64 = 856898338
	VaultOfGlassAgeOfTriumph   int64 = 4038697181

	CrotasEndNormal         int64 = 1836893116
	CrotasEndHeroic         int64 = 1836893119
	CrotasEndAgeOfTriumphCM int64 = 4000873610
	CrotasEndAgeOfTriumph   int64 = 2324706853

	KingsFallNormal         int64 = 1733556769
	KingsFallHeroic         int64 = 3534581229
	KingsFallAgeOfTriumphCM int64 = 3978884648
	KingsFallAgeOfTriumph   int64 = 1016659723

	WrathOfTheMachineNormal         int64 = 260765522
	WrathOfTheMachineHeroic         int64 = 1387993552
	WrathOfTheMachineAgeOfTriumphCM int64 = 3356249023
	WrathOfTheMachineAgeOfTriumph   int64 = 0
)

const (
	NormalRaid int = iota
	HeroicRaid
	AgeOfTriumph
	AgeOfTriumphCM
)

type RaidDetails struct {
	LightLevel int
	Difficulty int
}

func (rd *RaidDetails) Heroic() bool {
	return rd.Difficulty != NormalRaid
}

var Raids = map[int64]*RaidDetails{
	VaultOfGlassNormal:         {130, NormalRaid},
	VaultOfGlassHeroic:         {150, HeroicRaid},
	VaultOfGlassAgeOfTriumphCM: {390, AgeOfTriumphCM},
	VaultOfGlassAgeOfTriumph:   {390, AgeOfTriumph},

	CrotasEndNormal:         {150, NormalRaid},
	CrotasEndHeroic:         {160, HeroicRaid},
	CrotasEndAgeOfTriumphCM: {390, AgeOfTriumphCM},
	CrotasEndAgeOfTriumph:   {390, AgeOfTriumph},

	KingsFallNormal:         {290, NormalRaid},
	KingsFallHeroic:         {310, HeroicRaid},
	KingsFallAgeOfTriumphCM: {390, AgeOfTriumphCM},
	KingsFallAgeOfTriumph:   {390, AgeOfTriumph},

	WrathOfTheMachineNormal:         {370, NormalRaid},
	WrathOfTheMachineHeroic:         {390, HeroicRaid},
	WrathOfTheMachineAgeOfTriumphCM: {390, AgeOfTriumphCM},
	WrathOfTheMachineAgeOfTriumph:   {390, AgeOfTriumph},
}
