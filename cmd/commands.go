package cmd

import (
	"fmt"
	"os"
	"portScanner/models"
	"portScanner/port"
	"time"

	"github.com/spf13/cobra"
)

func Execute() {
	var (
		hostname string
		protocol string
		start    int
		end      int
	)
	channel := make(chan models.AppResult)
	//cobra tool is used to create a command line interface
	//for details see https://cobra.dev/
	goScanCmd := &cobra.Command{
		Use:   "GOPORT",
		Short: "simple port scanner",
		Long:  `simple port scanner`,
		Run: func(cmd *cobra.Command, args []string) {
			if end > 65535 || start < 1 {
				fmt.Println("Port range is must be between 1 and 65535")
				os.Exit(1)
			}
			go port.InitialScan(hostname, protocol, start, end, &channel)
			fmt.Println("Started: ", time.Now().Local().UTC())
			var result models.AppResult = <-channel
			fmt.Println("Finished: ", time.Now().Local().UTC())
			fmt.Println("Results saved to: ", result.FileName)
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
