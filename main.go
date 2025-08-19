package main

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var icon = `AAABAAEAICAAAAEAIACoEAAAFgAAACgAAAAgAAAAQAAAAAEAIAAAAAAAABAAAMMOAADDDgAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/3Z7AP92ewD/dnsC/3V4AP91eRb/dnxy/3h/cv95gxb/dnsA/36NBf+AlAL/gJMA/4CSAP+BlQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP93fQD/d30A/3V6Af91eQD/dXoh/3Z7pf92fFT/eIBV/3qFpP98iR3/fIkA/36OAv+DmQH/g5kA/4OZAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/3d/AP91eQD/XzsA/3+cAP92e6P/dnsx/3V4AP99iwD/e4cu/32Lvf9/kXT/gJMo/4CTAP+AkgH/fo4AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/eIEA/3Z8A/92fAD/dnxe/3Z8d/92fAD/eIEJ/3+RAP9+jVf/fo6E/3+RQ/+BlaT/g5kT/4OZAP+ClwEAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/b2sA/1EeAP9rYAD/b2wB/3FvAv9ycgL/dHYB/3V6Af93fgD/e4YA/3mBAP93fpr/d38a/3iAAP9+jQT/f5AA/3+Pi/9/kS7/gpcA/4Oafv+Dmj//g5oA/4WeA/+MrwH/jrIC/4+2Av+RuQH/l8gA/6LgAP+SuwD/kbsA/25qAP9uagD/bmoB/29rA/9taAD/bWcA/3JyAP94gAD/eIAA/3iABP93fwD/eIA1/3iAh/94gAD/eIAD/4CTA/+AkwD/gJNM/4GVm/+Em2T/hJyc/4eiDP9xcQD/gpgA/4moAP+TvAD/lMEA/5TBAP+UwAP/lMEB/5TBAP+UwQD/cm8A/25pAP9wbQH/cHAA/29sGv9xbzj/cnI6/3R2K/91eRD/eoQD/3mDAP95gm3/eYNS/3h/AP97iAT/cX8A/4OaAf+AkwD/gpcw/4Scmf+Fn2z/hZ4A/4OaA/+LrA//jK8q/46yOf+Ptjj/kbkb/4+3AP+VwwH/j7YA/5fHAP9uaQH/bWcA/21nCP9va4b/cG2T/3Fwef9yc3r/dHaG/3Z6lf93fpb/eYJo/3qFpf97hiH/e4YA/3uHAf+DmwP/hZ4D/4ejAf+HogD/h6Ih/4ejpv+JqGn/iqqV/4utk/+NsIX/jrN5/4+2ef+RupX/k7+I/5bFCP+WxQD/lMEB/29rBP9vawD/b2t3/29raP9vawD/b2sC/3N2AP+IqAD/eYMB/3d+F/95g1j/e4jO/32Mmv9/kHH/gZUe/3+QAP+IpAD/hp8e/4ahc/+Ho5z/iabO/4moWP+Lqxf/jbEB/2leAP+PtgD/lcQC/5XEAP+Uwmn/lcR3/5XDAP+VwwT/cG4C/3BuAP9wbYz/cG4b/3FvAP9xbwX/cnME/3R2A/96hAL/fIoA/32LKP99i4T/g5gG/3+QVv+BlY7/g5qI/4WeiP+Fno7/hqFV/4OaBv+Kq4T/i6so/4usAP+LrAL/jK8D/46zBP+RugX/mcwA/5fHHf+XyI3/l8cA/5fHAv9ycgT/cnIA/3JxfP9ycj//cXEA/3JxAv93fQD/e4cA/32MA/9+jgD/fo1K/36Odf9+jwD/g5oP/4OahP+Em63/hp+t/4ikg/+NsQ//jK8A/4yudf+Mr0n/jK8A/4yvA/+MrwD/mcsA/5jKAv+YygD/mMpB/5jLff+YygD/mMsE/3N1Av9zdQD/c3U3/3R3k/9ycwD/dnsE/36OAP+QuQD/f5AE/3+PAP9/kFz/f5Fg/4KYWv+CmKr/hJtI/32OAP+LqwD/iKVI/4qpq/+MrVr/jrJg/46yXP+NsAD/jrQE/4SfAP+c0wD/ms0E/5nMAP+ZzJX/mc03/5nNAP+ZzQL/dXoA/3d9A/93fQD/dnuL/3d9Vv92ewD/eYEE/4GUAf+Akwb/gJMA/4CTdP+Blc//gphf/4ScBf+HogD/h6IJ/4ikCf+HowD/iqgF/4yvX/+PtND/kLdz/5C2AP+SvQb/lsYB/5rOBP+azgD/ms5V/5rOjP+azgD/mc0D/5vPAP9ycwD/eYIB/3d/AP91eQz/eYGt/3qFLf94gQD/fo0E/2lmAP+Ak2f/gZbe/4KYQf+ClwD/fIsC/4ikbv+Jppn/iqqZ/4yubf+WxQL/kbkA/5C4Qf+Rut7/k79n////AP+ZzAT/ms8A/5rOLP+az6v/mtAM/5vRAP+b0AH/mcwA/3V5AP9xcQD/eoQC/3uGAP96hCX/fIi3/32NH/+BlQn/gJOl/4CTYv+Dmnz/hJs3/4ahAP+IpG//iKWC/3NzAP+9/wD/jbCD/46zbv+QuAD/krw4/5K9e/+UwWL/lsWl/5jLCf+azx7/ms+y/5vQI/+b0AD/ms8C/5rOAP+azgD/fYwA/32MAP99jAD/fYsC/3yIAP97iB//fo+v/4CTs/+Akzz/g5oA/4SdhP+FnS7/g5kA/4mml/+Jpgr/h6MA/4+2AP+Oswr/kLeX/5TAAP+UwC//lMGD/5XCAP+WxTz/mcyy/5rPrP+aziL/ms4A/5rOAv+azwD/ms8AAAAAAP96hgD/fYsC/3yKAP98ijD/fYyU/3+Qfv+AkqL/gZW0/4ScPP+FnwD/hqCE/4ahLv+GoAD/iqqX/4yuCv+LqgD/k74A/5G7Cv+Ru5f/l8cA/5XEL/+WxIT/l8cA/5rOOv+azrD/nNKv/57ZIv+d1QD/n9sC/6PjAP+j4wAAAAAA/32LAf99iwD/fYsN/32LpP99iyf/gJJe/4GViv+IpgL/hJym/4agYv+IpHz/iKU3/4qpAP+Mrm//jbGC/29pAP+o7gD/kryD/5K9bv+UwAD/lsY4/5fHe/+azmD/ms6j/5jMCf+d1SD/ntm1/6DeI/+f2gD/oN0C/6XpAP+m6QD/fo0B/36NAP9+jSb/fo6M/3+RAP+CmBj/g5mX/4OZAP95hQH/h6Jo/4mn3v+KqUH/iqgA/4WeAv+Os27/kLeZ/5G7mf+SvW7/n9gC/5jKAP+YyUH/mcvd/5rOZv+4/wD/nNQE/6LgAP+g3C7/oeCs/6XoDP+i4gD/od8B/6juAP9/kQD/gpcB/4GVAP+Ak6z/gZaO/4Oanv+Em03/hJsA/4ahCP+KqwD/iqp0/4ut0P+Osl//kLkF/5O+AP+TvQn/lL8J/5O/AP+WxQX/mMlf/5nMz/+aznP/ms4A/5rOBv+ZzQH/oeAE/6TmAP+j5Fb/pOaL/6PkAP+j5QP/pecA/4GUAv+BlAD/gZQ1/4GVjP+CmCv/g5ko/4KYAP+ClwH/jK0E/42xAP+Mr1z/jK9g/460W/+Qt6r/krxI/4+1AP+d0gD/l8ZI/5jJq/+YyVr/m9Bg/5vRXP+b0gD/ntcG/6TlAv///wD/pOYE/6juAP+m6pT/p+w2/6fsAP+n7AL/gpYE/4KXAP+Clnz/gpc//4KXAP+ClwP/g5oC/4KYAP+OsgP/jrIA/46ySf+Os3X/jrIA/42wEP+SvYT/lMKt/5bGrf+Xx4T/l8cP/5zSAP+c03X/nNRJ/5vSAP+c1AP/p+8A/6fsAv+o7wL/qfAA/6jvQP+o8Hv/qO8A/6jvBP+DmgL/g5oA/4OZjP+Dmhv/gZQA/4ikBv+MrwT/jbID/4+1Av+PtgD/j7Yo/5C2hP+YyQX/lMBT/5XDjP+Vw4f/l8eI/5nMjv+b0VX/l8kG/53WhP+d1ij/ntkA/6LiPv+j5UL/qfAC/6rzBP+q8wD/qvMc/6r0jP+q8wD/qvMC/4WeBP+FngD/hZ13/4afaf+FnQD/hZ0C/4qqAP+b0wD/ja4B/4+2Fv+RuVb/kbvN/5O+l/+UwG//lcIe/5K9AP+b0AD/mcwe/5vRcf+d1Zr/n9nR/6DdS/+i4XD/ouKO/6TniP+m6nn/pOUA/635AP+r9mn/q/Z3/6v2AP+r9gT/hqAB/4ScAP+EnAj/h6KH/4mnlv+Lq3r/jK56/42xhf+PtJP/kLeU/5G5Z/+TvqX/lMAh/5TAAP+TvgH/lcQD/5fGA/+f2gH/n9sA/5/bIv+g3KL/od9z/6Pjt//e/wD/qfEA/6ftov+p8Yj/q/WO/6v2hf+t+gj/rfoA/6z4Af+DmQD/iqsA/4WeAf+LqgD/iaga/4urOP+Mrzr/jrIr/4+1EP+WxgP/lcMA/5XDbf+Vw1L/lcQA/5XDBP+NsgD/mcsA/5/ZBP+i4gD/od5T/6Hfbv+j4gD/peeR/6bqWf+o8Ej/qfCf/6ryM/+r9Rr/qvEA/6r0Af+s+QD/qfMA/4agAP+GoAD/hqAB/4ahA/+GnwD/hp8A/4elAP+RuAD/kLgA/5O/A/+ZzQD/lsY1/5fHh/+WxQD/l8cD/5nLAP+i4AD/ouED/6LhAP+i4Yf/ouE2/6PkAP+l6BT/p+x6/6jwgf+p8hr/qfIA/6jwAf+r9gP/rPcB/6z3AP+s9wD/hqIA/4imAP9jUQD/gZQA/4moAf+LqwL/jK8C/46yAf+PtQH/lsYA/6PkAP+g3AD/mcya/5rOGv+YygD/mcsB/6LhAf+i4QD/o+Ib/6Pjmv+i4QD/oeAB/6PkAf+f2wD/qfEA/6jvAf+p8gL/q/UB/67/AP/i/wD/q/YAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAP+WxAD/m9AD/5vRAP+b0F7/m9F2/5vQAP+e2Qf/pOUH/6TlAP+k5Xf/pOVd/6TlAP+k5QP/pekB/6fsBP+o8AT/qfEB/6nxAP+p8QD/qfIAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/5jJAP+d1wD/ZlgA////AP+d1qP/n9kx/53WAP+l6QD/pOYx/6Tmo/+f1wD/AAAA/6XoAP+j4gD/p+8A/6fvAP+n7wD/p+8AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/mcwA/5rOAP+f2gH/n9oA/57YIv+g3KX/ouFU/6TlVP+k5qX/pech/6XoAP+l5wH/o+QA/6PkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAD/nNMA/5zTAP+f2wL/o+QA/6HeF/+i4nL/pOVy/6XoFv+l6QD/pOYC/6TmAP+k5gAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA/9Ql///IEv//mYX//5KC/6ASkgVIpIISECRECEAEIAJIAYASSEACEkhIEhJLQYLSJEJCJCIIEERQCZAKSEmSFSBJkhVACZAKSQgQRCFCQiRCQYISSUgSkkhAAhJIAYAyQAQgghAlpAhIpCQioBJJhf+SSC//mZH//8gT///UK/8=`

func showFullScreenPopup(a fyne.App) {
	var pop fyne.Window
	fyne.CurrentApp().Driver().DoFromGoroutine(func() {
		pop = a.NewWindow("休息提醒")
		pop.SetFullScreen(true)
	}, false)

	done := make(chan bool)
	// 创建按钮和计时器标签
	countdownLabel := widget.NewLabel("")
	//yesButton := widget.NewButton("是", nil)
	noButton := widget.NewButton("提前退出", func() {
		pop.Close()
	})

	// 创建按钮容器
	buttons := container.NewHBox( /*yesButton,*/ noButton)
	content := container.NewVBox(
		countdownLabel,
		buttons,
	)

	go func() {
		for i := 20; i >= 0; i-- {
			fyne.CurrentApp().Driver().DoFromGoroutine(func() {
				countdownLabel.SetText(fmt.Sprintf("请远眺20英尺（约6m）休息，还剩 %d 秒", i))
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
		pop.SetContent(container.NewCenter(content))
		pop.Show()
	}, false)
	// 等待结束
	<-done
	fyne.CurrentApp().Driver().DoFromGoroutine(func() {
		pop.Close()
	}, false)

}

func main() {
	checkMutex()
	a := app.New()
	// 检查是否支持系统托盘
	if desk, ok := a.(desktop.App); ok {
		byteIcon, err := base64.StdEncoding.DecodeString(icon)
		if err != nil {
			panic(err)
		}
		iconR := fyne.NewStaticResource("icon", byteIcon)
		fyne.CurrentApp().SetIcon(iconR)
		// 创建托盘菜单
		autostartItem := fyne.NewMenuItem("开机启动", nil)
		autostartItem.Checked = isAutoStartEnabled()
		autostartItem.Action = func() {
			var err error
			if autostartItem.Checked {
				err = toggleAutoStart(false)
			} else {
				err = toggleAutoStart(true)
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
			fyne.NewMenuItem("退出", func() {
				a.Quit()
			}),
		)
		desk.SetSystemTrayIcon(iconR) // 需要添加一个图标资源
		desk.SetSystemTrayMenu(menu)
	}

	// 循环弹全屏窗口
	go func() {
		for {
			time.Sleep(20 * time.Minute) // 等待 20 分钟
			showFullScreenPopup(a)
		}
	}()

	a.Run()
}

func getIcon() {
	bytes, err := os.ReadFile(filepath.Clean("icon.ico"))
	if err != nil {
		return
	}
	fmt.Println("读取图标文件成功：", base64.StdEncoding.EncodeToString(bytes))
}
