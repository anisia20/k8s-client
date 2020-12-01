# K8S TEST RESTAPI

- getpods local TEST
  - http://localhost:3000/api/getpods

- command
```command
go mod init getpods
go get -u
go get k8s.io/client-go@v0.19.0
go run cmd/main.go
```

- local test kill
```sh
kill -9 `netstat -avn | grep 3000 | grep LISTEN | awk '{print $9;}'`
```
