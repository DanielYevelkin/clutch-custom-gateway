package vmapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"bytes"
	"net/http"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"

	vmapi "github.com/DanielYevelkin/clutch-custom-gateway/backend/api/vmapi/v1" 
	"github.com/lyft/clutch/backend/service"
)

const Name = "gateway.service.vmapi"

const apiHost = "http://localhost:7070" // Replace with the actual host and port where your VM API is running

func New(cfg *any.Any, logger *zap.Logger, scope tally.Scope) (service.Service, error) {
	return &client{http: &http.Client{}}, nil
}

type Client interface {
	GetVMs(ctx context.Context) ([]*vmapi.VirtualMachine, error)
	AddVM(ctx context.Context, vm *vmapi.VirtualMachine) (*vmapi.VirtualMachine, error)
}

type client struct {
	http *http.Client
}

func vmsFromJSON(data []byte) ([]*vmapi.VirtualMachine, error) {
	var raw []vmapi.VirtualMachine
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, err
	}

	// Convert slice of structs to slice of pointers to structs
	ptrs := make([]*vmapi.VirtualMachine, len(raw))
	for i := 0; i < len(raw); i++ {
		ptrs[i] = &raw[i]
	}

	return ptrs, nil
}

func (c *client) GetVMs(ctx context.Context) ([]*vmapi.VirtualMachine, error) {
	url := fmt.Sprintf("%s/vms/get", apiHost)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, status.Error(service.CodeFromHTTPStatus(resp.StatusCode), string(body))
	}
	return vmsFromJSON(body)
}

func (c *client) AddVM(ctx context.Context, vm *vmapi.VirtualMachine) (*vmapi.VirtualMachine, error) {
	url := fmt.Sprintf("%s/vms/add", apiHost)
	payload, err := json.Marshal(vm)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, status.Error(service.CodeFromHTTPStatus(resp.StatusCode), string(body))
	}

	var newVM vmapi.VirtualMachine
	if err := json.Unmarshal(body, &newVM); err != nil {
		return nil, err
	}
	return &newVM, nil
}
