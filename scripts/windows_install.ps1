Write-Host "Downloading axon..."

$url = "https://github.com/Shravan-1908/axon/releases/latest/download/axon-windows-amd64.exe"

$dir = $env:USERPROFILE + "\.axon"
$filepath = $env:USERPROFILE + "\.axon\axon.exe"

[System.IO.Directory]::CreateDirectory($dir)
(Invoke-WebRequest -Uri $url -OutFile $filepath)

Write-Host "Adding axon to PATH..."
[Environment]::SetEnvironmentVariable(
    "Path",
    [Environment]::GetEnvironmentVariable("Path", [EnvironmentVariableTarget]::Machine) + ";"+$dir,
    [EnvironmentVariableTarget]::Machine)

Write-Host "axon installation is successful!"
Write-Host "You need to restart your shell to use axon."
