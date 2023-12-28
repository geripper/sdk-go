package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CommitmentTimestampResponse struct {
	Timestamp uint64 `json:"timestamp"`
}

type GetCommitmentResponse struct {
	Commitment string `json:"commitment"`
}

type TokenIdResponse struct {
	TokenId string `json:"token_id"`
}

type IsValidNameResponse struct {
	IsValidName bool `json:"is_valid_name"`
}

type MaxCommitmentAgeResponse struct {
	Age uint64 `json:"age"`
}

type MinCommitmentAgeResponse struct {
	Age uint64 `json:"age"`
}

type NodeInfoResponse struct {
	Label   []uint `json:"label"`
	Node    []uint `json:"node"`
	TokenId string `json:"token_id"`
}

type MinRegistrationDurationResponse struct {
	Duration uint64 `json:"duration"`
}

type NodeHashResponse struct {
	Node []uint `json:"node"`
}

type OwnerResponse struct {
	Owner string `json:"owner"`
}

type PriceResponse struct {
	Tier1Price uint64 `json:"tier1_price"`
	Tier2Price uint64 `json:"tier2_price"`
	Tier3Price uint64 `json:"tier3_price"`
}

type RegistrarResponse struct {
	RegistrarAddress string `json:"registrar_address"`
}

type RentPriceResponse struct {
	Price sdk.Uint `json:"price"`
}

type PaymentChanResponse struct {
	Id             uint32 `json:"id"`
	MaxAmount      string `json:"max_amount"`
	WithdrawlSeq   uint64 `json:"withdrawl_seq"`
	FaceValue      string `json:"face_value"`
	AutoRelease    uint64 `json:"auto_release"`
	ApproveSigners string `json:"approve_signers"`
}

type StakedResponse struct {
	Denom map[string]string `json:"denom"`
	Stake StakeData         `json:"stake"`
}

type StakeData struct {
	Amount   string `json:"amount"`
	Metadata []uint `json:"metadata"`
	Status   string `json:"status"`
}
