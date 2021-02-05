package parallel

import "github.com/pmalek-sumo/telegraf"

type Parallel interface {
	Enqueue(telegraf.Metric)
	Stop()
}
