# Todo Application (Server)

## Useful commands
$ docker pull postgres
$ docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
$ docker ps
$ migrate create -ext sql -dir ./schema -seq init
$ migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
$ docker exec -it <CONTAINER_ID> /bin/bash
root@<CONTAINER_ID>:/# psql -U postgres
postgres=# \d
postgres=# select * from users;
postgres=# exit
root@<CONTAINER_ID>:/# exit