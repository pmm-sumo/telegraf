// +build !linux,!freebsd

package zfs

import (
	"github.com/pmalek-sumo/telegraf"
	"github.com/pmalek-sumo/telegraf/plugins/inputs"
)

func (z *Zfs) Gather(acc telegraf.Accumulator) error {
	return nil
}

func init() {
	inputs.Add("zfs", func() telegraf.Input {
		return &Zfs{}
	})
}
