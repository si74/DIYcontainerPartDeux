package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type Mount struct {
	Source string
	Target string
	Fs     string
	Flags  int
	Data   string
}

type Cfg struct {
	Path     string
	Args     []string
	Hostname string
	Mounts   []Mount
	Rootfs   string
}

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what?")
	}
}

func run() {

	// link to currently running process
	// http://unix.stackexchange.com/questions/333225/which-process-is-proc-self-for
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...) // link to currently running process
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("running %v as pid%v\n", os.Args[2:], os.Getpid())

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	must(cmd.Run())
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
