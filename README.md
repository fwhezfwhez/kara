## kara
Kara is a middleware server.It has aimed goals on different version stage.

v1.x.x
- Solving multiple servers share the same code and execute a job specific times limit.

v2.x.x
- Imitating redis and provide grpc/http access it, making client more efficiently programming.

v3.x.x
- Making kara cloud-able.

## 1. Start
** Using GOPATH **
```go
cd ${GOPATH}/src
git clone https://github.com/fwhezfwhez/kara.git
cd kara/karad
karad -httpPort :8080 -grpcPort :8081
```
or

```go
go get -u github.com/fwhezfwhez/kara
cd ${GOPATH}/src/kara/karad
karad -httpPort :8080 -grpcPort :8081
```

** Using GOMODULE **
```go
git clone https://github.com/fwhezfwhez/kara.git
cd kara/karad
karad -httpPort :8080 -grpcPort :8081
```
**To verify its heath**
```go
>> karad cli ping
pong
```

## 2. Usage

#### 2.1 HTTP

#### 2.2 GRPC

