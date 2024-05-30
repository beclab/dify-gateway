# __**Dify Gateway**__
## Description
**Dify Gateway** acts as a bridge, establishing a connection between **Dify** and other services such as **Files** and **Agent**.

**Dify Gateway** encompasses two main sets of functionalities:
1. **Watcher**

Depending on the configuration, it monitors specific folders in the **files** service for each identified dataset in **Dify**.

When changes occur in a folder, the modified contents are sent to the corresponding datasets for future utilization by the **agent** service.

2. **Gateway**

This package comprises **Dify** APIs that form the core for fulfilling our requirements.

Additionally, it offers management of user-agent relationships, a feature that is not present in the original **Dify**.

## Getting Started
You can run **Dify Gateway** locally, but it will only function correctly when integrated with **Dify**.

About **Dify**: https://github.com/beclab/dify

### Clone the repository
```shell
git clone --recursive https://github.com/beclab/search2.git
```
### Build
```shell
cd search2/gateway
go mod tidy
go build -o wzinc
```
### Set OS Environments
```shell
export WATCH_DIR=/path/for/watching
export DIFY_ADMIN_USER_EMAIL=admin_username@example.com
export DIFY_ADMIN_USER_PASSWORD=*******
export DIFY_USER_NAME=username
export DIFY_USER_PASSWORD=*************
export PG_USERNAME=postgres
export PG_PASSWORD=*********
export PG_HOST=localhost
export PG_PORT=5432
export PG_DATABASE=dify
export DIFY_HOST=http://localhost
```
### Run
```shell
./wzinc start
```
