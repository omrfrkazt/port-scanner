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
	conn, err := net.DialTimeout(protocol, hostname+":"+strconv.Itoa(port), time.Second)
	if err != nil {
		model := ScanResult{strconv.Itoa(port), "CLOSED"}
		*channel <- model
		return
	}
	defer conn.Close()
	model := ScanResult{strconv.Itoa(port), "OPEN"}
	*channel <- model
}

func InitialScan(hostname string, start int, end int, mainChannel *chan []ScanResult) {
	if start < 1 {
		start = 1
		end = 2
	}
	channel := make([]chan ScanResult, end)
	for i := start; i <= end; i++ {
		channel[i-1] = make(chan ScanResult)
	}
	for i := start; i <= end; i++ {
		go ScanPort("tcp", hostname, i, &channel[i-1])
	}
	collection := make([]ScanResult, end)
	for i, v := range channel {
		collection[i] = <-v
	}
	*mainChannel <- collection
}
