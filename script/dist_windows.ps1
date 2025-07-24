# dist_windows.ps1
# 編譯 Release 版並更新版本

& "$PSScriptRoot\buildApp_windows.ps1" -buildType r -updateVersion t -outputZip t
