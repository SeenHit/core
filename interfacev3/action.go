package interfacev3

import (
	"github.com/hacash/core/fields"
	"github.com/hacash/core/interfaces"
)

type Action interface {

	// base super
	interfaces.Field

	// the action type number
	Kind() uint16

	// Addresses that need to verify signatures
	RequestSignAddresses() []fields.Address

	// change chain state
	WriteInChainState(ChainStateOperation) error

	// help func
	SetBelongTrs(Transaction)
	Describe() map[string]interface{} // json api

	// burning fees
	IsBurning90PersentTxFees() bool // 是否销毁本笔交易的 90% 的交易费用
}
