package model

import (
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"time"
)

type MsgCreateClawbackVestingAccountParams struct {
	RawMsgCreateClawbackVestingAccount
}

type RawMsgCreateClawbackVestingAccount struct {
	FromAddress    string          `mapstructure:"from_address" json:"from_address"`
	ToAddress      string          `mapstructure:"to_address" json:"to_address"`
	StartTime      time.Time       `mapstructure:"start_time" json:"start_time"`
	LockupPeriods  vesting.Periods `mapstructure:"lockup_periods" json:"lockup_periods"`
	VestingPeriods vesting.Periods `mapstructure:"vesting_periods" json:"vesting_periods"`
	Merge          bool            `mapstructure:"merge" json:"merge"`
}
