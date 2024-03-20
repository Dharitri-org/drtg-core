package transaction

import "github.com/Dharitri-org/drtg-core/data/block"

// Marshalizer is able to encode an object to its byte slice representation
type Marshalizer interface {
	Marshal(obj interface{}) ([]byte, error)
	IsInterfaceNil() bool
}

// StatusComputerHandler computes a transaction status
type StatusComputerHandler interface {
	ComputeStatusWhenInStorageKnowingMiniblock(miniblockType block.Type, tx *ApiTransactionResult) (TxStatus, error)
	ComputeStatusWhenInStorageNotKnowingMiniblock(destinationShard uint32, tx *ApiTransactionResult) (TxStatus, error)
	SetStatusIfIsRewardReverted(
		tx *ApiTransactionResult,
		miniblockType block.Type,
		headerNonce uint64,
		headerHash []byte,
	) (bool, error)
}
