package linuxwifi

import (
	"os"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func platformReadings(path string) (map[string]interface{}, error) {
	dump, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	lines := strings.Split(strings.TrimSpace(string(dump)), "\n")
	for i, line := range lines {
		if i < 2 {
			continue
		}
		iface, readings, err := readingsByInterface(line)
		if err != nil {
			return nil, err
		}
		result[iface] = readings
	}

	return result, nil
}

func readingsByInterface(line string) (string, map[string]interface{}, error) {
	fields := strings.Fields(line)

	iface := strings.TrimRight(fields[0], ":")

	link, err := strconv.ParseInt(strings.TrimRight(fields[2], "."), 10, 32)
	if err != nil {
		return "", nil, errors.Wrap(err, "invalid link quality reading")
	}
	level, err := strconv.ParseInt(strings.TrimRight(fields[3], "."), 10, 32)
	if err != nil {
		return "", nil, errors.Wrap(err, "invalid wifi level reading")
	}
	noise, err := strconv.ParseInt(fields[4], 10, 32)
	if err != nil {
		return "", nil, errors.Wrap(err, "invalid wifi noise reading")
	}

	return iface, map[string]interface{}{
		"link_quality": int(link),
		"level_dBm":    int(level),
		"noise_dBm":    int(noise),
	}, nil
}
