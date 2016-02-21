## Kallice CI 

A UNIX continuous integration server.

### Overview

The simplest job is based on a periodic schedule (you can even use the crontab
format if you'd like). More flexible jobs can use a job producer. A job producer
is a process that tells the master that a job should be run. The scheduled job
producer is a job producer that uses cron.

When Kallice starts (kallice-server /path/to/config) it reads the config and
starts producers, and then waits. When a producer sends a job to the master, the
master spawns a job runner of the type described in the config file.  The runner 
executes the job and monitors execution. When the job is finished, the runner
reports the results back to the master. 

There are different types of runners.

simple: The simplest runner does a fork/exec and checks the return code. If it
is non-zero, the job has failed.
pipe: read from a pipe
file: monitor a file
tap: read tap output
http: look for a specific response to a request
net: monitor fom a TCP or UDP socket
unix-socket: monitor a UNIX socket

### Install

Copy the binaries to a suitable location in your PATH.

What perms does Kallice need? Where should it write its data?

Might want to make distribution / install a tarball that contains a Makefile
we can use to install: make install, etc.

### Setup

### Use

### Why?

### How?
