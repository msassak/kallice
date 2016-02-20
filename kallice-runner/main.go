package main

import (
	"flag"
	"log"
	"os/exec"
)

func main() {
	flag.Parse()
	runcmd := flag.Arg(0)
	argcmd := flag.Arg(1)
	log.Printf("Running %s", runcmd)
	cmd := exec.Command(runcmd, argcmd)
	out, err := cmd.Output()
	if err == nil {
		log.Printf("Command finished successfully: %s", out)
	} else {
		log.Printf("Command finished with error: %v", err)
	}
}
