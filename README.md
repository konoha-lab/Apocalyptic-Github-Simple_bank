# [Apocalyptic](#apocalyptic)

This is simple bank testing base on https://dev.to/techschoolguru/, 

## Setup local development
Download pre-requirement tools (Windows, MacOS, or Linux)

> Windows

1. [powershell](https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell-on-windows?view=powershell-7.2&viewFallbackFrom=powershell-6)
    **Installation :**

    ```
    https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell-on-windows?view=powershell-7.2&viewFallbackFrom=powershell-6
    ```

2. [scoop](https://scoop.sh/)
**Installation (in powershell admin)**
     
    ```
    Set-ExecutionPolicy RemoteSigned -scope CurrentUser
    Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')
    ```

3. [Golang](https://go.dev/dl/)
**Installation :**

    ```
    https://go.dev/dl/
    ```

4. [Docker desktop](https://www.docker.com/products/docker-desktop)
**Installation :**

    ```
    https://www.docker.com/products/docker-desktop
    ```

5. [TablePlus](https://tableplus.com/)
**Installation :**

    ```
    https://tableplus.com/
    ```

4. [Git](https://git-scm.com/downloads)
**Installation :**

    ```
    https://git-scm.com/downloads
    ```

5. [sqlc](https://sqlc.dev/)
**Installation :**

    ```
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    ```


6. [Make](http://gnuwin32.sourceforge.net/packages/make.htm)
**Installation :**
    > [choco](https://chocolatey.org/install) required :
        Install choco in windows (in powershell admin) :

   ```
   Set-ExecutionPolicy RemoteSigned -scope CurrentUser
   Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1')
   ```

    ```
    choco install make
    ```

7. [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate#installation)
**Installation :**

    ```
    scoop install migrate
    ```

### How to use
- ```$ clone https://github.com/konoha-lab/Apocalyptic-Github-Simple_bank.git```
- ```$ docker pull postgres:12-alpine (if need to use mysql : docker pull mysql:8.0/docker pull mysql:latest)```
- ```$ make postgrescreate / make mysqlcreate```
- ```$ make migrateup```
- ```$ make server```

### Notes
#notes
1.  Requires a version of Go that supports modules. e.g. Go 1.15+
2.  These examples build the with postgres/mysql Database. 

Thank you so much for reading! ☺
