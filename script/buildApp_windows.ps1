
# buildApp_windows.ps1
param
(
    [string]$buildType = "d",    # r=Release, d=Debug (預設d)
    [string]$updateVersion = "f",   # t=更新版本, f=不更新 (預設f)
    [string]$outputZip = "f"   # t=打包成zip, f=不打包 (預設f
)
Import-Module "$PSScriptRoot\utilty.ps1" -Force
$ErrorActionPreference = "Stop"

$projectRoot = (Get-Location).Path
$appName = "anigi"

# 路徑與 build 參數
if ($buildType.ToLower() -eq "r") {
    $binDir = Join-Path $projectRoot "bin\Release"
    $goBuildFlags = @('-ldflags', '-s -w')
    $zipType = "Release"
}
else {
    $binDir = Join-Path $projectRoot "bin\Debug"
    $goBuildFlags = @()
    $zipType = "Debug"
}

# 是否更新版本
if ($updateVersion -eq "t") {
    Write-YellowToConsole "正在更新版本號..."
    $updateVer = & "$projectRoot\script\update_version.ps1" 2>&1
    if ($LASTEXITCODE -ne 0) {
        Write-RedToConsole "版本更新失敗：$updateVer"
        exit 1
    }
}

# 建立輸出資料夾
if (!(Test-Path $binDir)) {
    New-Item -ItemType Directory -Path $binDir | Out-Null
}

# 編譯 Go 專案
$exePath = "$binDir\$appName.exe"
$md5Path = "$binDir\md5.txt"
Write-YellowToConsole "正在編譯 $appName 為 $zipType 版..."
if ($goBuildFlags.Count -gt 0) {
    go build $goBuildFlags -o $exePath "$projectRoot\main.go"
}
else {
    go build -o $exePath "$projectRoot\main.go"
}

if ($LASTEXITCODE -ne 0) {
    Write-RedToConsole "編譯失敗，請檢查錯誤訊息。"
    exit 1
}
Write-YellowToConsole "編譯成功！輸出檔案：$exePath"

if ($outputZip -eq "t") {
    # 產生 md5.txt
    Write-YellowToConsole "正在產生 MD5..."
    $md5 = Get-FileHash $exePath -Algorithm MD5 | Select-Object -ExpandProperty Hash
    Set-Content -Path $md5Path -Value $md5
    Write-YellowToConsole "MD5 已產生：$md5Path"

    # 打包成 zip
    $now = Get-Date -Format "yyyyMMddHHmm"
    $distDir = Join-Path $projectRoot "dist"
    if (!(Test-Path $distDir)) {
        New-Item -ItemType Directory -Path $distDir | Out-Null
    }
    $zipName = "${appName}_${zipType}_${now}.zip"
    $zipPath = Join-Path $distDir $zipName

    Write-YellowToConsole "正在打包 zip..."
    # 複製 example\config.json 到 bin 目錄，確保壓縮時在根目錄
    $configSrc = Join-Path $projectRoot "example\config.json"
    $configDst = Join-Path $binDir "config.json"
    if (Test-Path $configSrc) {
        Copy-Item $configSrc -Destination $configDst -Force
    }
    Compress-Archive -Path $exePath, $md5Path, $configDst -DestinationPath $zipPath -Force
    if (Test-Path $configDst) {
        Remove-Item $configDst
    }
    Write-YellowToConsole "已完成打包：$zipPath"

}