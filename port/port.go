package port

import (
	"net"
	"portScanner/models"
	"portScanner/utils"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

// ScanPort scans port
func ScanPort(protocol, hostname string, port int, fileName string) {
	//net.Dialtimeout is a function that returns a net.Conn that is connected to the specified network address.
	conn, err := net.DialTimeout(protocol, hostname+":"+strconv.Itoa(port), time.Second*1)
	//if has error then port is closed
	if err != nil {
		model := models.ScanResult{Port: port, State: "CLOSED"}
		utils.WriteTxt(model, fileName)
		wg.Done()
		return
	}
	//we need close every connection
	defer conn.Close()
	//if has no error then port is open
	model := models.ScanResult{Port: port, State: "OPEN"}
	utils.WriteTxt(model, fileName)
	//response is a channel that will be closed when the response is received
	wg.Done()
}

// InitialScan scans ports from start to end
func InitialScan(hostname, protocol string, start, end int, mainChannel *chan models.AppResult) {
	//add wait group every port
	wg.Add(end - start + 1)
	//create file
	fileName := utils.CreateTxt(hostname)
	//fill channel with channels
	for i := start; i <= end; i++ {
		//goroutine is a lightweight thread of execution.
		go ScanPort(protocol, hostname, i, fileName)
	}
	//wait all goroutines
	wg.Wait()
	//return channel
	*mainChannel <- models.AppResult{FileName: fileName}
}
