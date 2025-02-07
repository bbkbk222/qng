package main

import (
	"github.com/Qitmeer/qng/core/blockchain"
	"github.com/Qitmeer/qng-core/database"
)

type INode interface {
	BlockChain() *blockchain.BlockChain
	DB() database.DB
}
