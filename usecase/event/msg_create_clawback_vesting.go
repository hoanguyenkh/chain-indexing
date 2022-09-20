package event

import (
	"bytes"
	vesting "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/crypto-com/chain-indexing/usecase/model"
	"time"

	entity_event "github.com/crypto-com/chain-indexing/entity/event"
	jsoniter "github.com/json-iterator/go"
	"github.com/luci/go-render/render"
)

const MSG_CREATE_CRAW_BACK_VESTING_ACCOUNT = "/evmos.vesting.v1.MsgCreateClawbackVestingAccount"
const MSG_CREATE_CRAW_BACK_VESTING_ACCOUNT_CREATED = "/evmos.vesting.v1.MsgCreateClawbackVestingAccount.Created"
const MSG_CREATE_CRAW_BACK_VESTING_ACCOUNT_FAILED = "/evmos.vesting.v1.MsgCreateClawbackVestingAccount.Failed"

type MsgCreateClawbackVestingAccount struct {
	MsgBase
	FromAddress string `json:"from_address"`
	// to_address specifies the account to receive the funds
	ToAddress string `json:"to_address"`
	// start_time defines the time at which the vesting period begins
	StartTime      time.Time       `json:"start_time"`
	LockupPeriods  vesting.Periods `json:"lockup_periods"`
	VestingPeriods vesting.Periods `json:"vesting_periods"`

	Merge bool `json:"merge"`
}

func NewMsgCreateClawbackVestingAccount(msgCommonParams MsgCommonParams,
	params model.MsgCreateClawbackVestingAccountParams) *MsgCreateClawbackVestingAccount {
	return &MsgCreateClawbackVestingAccount{
		NewMsgBase(MsgBaseParams{
			MsgName:         MSG_CREATE_CRAW_BACK_VESTING_ACCOUNT,
			Version:         1,
			MsgCommonParams: msgCommonParams,
		}),
		params.FromAddress,
		params.ToAddress,
		params.StartTime,
		params.LockupPeriods,
		params.VestingPeriods,
		params.Merge,
	}
}

// ToJSON encodes the event into JSON string payload
func (event *MsgCreateClawbackVestingAccount) ToJSON() (string, error) {
	encoded, err := jsoniter.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(encoded), nil
}

func (event *MsgCreateClawbackVestingAccount) String() string {
	return render.Render(event)
}

func DecodeMsgCreateClawbackVestingAccount(encoded []byte) (entity_event.Event, error) {
	jsonDecoder := jsoniter.NewDecoder(bytes.NewReader(encoded))
	jsonDecoder.DisallowUnknownFields()

	var event *MsgCreateClawbackVestingAccount
	if err := jsonDecoder.Decode(&event); err != nil {
		return nil, err
	}

	return event, nil
}
