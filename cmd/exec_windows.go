//go:build windows

package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

// execSelf 在 Windows 下的 Restart 实现：spawn 一份新的自身进程，
// 继承 stdio 和参数/环境，然后父进程 os.Exit(0) 退出。
// PID 会发生变化（Windows 无 exec 原语），其余语义尽量贴近 POSIX。
// 成功时不返回（通过 os.Exit 终止当前进程）；失败时返回 error 给调用方。
func execSelf() error {
	bin, err := os.Executable()
	if err != nil {
		bin = os.Args[0]
	}

	child := exec.Command(bin, os.Args[1:]...)
	child.Stdin = os.Stdin
	child.Stdout = os.Stdout
	child.Stderr = os.Stderr
	child.Env = os.Environ()
	if wd, err := os.Getwd(); err == nil {
		child.Dir = wd
	}

	if err := child.Start(); err != nil {
		return fmt.Errorf("spawn child process failed: %w", err)
	}

	// 让子进程脱离父进程生命周期
	_ = child.Process.Release()

	// 父进程干净退出，控制权交给新子进程
	os.Exit(0)
	return nil
}
