package consensus

import (
	"github.com/Qitmeer/qng-core/consensus"
	"github.com/Qitmeer/qng-core/core/types"
)

type VMI interface {
	VerifyTx(tx consensus.Tx) (int64, error)
	GetVM(id string) (consensus.ChainVM, error)
	ConnectBlock(block *types.SerializedBlock) error
	DisconnectBlock(block *types.SerializedBlock) error
}
