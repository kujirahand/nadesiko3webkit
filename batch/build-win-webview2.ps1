# nadesiko3 webview2 build script for Windows (PowerShell version)

$ScriptDir = $PSScriptRoot
$RootDir = Split-Path -Path $ScriptDir -Parent
$TargetDir = Join-Path $ScriptDir "win-nako3-webview2"

Write-Host "--------------------------------------------------------------"
Write-Host "ROOT_DIR=$RootDir"
Write-Host "SCRIPT_DIR=$ScriptDir"
Write-Host "TARGET_DIR=$TargetDir"
Write-Host "--------------------------------------------------------------"

# --- reset ---
if (Test-Path $TargetDir) {
    Remove-Item -Recurse -Force $TargetDir
}
New-Item -ItemType Directory -Force -Path $TargetDir | Out-Null

# --- build ---
Copy-Item (Join-Path $ScriptDir "res\nadesiko3webkit.syso") (Join-Path $RootDir "nadesiko3webkit.syso") -Force

# Run build in the root directory
$OriginalDir = Get-Location
Set-Location $RootDir
try {
    go build -ldflags="-H windowsgui" -o (Join-Path $TargetDir "nadesiko3.exe")
}
finally {
    Set-Location $OriginalDir
}

# Copy webapp resources recursively
Copy-Item -Path (Join-Path $RootDir "webapp") -Destination $TargetDir -Recurse -Force

# Copy README.md
Copy-Item (Join-Path $ScriptDir "res\README.md") $TargetDir -Force

Write-Host "ok"
Read-Host -Prompt "Press Enter to continue..."
