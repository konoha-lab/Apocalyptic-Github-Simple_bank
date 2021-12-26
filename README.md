### Apocalyptic

This is simple bank testing base on https://dev.to/techschoolguru/, 

## Setup local development
Download pre-requirement tools (Windows, MacOS, or Linux)

> Windows

1. powershell
**Installation : 
https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell-on-windows?view=powershell-7.2&viewFallbackFrom=powershell-6

2. scoop
** Install scoop (in powershell admin)
    Set-ExecutionPolicy RemoteSigned -scope CurrentUser
    Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')

3. Golang
** Installation :
    https://go.dev/dl/

4. Docker desktop
** Installation :
    https://www.docker.com/products/docker-desktop

5. TablePlus
** Installation :
    https://tableplus.com/

4. Git
** Installation :
    https://git-scm.com/downloads

5. sqlc
** Installation :
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest


6. Make
** Installation :
    > choco required :
        Install choco in windows (in powershell admin) :
            Set-ExecutionPolicy RemoteSigned -scope CurrentUser
            Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1')

    choco install make

7. migrate 
** Installation :
    scoop install migrate

How to use :
    1. CMD > clone https://github.com/konoha-lab/Apocalyptic-Github-Simple_bank.git
    2. CMD > docker pull postgres:12-alpine (if need to use mysql : docker pull mysql:8.0/docker pull mysql:latest)
    3. CMD > make postgrescreate
    4. CMD > make migrateup
    5. CMD > make server

Notes
    1.  Requires a version of Go that supports modules. e.g. Go 1.15+
    2.  These examples build the with postgres/mysql Database. 

Thank you so much for reading! ☺