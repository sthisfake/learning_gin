package commands

import (
	"fmt"
	"os/models"
)

func ProcessManager(processes []*models.Process, processID int, command string, cpuAvailable *int, ramAvailable *int, bwAvailable *int) {
	for i, process := range processes {
		if process.ID == processID {
			switch command {
			case "kill":
				killProcess(processes, i, process, cpuAvailable, ramAvailable, bwAvailable)
			default:
				fmt.Println("Invalid command:", command)
			}
			return
		}
	}

	fmt.Println("Process not found with ID:", processID)
}

func killProcess(processes []*models.Process, index int, process *models.Process, cpuAvailable *int, ramAvailable *int, bwAvailable *int) {

	if process.IsCompleted {
		fmt.Printf("Process %d (%s) is not running \n", process.ID, process.Name)
	} else {
		fmt.Printf("Process %d (%s) has been killed and resources returned.\n", process.ID, process.Name)

		*cpuAvailable += process.RequiredCPU
		*ramAvailable += process.RequiredRAM
		*bwAvailable += process.RequiredBW

		process.IsResourceFulfilled = false
		process.IsCompleted = true

		if process.Name == "VPN" {
			close(process.TerminateCh)
		}

		// Remove the killed process from the list of processes
		copy(processes[index:], processes[index+1:])
		processes = processes[:len(processes)-1]
	}
}
