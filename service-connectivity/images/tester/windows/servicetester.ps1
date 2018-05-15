#!Powershell.exe
while($true) {
    echo "*********************************************"
    echo "[$(Get-Date)]: Checking the URL $env:SERVICE_URL"
    echo (Invoke-WebRequest -UseBasicParsing $env:SERVICE_URL).Content
    echo "[$(Get-Date)]: Waiting $SERVICE_DELAY"
    Start-Sleep -Seconds 5
}