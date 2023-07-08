package commands

import (
	"os/models"
)

// ...

func CreateVPNProcess(runningProcesses *[]*models.Process,cpuAvailable *int, ramAvailable *int, bwAvailable *int , newProcessCh chan *models.Process) {

	vpnProcess := models.NewVPNProcess(len(*runningProcesses) + 1)
	*runningProcesses = append(*runningProcesses, vpnProcess)

}






