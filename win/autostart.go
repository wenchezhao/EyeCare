package win

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"syscall"
)

// Add this function to handle autostart
func ToggleAutoStart(enable bool) error {
	if runtime.GOOS != "windows" {
		return fmt.Errorf("autostart only supported on Windows")
	}

	startupDir := filepath.Join(os.Getenv("APPDATA"), "Microsoft\\Windows\\Start Menu\\Programs\\Startup")
	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	shortcutPath := filepath.Join(startupDir, "EyeCare.lnk")

	if enable {
		// Create shortcut using PowerShell
		psCmd := `$WS = New-Object -ComObject WScript.Shell; ` +
			`$SC = $WS.CreateShortcut('` + shortcutPath + `'); ` +
			`$SC.TargetPath = '` + exePath + `'; ` +
			`$SC.Save()`

		cmd := exec.Command("powershell", "-NoProfile", "-NonInteractive", "-WindowStyle", "Hidden", "-Command", psCmd)

		// 隐藏窗口
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		return cmd.Run()
	} else {
		// Remove shortcut
		return os.Remove(shortcutPath)
	}
}

// Add this function to check if autostart is enabled
func IsAutoStartEnabled() bool {
	if runtime.GOOS != "windows" {
		return false
	}
	startupPath := filepath.Join(os.Getenv("APPDATA"), "Microsoft\\Windows\\Start Menu\\Programs\\Startup", "EyeCare.lnk")
	_, err := os.Stat(startupPath)
	return err == nil
}
