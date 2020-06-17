package daemon

import (
	"os"
	"os/exec"
	"syscall"
)

func Daemon() {
	if syscall.Getppid() == 1 {
		return
	}

	args := os.Args[1:]

	cmd := exec.Command(os.Args[0], args...)
	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	err := cmd.Start()
	if err == nil {
		cmd.Process.Release()
		os.Exit(0)
	}
}
