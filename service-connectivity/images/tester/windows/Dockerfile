FROM microsoft/powershell:6.0.2-nanoserver-1709
LABEL maintainer="Raj Chaudhuri <rajch@hotmail.com>"
ENV SERVICE_URL="http://www.google.com"
COPY windows/servicetester.ps1 C:/app/servicetester.ps1
ENTRYPOINT [ "Pwsh.exe", "-File", "C:/app/servicetester.ps1", "-ExecutionPolicy", "-Bypass" ]


