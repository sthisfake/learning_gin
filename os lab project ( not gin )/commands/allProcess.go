package commands

import (
	"fmt"
	"os/models"
)

func ShowAllProcesses(runningProcesses []*models.Process){
	if len(runningProcesses) == 0 {
		fmt.Println("No processes running at the moment.")
		return
	}

	fmt.Println("Currently running processes:")
	for _, process := range runningProcesses {
		fmt.Printf("Process ID: %d, Name: %s\n", process.ID, process.Name)
		fmt.Printf("   CPU Usage: %d%%, RAM Usage: %d, Bandwidth Usage: %d\n", process.RequiredCPU, process.RequiredRAM, process.RequiredBW)
	}
}