package speedtest

import (
	"fmt"
)

var (
	bits        = 8
	kb          = uint64(1024)
	mb          = 1024 * kb
	gb          = 1024 * mb
	tb          = 1024 * gb
	pb          = 1024 * tb
	tooDamnFast = "Too fast to test"
)

func HumanSpeed(bps uint64) string {
	if bps > pb {
		return tooDamnFast
	} else if bps > tb {
		return fmt.Sprintf("%.02f Tbps/s", float64(bps)/float64(tb))
	} else if bps > gb {
		return fmt.Sprintf("%.02f Gbps/s", float64(bps)/float64(gb))
	} else if bps > mb {
		return fmt.Sprintf("%.02f Mbps/s", float64(bps)/float64(mb))
	} else if bps > kb {
		return fmt.Sprintf("%.02f Kbps/s", float64(bps)/float64(kb))
	}
	return fmt.Sprintf("%d bps", bps)
}
