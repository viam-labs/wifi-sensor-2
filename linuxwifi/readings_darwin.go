package linuxwifi

import (
	"os/exec"
	"strings"
)

const airportCli string = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport"

func platformReadings(_ string) (map[string]interface{}, error) {
	out, err := exec.Command(airportCli, "-I").Output()
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{})
	for _, line := range strings.Split(string(out), "\n") {
		if before, after, found := strings.Cut(strings.Trim(line, " "), ": "); found {
			ret[before] = after
		}
	}
	return ret, nil
}
