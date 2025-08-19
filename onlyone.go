package main

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows"
)

func checkMutex() {
	// 创建全局唯一互斥量名字
	mutexName := "Global\\MyUniqueGoAppMutexEyeCare"

	// 创建互斥量
	_, err := windows.CreateMutex(nil, false, windows.StringToUTF16Ptr(mutexName))
	if err != nil && err == windows.ERROR_ALREADY_EXISTS {
		fmt.Println("已有程序实例在运行，退出", err)
		os.Exit(0)
	}
}
