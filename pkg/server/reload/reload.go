package reload

import "sync"

var (
	reloadChan  chan struct{}
	restartChan chan struct{}
	reloadOnce  sync.Once
	restartOnce sync.Once
)

// GetReloadChan 获取重载通道（单例模式）
func GetReloadChan() <-chan struct{} {
	reloadOnce.Do(func() {
		reloadChan = make(chan struct{}, 1)
	})
	return reloadChan
}

// TriggerReload 触发重载信号
func TriggerReload() {
	reloadOnce.Do(func() {
		reloadChan = make(chan struct{}, 1)
	})
	select {
	case reloadChan <- struct{}{}:
		// 成功发送重载信号
	default:
		// 通道已满，忽略本次请求
	}
}

// Reload 触发软重载：重置容器、重新装配、重启服务，不退出当前进程。
// 与 SIGHUP 行为一致。非阻塞，若已有重载待处理则忽略本次调用。
func Reload() {
	TriggerReload()
}

// GetRestartChan 获取进程重启通道（单例模式）
func GetRestartChan() <-chan struct{} {
	restartOnce.Do(func() {
		restartChan = make(chan struct{}, 1)
	})
	return restartChan
}

// Restart 触发进程自替换：先优雅停止所有服务，再用 syscall.Exec 用当前
// 可执行文件替换进程映像。PID 保持不变，gomander 管理的 PID 文件仍有效。
// 非阻塞，调用方可在返回后先完成当前 HTTP 响应再由外层 exec 替换进程。
// Windows 下返回无错，但实际 exec 会在 cmd 层失败并记录日志。
func Restart() {
	restartOnce.Do(func() {
		restartChan = make(chan struct{}, 1)
	})
	select {
	case restartChan <- struct{}{}:
	default:
	}
}
