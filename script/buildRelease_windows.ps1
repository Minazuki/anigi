# buildRelease_windows.ps1
# 編譯 Release 版，不更新版本

& "$PSScriptRoot\buildApp_windows.ps1" -buildType r -updateVersion f -outputZip f
