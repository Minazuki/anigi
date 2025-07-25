package main

import (
	"anigi/anigi"
	"anigi/config"
	"context"
	"flag"
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/viper"
)

var (
	workDir      = flag.String("d", ".", "工作目錄")
	showVersion  = flag.Bool("v", false, "顯示版本號")
)

type Config struct {
    Anigi config.AnigiCfg `json:"anigi"`
}

func main() {
	getWin := syscall.NewLazyDLL("user32.dll").NewProc("GetConsoleWindow")
	showWin := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	if getWin.Find() == nil && showWin.Find() == nil {
		hwnd, _, _ := getWin.Call()
		if hwnd != 0 {
			showWin.Call(hwnd, 0) // SW_HIDE
		}
	}
	for _, arg := range os.Args[1:] {
		if arg == "-help" || arg == "--help" || arg == "help" {
			fmt.Println("使用說明：")
			fmt.Println("  -v                 顯示版本號")
			fmt.Println("  help               顯示本說明")
			os.Exit(0)
		}
	}

	flag.Parse()

	if *showVersion {
		fmt.Printf("anigi 版本： %v \n", config.Version)
		os.Exit(0)
	}

	os.Chdir(*workDir)

	config := viper.New()
	config.AddConfigPath(*workDir)
	config.SetConfigName("config")
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Printf("找不到 config.json \n")
		} else {
			fmt.Printf("讀取 config.json 時發生錯誤： %v \n", err)
		}
	}

	jc := Config{}
	if err := config.Unmarshal(&jc); err != nil {
		fmt.Printf("解析 config.json 時發生錯誤： %v \n", err)
	}

	fmt.Printf("config: %v \n", jc)

	ctx, cancel := context.WithCancel(context.Background())

	a , err:= anigi.NewAnigi(ctx,jc.Anigi)
	if err != nil {
		fmt.Printf("初始化 anigi 時發生錯誤： %v \n", err)
		return
	}
	fmt.Printf("anigi 初始化完成，開始運行... \n")
	a.Run()
	cancel()
	fmt.Printf("anigi 已經結束運行。 大家掰掰 \n")

}
