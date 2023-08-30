package main

import (
	"github.com/lyft/clutch/backend/gateway"

	"github.com/DanielYevelkin/clutch-custom-gateway/backend/cmd/assets"
	"github.com/DanielYevelkin/clutch-custom-gateway/backend/module/echo"
	amiibomod "github.com/DanielYevelkin/clutch-custom-gateway/backend/module/amiibo"
    amiiboservice "github.com/DanielYevelkin/clutch-custom-gateway/backend/service/amiibo"
	vmapimod "github.com/DanielYevelkin/clutch-custom-gateway/backend/module/vmapi"
    vmapiservice "github.com/DanielYevelkin/clutch-custom-gateway/backend/service/vmapi"
)

func main() {
	flags := gateway.ParseFlags()

	components := gateway.CoreComponentFactory
	components.Modules[amiibomod.Name] = amiibomod.New
    components.Services[amiiboservice.Name] = amiiboservice.New
	components.Modules[vmapimod.Name] = vmapimod.New
    components.Services[vmapiservice.Name] = vmapiservice.New

	// Add custom components.
	components.Modules[echo.Name] = echo.New

	gateway.Run(flags, components, assets.VirtualFS)
}
