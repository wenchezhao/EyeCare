package fyne

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/wenchezhao/EyeCare/config"
)

func ShowFullScreenPopup(a fyne.App, cfg *config.Config) {
	var pop fyne.Window
	fyne.CurrentApp().Driver().DoFromGoroutine(func() {
		pop = a.NewWindow("休息提醒")
		pop.SetFullScreen(true)
		pop.RequestFocus()
	}, false)

	done := make(chan bool)
	// 创建按钮和计时器标签
	countdownLabel := widget.NewLabel("")
	noButton := widget.NewButton("提前退出", func() {
		pop.Close()
	})

	// 创建按钮容器并居中
	buttons := container.NewCenter(container.NewHBox(noButton))
	// 创建标签容器并居中
	labelContainer := container.NewCenter(countdownLabel)

	// 创建垂直布局并居中所有元素
	content := container.NewCenter(
		container.NewVBox(
			labelContainer,
			buttons,
		),
	)

	go func() {
		for i := cfg.RestDuration; i >= 0; i-- {
			fyne.CurrentApp().Driver().DoFromGoroutine(func() {
				countdownLabel.SetText(fmt.Sprintf("请远眺20英尺（约6米）休息，还剩 %d 秒", i))
				pop.RequestFocus()
			}, false)
			time.Sleep(1 * time.Second)
			// 检查是否已经被提前关闭
			select {
			case <-done:
				return
			default:
			}
		}
		close(done)
	}()
	noButton.OnTapped = func() {
		defer close(done)
	}
	fyne.CurrentApp().Driver().DoFromGoroutine(func() {
		pop.SetContent(content)
		pop.Show()
		pop.RequestFocus()
	}, false)
	// 等待结束
	<-done
	fyne.CurrentApp().Driver().DoFromGoroutine(func() {
		pop.Close()
	}, false)

}
