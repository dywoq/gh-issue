package process

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/dywoq/dywoqlib/err"
)

type config struct {
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
	Token      string `json:"token"`
}

func newConfig(filename string) (config, err.Context) {
	gotData, err1 := os.ReadFile(filename)
	if err1 != nil {
		return config{}, err.NewContext(err1, "source is process.newConfig(string) (config, err.Context)")
	}
	var (
		d    = json.NewDecoder(strings.NewReader(string(gotData)))
		conf config
	)
	err1 = d.Decode(&conf)
	if err1 != nil {
		return config{}, err.NewContext(err1, "source is process.newConfig(string) (config, err.Context)")
	}
	return conf, err.NoneContext()
}
