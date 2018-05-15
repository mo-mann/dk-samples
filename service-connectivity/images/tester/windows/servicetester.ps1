#!Powershell.exe
while($true) {
    echo "*********************************************"
    echo "[$(Get-Date)]: Checking the URL $env:SERVICE_URL"
    echo (Invoke-WebRequest -UseBasicParsing $env:SERVICE_URL).Content
    echo "[$(Get-Date)]: Waiting 5 seconds"
    Start-Sleep -Seconds 5
}