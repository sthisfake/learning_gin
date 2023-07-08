package scheduling

import (
	"os/models"
	"time"
)

// func FCFS(processes []*models.Process, cpuAvailable *int, ramAvailable *int, bwAvailable *int) {
// 	resourceCh := make(chan struct{}, 1) // Channel for resource availability signal, buffer size 1 to avoid blocking

// 	for _, process := range processes {
// 		if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
// 			*cpuAvailable -= process.RequiredCPU
// 			*ramAvailable -= process.RequiredRAM
// 			*bwAvailable -= process.RequiredBW
// 			if process.Name == "VPN"{
// 				go runVpnProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh )
// 			}else{
// 				go runProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 			}

// 		} else {
// 			for {
// 				if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
// 					*cpuAvailable -= process.RequiredCPU
// 					*ramAvailable -= process.RequiredRAM
// 					*bwAvailable -= process.RequiredBW
// 					if process.Name == "VPN"{
// 						go runVpnProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh )
// 					}else{
// 						go runProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 					}
// 					break
// 				}
// 				time.Sleep(100 * time.Millisecond)
// 			}
// 		}
// 	}
// }

// func FCFS(processes *[]*models.Process, cpuAvailable *int, ramAvailable *int, bwAvailable *int, newProcessCh chan *models.Process) {
// 	resourceCh := make(chan struct{}, 1) // Channel for resource availability signal, buffer size 1 to avoid blocking

// 		for {
// 			select {
// 			case process := <-newProcessCh:
// 				if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
// 					*cpuAvailable -= process.RequiredCPU
// 					*ramAvailable -= process.RequiredRAM
// 					*bwAvailable -= process.RequiredBW
// 					if process.Name == "VPN" {
// 						go runVpnProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 					} else {
// 						go runProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 					}
// 				} else {
// 					for {
// 						if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
// 							*cpuAvailable -= process.RequiredCPU
// 							*ramAvailable -= process.RequiredRAM
// 							*bwAvailable -= process.RequiredBW
// 							if process.Name == "VPN" {
// 								go runVpnProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 							} else {
// 								go runProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 							}
// 							break
// 						}
// 						time.Sleep(100 * time.Millisecond)
// 					}
// 				}
// 			default:
// 				// Continue running the FCFS logic without blocking
// 				for _, process := range *processes {
// 					if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
// 						*cpuAvailable -= process.RequiredCPU
// 						*ramAvailable -= process.RequiredRAM
// 						*bwAvailable -= process.RequiredBW
// 						if process.Name == "VPN"{
// 							go runVpnProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh )
// 						}else{
// 							go runProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 						}

// 					} else {
// 						for {
// 							if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
// 								*cpuAvailable -= process.RequiredCPU
// 								*ramAvailable -= process.RequiredRAM
// 								*bwAvailable -= process.RequiredBW
// 								if process.Name == "VPN"{
// 									go runVpnProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh )
// 								}else{
// 									go runProcess(processes , process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
// 								}
// 								break
// 							}
// 							time.Sleep(100 * time.Millisecond)
// 						}
// 					}
// 				}
// 			}
// 			time.Sleep(100 * time.Millisecond) // Sleep to avoid excessive CPU usage
// 		}

// }

func FCFS(processes *[]*models.Process, cpuAvailable *int, ramAvailable *int, bwAvailable *int, newProcessCh chan *models.Process) {
	resourceCh := make(chan struct{}, 1) // Channel for resource availability signal, buffer size 1 to avoid blocking

	for _, _ = range *processes {
		select {
		case process := <-newProcessCh:
			if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
				*cpuAvailable -= process.RequiredCPU
				*ramAvailable -= process.RequiredRAM
				*bwAvailable -= process.RequiredBW
				if process.Name == "VPN" {
					go runVpnProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
				} else {
					go runProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
				}
			} else {
				for {
					if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
						*cpuAvailable -= process.RequiredCPU
						*ramAvailable -= process.RequiredRAM
						*bwAvailable -= process.RequiredBW
						if process.Name == "VPN" {
							go runVpnProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
						} else {
							go runProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
						}
						break
					}
					time.Sleep(100 * time.Millisecond)
				}
			}
		default:
			// Continue running the FCFS logic without blocking
			for _, process := range *processes {
				if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
					*cpuAvailable -= process.RequiredCPU
					*ramAvailable -= process.RequiredRAM
					*bwAvailable -= process.RequiredBW
					if process.Name == "VPN" {
						go runVpnProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
					} else {
						go runProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
					}
				} else {
					for {
						if process.RequiredCPU <= *cpuAvailable && process.RequiredRAM <= *ramAvailable && process.RequiredBW <= *bwAvailable {
							*cpuAvailable -= process.RequiredCPU
							*ramAvailable -= process.RequiredRAM
							*bwAvailable -= process.RequiredBW
							if process.Name == "VPN" {
								go runVpnProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
							} else {
								go runProcess(processes, process, cpuAvailable, ramAvailable, bwAvailable, resourceCh)
							}
							break
						}
						time.Sleep(100 * time.Millisecond)
					}
				}
			}
		}
		time.Sleep(100 * time.Millisecond) // Sleep to avoid excessive CPU usage
	}
}