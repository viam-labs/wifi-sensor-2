// Package linuxwifi implements a wifi strength sensor
package linuxwifi

import (
	"context"
	"os"
	"path/filepath"

	"github.com/edaniels/golog"
	"github.com/pkg/errors"

	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/resource"
)

// Model represents a linux wifi strength sensor model.
var Model = resource.NewModel("viam", "sensor", "linux-wifi")

const wirelessInfoPath string = "/proc/net/wireless"

// stub config to satisfy resource.Registration
type StubConfig struct{}

func (cfg StubConfig) Validate(path string) ([]string, error) {
	return []string{}, nil
}

func init() {
	resource.RegisterComponent(
		sensor.API,
		Model,
		resource.Registration[sensor.Sensor, StubConfig]{
			Constructor: func(
				_ context.Context,
				_ resource.Dependencies,
				_ resource.Config,
				logger golog.Logger,
			) (sensor.Sensor, error) {
				return newWifi(logger, wirelessInfoPath)
			},
		},
	)
}

func newWifi(logger golog.Logger, path string) (sensor.Sensor, error) {
	if _, err := os.ReadFile(filepath.Clean(path)); err != nil {
		return nil, errors.Wrap(err, "wifi readings not supported on this system")
	}
	return &wifi{logger: logger, path: path}, nil
}

type wifi struct {
	resource.Named
	resource.TriviallyCloseable
	resource.TriviallyReconfigurable
	logger golog.Logger

	path string // for testing
}

// DoCommand always returns unimplemented but can be implemented by the embedder.
func (sensor *wifi) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	return nil, resource.ErrDoUnimplemented
}

// Readings returns Wifi strength statistics.
func (sensor *wifi) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
	return platformReadings(sensor.path)
}
