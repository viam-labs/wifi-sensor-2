// package main is a module with a linux wifi sensor component
package main

import (
	"context"

	"github.com/edaniels/golog"
	"github.com/viam-labs/wifi-sensor/linuxwifi"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, golog.NewDevelopmentLogger("wifi-sensor"))
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) error {
	wifiSensorModule, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	wifiSensorModule.AddModelFromRegistry(ctx, sensor.API, linuxwifi.Model)

	err = wifiSensorModule.Start(ctx)
	defer wifiSensorModule.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
