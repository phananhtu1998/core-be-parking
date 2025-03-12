@echo off
setlocal

:: Định nghĩa biến môi trường
set GOOSE_DRIVER=mysql
set GOOSE_DBSTRING=root:root1234@tcp(127.0.0.1:3306)/parkingdevgo
set GOOSE_MIGRATION_DIR=sql/schema
set APP_NAME=server
set SWAGTOMAIN=cmd/server/main.go
set SWAGOUT=cmd/swag/docs

:: Kiểm tra tham số đầu vào
if "%1"=="" (
    echo Usage: makefile.bat [dev|docker_build|docker_stop|docker_up|upse|downse|resetse|up-by-one|create_migration name=migration_name]
    exit /b
)

:: Xử lý tham số đầu vào
if "%1"=="dev" goto RUN_APP
if "%1"=="docker_build" set BUILD_FLAG=--build && goto DOCKER_UP
if "%1"=="docker_up" goto DOCKER_UP
if "%1"=="docker_stop" goto DOCKER_STOP
if "%1"=="sqlgen" goto SQLGEN
if "%1"=="create_migration" goto CREATE_MIGRATION
if "%1"=="upse" set ACTION=up && goto RUN_GOOSE
if "%1"=="downse" set ACTION=down && goto RUN_GOOSE
if "%1"=="resetse" set ACTION=reset && goto RUN_GOOSE
if "%1"=="up_by_one" set ACTION=up-by-one && goto RUN_GOOSE
if "%1"=="swag" goto SWAG

:: Hiển thị hướng dẫn nếu tham số không hợp lệ
echo Invalid option: %1
echo Usage: makefile.bat [dev|swag|docker_build|docker_stop|docker_up|upse|downse|resetse|up-by-one|create_migration name=migration_name]
exit /b

:DOCKER_UP
echo Running Docker Compose...
docker compose up -d %BUILD_FLAG%
docker compose ps
exit /b

:DOCKER_STOP
echo Stopping Docker...
docker compose down
exit /b

:RUN_GOOSE
echo Running goose %ACTION%...
goose -dir=%GOOSE_MIGRATION_DIR% %ACTION% %GOOSE_DRIVER% "%GOOSE_DBSTRING%"
exit /b

:CREATE_MIGRATION
if "%2"=="" (
    echo Missing migration name. Usage: makefile.bat create_migration name=migration_name
    exit /b
)
set NAME=%2
echo Running goose create migration: %NAME%...
goose -dir=%GOOSE_MIGRATION_DIR% create %NAME% sql
exit /b

:SQLGEN
echo generating SQL
sqlc generate
exit /b

:SWAG
swag init -g %SWAGTOMAIN% -o %SWAGOUT% 
exit /b

:RUN_APP
echo Running the application...
go run ./cmd/%APP_NAME%
exit /b
