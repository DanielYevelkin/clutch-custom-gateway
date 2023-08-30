package vm

import (
	"context"
	"errors"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"

	vmapi "github.com/DanielYevelkin/clutch-custom-gateway/backend/api/vmapi/v1" // Replace with the actual path to your .proto generated Go file
	"github.com/lyft/clutch/backend/module"
	"github.com/lyft/clutch/backend/service"
	vmservice "github.com/DanielYevelkin/clutch-custom-gateway/backend/service/vmapi" // Replace with the actual path to your VM service
)

const Name = "gateway.module.vmapi"

func New(_ *any.Any, _ *zap.Logger, _ tally.Scope) (module.Module, error) {
	svc, ok := service.Registry["gateway.service.vmapi"]
	if !ok {
		return nil, errors.New("no VM service was registered")
	}

	client, ok := svc.(vmservice.Client)
	if !ok {
		return nil, errors.New("VM service in registry was the wrong type")
	}

	return &mod{client: client}, nil
}

type mod struct {
	client vmservice.Client
}

func (m *mod) Register(r module.Registrar) error {
	vmapi.RegisterVirtualMachineServiceServer(r.GRPCServer(), m)
	return r.RegisterJSONGateway(vmapi.RegisterVirtualMachineServiceHandler)
}

func (m *mod) GetVMs(ctx context.Context, request *vmapi.GetVMsRequest) (*vmapi.GetVMsResponse, error) {
	vms, err := m.client.GetVMs(ctx)
	if err != nil {
		return nil, err
	}
	return &vmapi.GetVMsResponse{Vms: vms}, nil
}

func (m *mod) AddVM(ctx context.Context, request *vmapi.AddVMRequest) (*vmapi.AddVMResponse, error) {
	vm, err := m.client.AddVM(ctx, request.Vm)
	if err != nil {
		return nil, err
	}
	return &vmapi.AddVMResponse{Vm: vm}, nil
}
