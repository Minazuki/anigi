  
![demo](./image/demo.gif)

# 使用
```
https://github.com/Minazuki/gif2ico/releases
請去隔壁下載工具
把gif轉成ico資訊
ex: 
PS D:\Source\gif2ico_Release_202507251529> .\gif2ico.exe -gif .\leafGlowstick.gif
找到了 20 個影格。 
已將影格 0 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_0.ico 
已將影格 1 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_1.ico
已將影格 2 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_2.ico
已將影格 3 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_3.ico
已將影格 4 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_4.ico
已將影格 5 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_5.ico
已將影格 6 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_6.ico
已將影格 7 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_7.ico
已將影格 8 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_8.ico
已將影格 9 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_9.ico
已將影格 10 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_10.ico
已將影格 11 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_11.ico
已將影格 12 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_12.ico
已將影格 13 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_13.ico
已將影格 14 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_14.ico
已將影格 15 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_15.ico
已將影格 16 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_16.ico
已將影格 17 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_17.ico
已將影格 18 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_18.ico
已將影格 19 儲存為 D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_19.ico
已輸出 ICO JSON: D:\Source\gif2ico_Release_202507251529\leafGlowstick\leafGlowstick_ico.json 
處理完成。
```
```
把產出來的資料夾擺到執行檔旁邊
d-----       2025/7/25  下午 08:18                leafGlowstick
-a----       2025/7/25  下午 08:36        5518336 anigi.exe
-a----       2025/7/25  下午 08:13            103 config.json
-a----       2025/7/25  下午 08:36             34 md5.txt
```
```
確定你的config.json內容是指向對的資料夾名稱
{
  "anigi": {
    "title": "leaf",
    "icon": "leafGlowstick",
    "sourceProvider":"cpu"
  }
}
執行anigi.exe就可以了
```
```
sourceProvider目前支援
cpu: 你現在cpu使用量的平均值
fakecpu: 0~100 每5秒+10%
mem: 你的記憶體使用量

程式本體沒有上鎖 你可以複製多個配合不同圖片資料啟動
```

# 配置
請參考  
.vscode_templet  

# build script  
debug  
```
PS D:\Source\anigi> .\script\buildDebug_windows.ps1  
正在編譯 anigi 為 Debug 版...
編譯成功！輸出檔案：D:\Source\anigi\bin\Debug\anigi.exe
```
release  
```
PS D:\Source\anigi> .\script\buildRelease_windows.ps1
正在編譯 anigi 為 Release 版...                                                                                   編譯成功！輸出檔案：D:\Source\anigi\bin\Release\anigi.exe   
```
distribution
```
PS D:\Source\anigi> .\script\dist_windows.ps1                                                                     正在更新版本號...                                                                                                 已更新版本號至 2025.07.25.2036 (D:\Source\anigi\config\version.go)                                                
正在編譯 anigi 為 Release 版...
編譯成功！輸出檔案：D:\Source\anigi\bin\Release\anigi.exe
正在產生 MD5...
MD5 已產生：D:\Source\anigi\bin\Release\md5.txt
正在打包 zip...
已完成打包：D:\Source\anigi\dist\anigi_Release_202507252036.zip
```
