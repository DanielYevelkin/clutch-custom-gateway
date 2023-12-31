package echo

import (
	"context"
	"testing"

	"github.com/lyft/clutch/backend/module/moduletest"
	"github.com/stretchr/testify/assert"
	"github.com/uber-go/tally/v4"
	"go.uber.org/zap/zaptest"

	echov1 "github.com/DanielYevelkin/clutch-custom-gateway/backend/api/echo/v1"
)

func TestModule(t *testing.T) {
	log := zaptest.NewLogger(t)
	scope := tally.NewTestScope("", nil)
	m, err := New(nil, log, scope)

	assert.NoError(t, err)

	r := moduletest.NewRegisterChecker()
	assert.NoError(t, m.Register(r))
	assert.NoError(t, r.HasAPI("DanielYevelkin.echo.v1.EchoAPI"))
	assert.True(t, r.JSONRegistered())
}

func TestEchoSayHello(t *testing.T) {
	api := newAPI()
	resp, err := api.SayHello(context.Background(), &echov1.SayHelloRequest{Name: "world"})
	assert.NoError(t, err)
	assert.Equal(t, "Hello world!", resp.Message)
}
