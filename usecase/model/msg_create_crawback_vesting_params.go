package model

import (
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"time"
)

type MsgCreateClawbackVestingAccountParams struct {
	FromAddress    string          `json:"from_address"`
	ToAddress      string          `json:"to_address"`
	StartTime      time.Time       `json:"start_time"`
	LockupPeriods  vesting.Periods `json:"lockup_periods"`
	VestingPeriods vesting.Periods `json:"vesting_periods"`
	Merge          bool            `json:"merge"`
}
