# buildDebug_windows.ps1
# 編譯 Debug 版，不更新版本

& "$PSScriptRoot\buildApp_windows.ps1" -buildType d -updateVersion f -outputZip f
