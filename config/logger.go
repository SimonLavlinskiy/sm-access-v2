package config

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func loggerInit() {
	redirectStdout()
	redirectStderr()
}

func redirectStdout() {
	f, err := os.OpenFile("/logs/outLog.log", syscall.O_APPEND|syscall.O_RDWR|syscall.O_CREAT, 0644)
	if err != nil {
		fmt.Println("can't open file for redirecting logs", err)
	}
	err = syscall.Dup2(int(f.Fd()), int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("can't redirect logs to file", err)
	} else {
		fmt.Println(time.Now(), "logs redirected to file")
	}
}

func redirectStderr() {
	f, err := os.OpenFile("/logs/errLog.log", syscall.O_APPEND|syscall.O_RDWR|syscall.O_CREAT, 0644)
	if err != nil {
		fmt.Println("can't open file for redirecting logs", err)
	}
	err = syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		fmt.Println("can't redirect logs to file", err)
	} else {
		fmt.Println(time.Now(), "logs redirected to file")
	}
}
