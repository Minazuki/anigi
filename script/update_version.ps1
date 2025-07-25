# update_version.ps1
# 功能：自動更新 config\version.go 的版本號為當前時間 (YYYY.MM.DD.HHmm)

$ErrorActionPreference = "Stop"
Import-Module "$PSScriptRoot\utilty.ps1" -Force

$projectRoot = (Get-Location).Path
$versionFile = Join-Path $projectRoot "config\version.go"

if (!(Test-Path $versionFile)) {
    Write-RedToConsole "找不到 version.go，請確認路徑是否正確：$versionFile"
    exit 1
}

$now = Get-Date -Format "yyyy.MM.dd.HHmm"

$content = @"
package config

// Version of anigi
const Version = "$now"
"@

Set-Content -Path $versionFile -Value $content -Encoding UTF8
Write-GreenToConsole "已更新版本號至 $now ($versionFile)"
exit 0