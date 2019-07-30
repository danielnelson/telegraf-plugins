package example

import (
	"os"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

const sampleConfig = `
  value = 42
`

type Example struct {
	Value int `toml:"value"`
}

func (e *Example) SampleConfig() string {
	return sampleConfig
}

func (e *Example) Description() string {
	return "Example go-plugin for Telegraf"
}

func (e *Example) Gather(acc telegraf.Accumulator) error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}

	acc.AddFields("example",
		map[string]interface{}{
			"value": e.Value,
		},
		map[string]string{
			"source": host,
		})
	return nil
}

func init() {
	inputs.Add("example", func() telegraf.Input {
		return &Example{}
	})
}
