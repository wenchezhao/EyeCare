package fyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"

	"github.com/wenchezhao/EyeCare/config"
	"github.com/wenchezhao/EyeCare/win"
)

func ShowTrayMenu(a fyne.App, cfg *config.Config, iconR *fyne.StaticResource) {
	if desk, ok := a.(desktop.App); ok {
		// 创建托盘菜单
		autostartItem := fyne.NewMenuItem("开机启动", nil)
		autostartItem.Checked = win.IsAutoStartEnabled()
		// 添加主题切换菜单项
		themeItem := fyne.NewMenuItem("暗色主题", nil)
		themeItem.Checked = cfg.DarkTheme
		themeItem.Action = func() {
			cfg.DarkTheme = !cfg.DarkTheme
			// 保存配置
			if err := config.SaveConfig(cfg); err != nil {
				dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
				return
			}
			// 更新主题
			if cfg.DarkTheme {
				a.Settings().SetTheme(theme.DarkTheme())
			} else {
				a.Settings().SetTheme(theme.LightTheme())
			}
			// 更新菜单项状态
			themeItem.Checked = cfg.DarkTheme
			// 重新创建菜单并更新
			menu := fyne.NewMenu("护眼提醒",
				autostartItem,
				themeItem,
				fyne.NewMenuItem("退出", func() {
					a.Quit()
				}),
			)
			desk.SetSystemTrayMenu(menu)
		}
		autostartItem.Action = func() {
			var err error
			if autostartItem.Checked {
				err = win.ToggleAutoStart(false)
			} else {
				err = win.ToggleAutoStart(true)
			}

			if err != nil {
				// 显示错误对话框而不是打印到控制台
				dialog.ShowError(err, fyne.CurrentApp().Driver().AllWindows()[0])
				return
			}

			// 只有在操作成功时才更新状态
			autostartItem.Checked = !autostartItem.Checked
			// 重新创建菜单并更新
			menu := fyne.NewMenu("护眼提醒",
				autostartItem,
				themeItem,
				fyne.NewMenuItem("退出", func() {
					a.Quit()
				}),
			)
			// 刷新菜单显示
			desk.SetSystemTrayMenu(menu)
		}
		// 创建托盘菜单
		menu := fyne.NewMenu("护眼提醒",
			autostartItem,
			themeItem,
			fyne.NewMenuItem("退出", func() {
				a.Quit()
			}),
		)
		desk.SetSystemTrayIcon(iconR)
		desk.SetSystemTrayMenu(menu)
	}
}
