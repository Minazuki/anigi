# utilty.ps1

function Write-HorisontalLine {
    $width = [Math]::Min($Host.UI.RawUI.WindowSize.Width - 1, 100)
    Write-Host ('=' * $width) -ForegroundColor Green
}

function Write-YellowToConsole ([string]$message) {
    Write-Host "$message" -ForegroundColor Yellow
}

function Write-RedToConsole ([string]$message) {
    Write-Host "$message" -ForegroundColor Red
}

function Write-GreenToConsole ([string]$message) {
    Write-Host "$message" -ForegroundColor Green
}

function Write-WhiteToConsole ([string]$message) {
    Write-Host "$message" -ForegroundColor White
}