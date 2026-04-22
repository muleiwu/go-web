//go:build !windows

package cmd

import (
	"os"
	"syscall"
)

// execSelf 使用 syscall.Exec 用当前可执行文件替换进程映像。
// PID 保持不变，成功时不会返回（控制权交给新进程映像）。
func execSelf() error {
	bin, err := os.Executable()
	if err != nil {
		bin = os.Args[0]
	}
	return syscall.Exec(bin, os.Args, os.Environ())
}
