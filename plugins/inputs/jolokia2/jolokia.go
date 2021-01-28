package jolokia2

import (
	"github.com/pmalek-sumo/telegraf"
	"github.com/pmalek-sumo/telegraf/plugins/inputs"
)

func init() {
	inputs.Add("jolokia2_agent", func() telegraf.Input {
		return &JolokiaAgent{
			Metrics:               []MetricConfig{},
			DefaultFieldSeparator: ".",
		}
	})
	inputs.Add("jolokia2_proxy", func() telegraf.Input {
		return &JolokiaProxy{
			Metrics:               []MetricConfig{},
			DefaultFieldSeparator: ".",
		}
	})
}
