package bungie

import "time"

type Basic struct {
	Value        float64 `json:"value"`
	DisplayValue string  `json:"displayValue"`
}

type Stat struct {
	Basic Basic `json:"basic"`
}

type StatWithID struct {
	StatID string `json:"statId"`
	Basic  Basic  `json:"basic"`
}

type Values struct {
	Assists                 StatWithID `json:"assists"`
	Deaths                  StatWithID `json:"deaths"`
	Kills                   StatWithID `json:"kills"`
	AverageScorePerKill     StatWithID `json:"averageScorePerKill"`
	AverageScorePerLife     StatWithID `json:"averateScorePerLife"`
	Completed               StatWithID `json:"completed"`
	KillsDeathsRatio        StatWithID `json:"killsDeathsRatio"`
	KillsDeathsAssists      StatWithID `json:"killsDeathsAssists"`
	ActivityDurationSeconds StatWithID `json:"activityDurationSeconds"`
	CompletionReason        StatWithID `json:"completionReason"`
	FireTeamID              StatWithID `json:"fireTeamId"`
	PlayerCount             StatWithID `json:"playerCount"`
	LeaveRemainingSeconds   StatWithID `json:"leaveRemainingSeconds"`
}

func (v *Values) Duration() time.Duration {
        return time.Duration(v.ActivityDurationSeconds.Basic.Value) * time.Second
}
