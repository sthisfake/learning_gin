package scheduling

import (
	"fmt"
	"os/models"
	"time"
)

func runProcess(processes *[]*models.Process, process *models.Process, cpuAvailable *int, ramAvailable *int, bwAvailable *int, resourceCh chan<- struct{}) {
	// process start
	process.IsResourceFulfilled = true
	process.IsCompleted = false

	time.Sleep(5 * time.Second)

	// Process finishes
	*cpuAvailable += process.RequiredCPU
	*ramAvailable += process.RequiredRAM
	*bwAvailable += process.RequiredBW

	process.IsResourceFulfilled = false
	process.IsCompleted = true

	// Remove the completed process from the list
	for i, p := range *processes {
		if p == process {
			copy((*processes)[i:], (*processes)[i+1:])
			*processes = (*processes)[:len(*processes)-1]
			break
		}
	}

	// Signal on the channel that resources are available
	resourceCh <- struct{}{}
}

func runVpnProcess(processes *[]*models.Process, process *models.Process, cpuAvailable *int, ramAvailable *int, bwAvailable *int, resourceCh chan<- struct{}) {
	// Process start
	process.IsResourceFulfilled = true
	process.IsCompleted = false

	for {
		select {
		case <-process.TerminateCh:
			fmt.Printf("VPN process  terminated.\n")
			process.IsCompleted = true

			// Remove the completed process from the list
			for i, p := range *processes {
				if p == process {
					copy((*processes)[i:], (*processes)[i+1:])
					*processes = (*processes)[:len(*processes)-1]
					break
				}
			}

			*cpuAvailable += process.RequiredCPU
			*ramAvailable += process.RequiredRAM
			*bwAvailable += process.RequiredBW
			resourceCh <- struct{}{}
			return
		}
	}
}
