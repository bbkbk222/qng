package chainvm

import (
	"context"
	"github.com/Qitmeer/qng/vm/chainvm/proto"
	"github.com/Qitmeer/qng/vm/common"
	"github.com/hashicorp/go-plugin"
)

type VMServer struct {
	proto.UnimplementedVMServer
	vm     common.ChainVM
	broker *plugin.GRPCBroker

	ctx    *context.Context
	closed chan struct{}
}

func NewServer(vm common.ChainVM, broker *plugin.GRPCBroker) *VMServer {
	return &VMServer{
		vm:     vm,
		broker: broker,
	}
}
