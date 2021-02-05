// +build !linux

package cgroup

import (
	"github.com/pmalek-sumo/telegraf"
)

func (g *CGroup) Gather(acc telegraf.Accumulator) error {
	return nil
}
