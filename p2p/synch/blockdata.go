/*
 * Copyright (c) 2017-2020 The qitmeer developers
 */

package synch

import (
	"github.com/Qitmeer/qng-core/common/hash"
	"github.com/Qitmeer/qng-core/core/types"
)

type BlockData struct {
	Hash *hash.Hash
	Block *types.SerializedBlock
}