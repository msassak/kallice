package main

import (
	"flag"
	"log"
	"net"
	"net/rpc"

	"github.com/BurntSushi/toml"
)

type MasterConfig struct {
	LogRoot    string `toml:"log_root"`
	LogLevel   string `toml:"log_level"`
	PidRoot    string `toml:"pid_root"`
	Include    string
	ListenPort int `toml:"listen_port"`
	Producers  map[string]ProducerConfig
	Runners    map[string]RunnerConfig
	Jobs       map[string]JobConfig
}

type ProducerConfig struct {
	Type     string
	Interval int
}

type RunnerConfig struct {
	Type       string
	ReturnWait int `toml:"return_wait"`
}

type JobConfig struct {
	Producer string
	Runner   string
	Command  string
	Args     []string
}

var conf MasterConfig

type JobResult struct {
	Result   string
	ExitCode int
}

type JobReporter int

func (t *JobReporter) ReportResult(res *JobResult, reply *string) error {
	log.Printf("Runner result %s with exit code %d", res.Result, res.ExitCode)
	*reply = "ack"
	return nil
}

func serve() {
	reporter := new(JobReporter)
	rpc.Register(reporter)
	l, e := net.Listen("unix", "/tmp/kallice.sock")
	if e != nil {
		log.Fatal("UNIX socket failure: ", e)
	}
	rpc.Accept(l)
}

func main() {
	flag.Parse()
	var configPath = flag.Arg(0)

	if len(configPath) == 0 {
		log.Fatal("Need path to config file.")
	}

	log.Printf("configPath is %s", configPath)

	var conf MasterConfig
	_, err := toml.DecodeFile(configPath, &conf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("kallice-server read config data! %v", conf)

	// makes no attempt to clean up the socket, so you have to delete by hand
	// when restarting
	serve()
}
