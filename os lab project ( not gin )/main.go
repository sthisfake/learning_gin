package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"os/initialprint"
	"os/commands"
	"os/scheduling"
	"os/models"
	"strconv"
)

var currentDir string


func main() {

	currentDir, _ = os.Getwd()

	scanner := bufio.NewScanner(os.Stdin)

	processes := []*models.Process{
		{ID: 1, Name: "Process A", RequiredCPU: 50, RequiredRAM: 2, RequiredBW: 5},
		{ID: 2, Name: "Process B", RequiredCPU: 30, RequiredRAM: 1, RequiredBW: 2},
		{ID: 3, Name: "Process C", RequiredCPU: 50, RequiredRAM: 2, RequiredBW: 5},
		{ID: 4, Name: "Process D", RequiredCPU: 30, RequiredRAM: 1, RequiredBW: 2},
		{ID: 5, Name: "Process E", RequiredCPU: 50, RequiredRAM: 2, RequiredBW: 5},
		{ID: 6, Name: "Process F", RequiredCPU: 30, RequiredRAM: 1, RequiredBW: 2},
		{ID: 7, Name: "Process G", RequiredCPU: 50, RequiredRAM: 2, RequiredBW: 5},
		{ID: 8, Name: "Process H", RequiredCPU: 30, RequiredRAM: 1, RequiredBW: 2},
	}

	cpuAvailable := 100 // You can set the initial values here
	ramAvailable := 6
	bwAvailable := 10

	var newProcessCh chan *models.Process


	go scheduling.FCFS(&processes, &cpuAvailable, &ramAvailable, &bwAvailable , newProcessCh)
	

	for {

		initialprint.PrintPrompt(currentDir)

		scanner.Scan()
		input := scanner.Text()

		if input == "" {
			continue
		}

		if input == "exit" {
			fmt.Println("Exiting Terminal")
			break
		}

		// Split the input by spaces to separate the command and arguments
		args := strings.Split(input, " ")
		cmd := args[0]
		cmdArgs := args[1:]

		switch cmd {
		case "ls":
			commands.ListFiles(cmdArgs)
		case "cd":
			newDir:= commands.ChangeDirectory(cmdArgs , currentDir)
			currentDir = newDir
		case "process_check": // New command for checking currently running processes
			checkRunningProcesses(processes)	
	    case "pmanager":
			processID, command := parseProcessManagerCommand(cmdArgs)
			commands.ProcessManager(processes, processID, command , &cpuAvailable, &ramAvailable, &bwAvailable)	
		case "vpn":
			commands.CreateVPNProcess(&processes , &cpuAvailable, &ramAvailable, &bwAvailable , newProcessCh)	
		case "all_process":
			commands.ShowAllProcesses(processes)				
		default:
			fmt.Println("Command not supported:", cmd)
		}
	}
}

func checkRunningProcesses(processes []*models.Process) {
	if len(processes) > 0 {
		fmt.Println("Currently running processes:")
		for _, process := range processes {
			if process.IsResourceFulfilled && !process.IsCompleted {
				fmt.Printf("Process ID: %d, Name: %s\n", process.ID, process.Name)
				fmt.Printf("   CPU Usage: %d%%, RAM Usage: %d GB, Bandwidth Usage: %d Mbps\n", process.RequiredCPU, process.RequiredRAM, process.RequiredBW)
			}
		}
	}else{
		fmt.Println("No Process running at the momment")
	}

}


func parseProcessManagerCommand(args []string) (int, string) {
	if len(args) < 2 {
		fmt.Println("Invalid 'pmanager' command. Usage: pmanager <process_id> <command>")
		return 0, ""
	}

	processID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid process ID:", args[0])
		return 0, ""
	}

	return processID, args[1]
}








