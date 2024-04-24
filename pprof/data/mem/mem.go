package mem

import (
	"CloudNative/pprof/constants"
	"log"
)

type Mem struct {
	buffer [][constants.MI]byte
}

func (m *Mem) Name() string {
	return "memory"
}
func (m *Mem) Run() {
	log.Println(m.Name(), "Run")
	for len(m.buffer)*constants.MI < constants.GI {
		m.buffer = append(m.buffer, [constants.MI]byte{})
	}
}
