package main

import (
	"flag"
	"log"
	"net/rpc"
	"os/exec"
)

type JobResult struct {
	Result   string
	ExitCode int
}

func main() {
	flag.Parse()
	runcmd := flag.Arg(0)
	argcmd := flag.Arg(1)
	log.Printf("Running %s", runcmd)
	cmd := exec.Command(runcmd, argcmd)
	// err is never present for some reason, even if we execute a command that
	// doesn't exist.
	out, err := cmd.Output()

	client, err := rpc.Dial("unix", "/tmp/kallice.sock")
	if err != nil {
		log.Fatal("Connection failed dialing: ", err)
	}
	var reply string

	args := &JobResult{}
	if err == nil {
		log.Printf("Command finished successfully: %s", out)
		// almost certainly the wrong way to convert []byte into a string
		args.Result = string(out)
		args.ExitCode = 0
	} else {
		log.Printf("Command finished with error: %v", err)
		args.Result = string(out)
		args.ExitCode = 255
	}

	err = client.Call("JobReporter.ReportResult", args, &reply)
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}
	log.Printf("Returned from server: %s", reply)
}
