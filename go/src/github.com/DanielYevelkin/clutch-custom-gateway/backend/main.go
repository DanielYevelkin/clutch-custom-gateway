package main

import (
	"github.com/lyft/clutch/backend/gateway"

	"github.com/DanielYevelkin/clutch-custom-gateway/backend/cmd/assets"
	"github.com/DanielYevelkin/clutch-custom-gateway/backend/module/echo"
)

func main() {
	flags := gateway.ParseFlags()

	components := gateway.CoreComponentFactory

	// Add custom components.
	components.Modules[echo.Name] = echo.New

	gateway.Run(flags, components, assets.VirtualFS)
}
