package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  string
	State string
}

func ScanPort(protocol, hostname string, port int, channel *chan ScanResult) {
	conn, err := net.DialTimeout(protocol, hostname+":"+strconv.Itoa(port), time.Second*1)
	if err != nil {
		model := ScanResult{strconv.Itoa(port), "CLOSED"}
		*channel <- model
		return
	}
	defer conn.Close()
	model := ScanResult{strconv.Itoa(port), "OPEN"}
	*channel <- model
}

func InitialScan(hostname, protocol string, start, end int, mainChannel *chan []ScanResult) {
	channel := make([]chan ScanResult, end-start+1)
	for i := start; i <= end; i++ {
		channel[i-start] = make(chan ScanResult)
		go ScanPort(protocol, hostname, i, &channel[i-start])
	}
	var result []ScanResult
	for i := start; i <= end; i++ {
		result = append(result, <-channel[i-start])
	}
	*mainChannel <- result
}
