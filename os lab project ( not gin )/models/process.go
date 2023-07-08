package models

type Process struct {
	ID                  int
	Name                string
	RequiredCPU         int // CPU required by the process (%)
	RequiredRAM         int // RAM required by the process (GB)
	RequiredBW          int // Network bandwidth required by the process (Mbps)
	IsCompleted         bool
	IsResourceFulfilled bool
	TerminateCh         chan struct{}
}

func NewVPNProcess(id int) *Process {
	return &Process{
		ID:          id,
		Name:        "VPN",
		RequiredBW:  2,
		TerminateCh: make(chan struct{}),
	}
}