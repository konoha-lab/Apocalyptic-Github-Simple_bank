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

2. [Golang](https://go.dev/dl/)
**Installation :**

    ```
    https://go.dev/dl/
    ```

3. [Docker desktop](https://www.docker.com/products/docker-desktop)
**Installation :**

    ```
    https://www.docker.com/products/docker-desktop
    ```

4. [TablePlus](https://tableplus.com/)
**Installation :**

    ```
    https://tableplus.com/
    ```

5. [Git](https://git-scm.com/downloads)
**Installation :**

    ```
    https://git-scm.com/downloads
    ```

6. [sqlc](https://sqlc.dev/)
**Installation :**

    ```
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
    ```


7. [Make](http://gnuwin32.sourceforge.net/packages/make.htm)
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
    > [scoop](https://scoop.sh/) required :
        Install scoop in windows (in powershell admin) :
     
    ```
    Set-ExecutionPolicy RemoteSigned -scope CurrentUser
    Invoke-Expression (New-Object System.Net.WebClient).DownloadString('https://get.scoop.sh')
    ```

    ```
    scoop install migrate
    ```

### How to use
- ```$ clone https://github.com/konoha-lab/Apocalyptic-Github-Simple_bank.git```
- ```$ docker pull postgres:12-alpine```
- ```$ make postgrescreate```
- ```$ make createdb```
- ```$ make migrateup```
- ```$ make server```

### Hot to generate CRUD
- Add table name inside schema folder (using postgres format or myssql format), see example inside diagram export
- Create CRUD script inside ```db/query/mysql``` or ```db/query/postgres```, see example inside folder
- Change file ```sqlc.yaml``` part ```queries``` & ```engine``` name, according to the file added in the folder ```db/query/mysql``` or ```db/query/postgres``` & what engine you are using (postgres/mysql)
  ```
  version: "1"
    packages: 
      - name: "repository"
    	path: "./db/sqlc"
    	𝑞𝑢𝑒𝑟𝑖𝑒𝑠: "./𝑑𝑏/𝑞𝑢𝑒𝑟𝑦/𝑚𝑦𝑠𝑞𝑙/𝑡𝑟𝑎𝑛𝑠𝑓𝑒𝑟.𝑠𝑞𝑙"
    	schema: "./schema/"
    	𝑒𝑛𝑔𝑖𝑛𝑒: "𝑚𝑦𝑠𝑞𝑙"
    	emit_json_tags: true
    	emit_prepared_queries: false
    	# The setting is called emit_empty_slices, and its default value is false. it option to emit an interface that contains all of the function of the Queries struct.
    	emit_interface: true
    	emit_exact_table_names: false
    	# The setting is called emit_empty_slices, and its default value is false. If we set this value to true, then the result returned by a many query will be an empty slice.
    	emit_empty_slices: true
  ```
- Run ```sqlc generate``` inside command prompt, this will generate file inside ```db\sqlc```
- Copy file ```db\sqlc\models.go``` to ```db\models``` and rename it as type of struct [table-name]
- Copy what inside ```db\sqlc\querier.go```
   ```
    type Querier interface {
        ...
	    CreateTransfer(ctx context.Context, arg CreateTransferParams) (sql.Result, error)
	    DeleteTransfer(ctx context.Context, id int64) error
	    GetTransfer(ctx context.Context, id int64) (Transfer, error)
	    ListTransfer(ctx context.Context, arg ListTransferParams) ([]Transfer, error)
	    UpdateTransfer(ctx context.Context, arg UpdateTransferParams) (sql.Result, error)
        ...
    }
   ```
   paste it inside ```db\repository\querier.go```.
   if using template mysql, change  ```sql.Result``` as return type (type [table-name] struct) for function ```Create[Transfer]``` & ```Update[Transfer]```
- Copy file ```db\sqlc\[table-name].db.go``` into ```db\repository\``` and rename it from ```[table-name].db.go``` to ```[table-name].repository.go```
  if using template mysql, change  ```sql.Result``` as well.
- Run ```mockgen -package mockapi -destination db/mock/handler.go simple_bank/api Handler``` inside command prompt, to generate morkgen.
- Delete what left inside ```db\sqlc```
- Dont forget to create test file inside ```db\test```
- happy coding :grin:

### Notes
1.  Requires a version of Go that supports modules. e.g. Go 1.15+
2.  These examples build the with postgres/mysql Database. 

Thank you so much for reading! ☺
