package cmd

import (
	"fmt"
	"os"
	"portScanner/port"
	"time"

	"github.com/spf13/cobra"
)

func Execute() {
	var hostname string
	var start int
	var end int
	var protocol string
	channel := make(chan []port.ScanResult)
	goScanCmd := &cobra.Command{
		Use:   "GOPORT",
		Short: "simple port scanner",
		Long:  `simple port scanner`,
		Run: func(cmd *cobra.Command, args []string) {
			go port.InitialScan(hostname, protocol, start, end, &channel)
			fmt.Println("Started: ", time.Now().Local().UTC())
			fmt.Println(<-channel)
			fmt.Println("Finished: ", time.Now().Local().UTC())
		},
	}

	goScanCmd.Flags().StringVarP(&hostname, "hostname", "a", "", "Hostname to scan")
	goScanCmd.Flags().IntVarP(&start, "start", "s", 1, "Start port")
	goScanCmd.Flags().IntVarP(&end, "end", "e", 65535, "End port")
	goScanCmd.Flags().StringVarP(&protocol, "protocol", "p", "tcp", "Protocol to scan")
	goScanCmd.MarkFlagRequired("hostname")
	goScanCmd.Example = `goScan -s 1 -e 65535 -a website.com or ip address`
	if err := goScanCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
