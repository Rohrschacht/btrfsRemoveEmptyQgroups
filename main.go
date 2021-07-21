package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/rohrschacht/checkroot"
)

func main() {
	checkroot.RootOrExit()

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Please specify path to btrfs filesystem!\n")
		os.Exit(1)
	}

	cmd := exec.Command("btrfs", "qgroup", "show", os.Args[1])
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(stdout)
	// skip first two lines of output
	scanner.Scan()
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Fields(line)

		if cols[1] == "0.00B" && cols[2] == "0.00B" {
			qgroupid := strings.TrimPrefix(cols[0], "0/")

			err := exec.Command("btrfs", "qgroup", "destroy", qgroupid, os.Args[1]).Run()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
