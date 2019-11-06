package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"kara/karad/src"
	"net"
	"net/http"
	"os"
	"time"
)

var httpPort string
var grpcPort string

const (
	version = "v1.0.0"
	help    = `
karad -httpPort disable                         # will not start http server
karad -grpcPort disable                         # will not start grpc server
karad version                                   # will print its version
karad cli <command>                             # will send cli request to execute commands
`
)

func main() {
	flag.StringVar(&httpPort, "httpPort", ":8080", "kara -httpPort :8080")
	flag.StringVar(&grpcPort, "grpcPort", ":8081", "kara -grpcPort :8081")
	flag.Parse()

	args := os.Args
	switch args[0] {
	case "version", "--version":
		fmt.Println(version)
		return
	case "help":
		fmt.Println(help)
		return
	case "cli":
		if args[1] == "ping" {
			fmt.Println("pong")
		}
		return
	}

	if httpPort != "" && httpPort != "disable" {
		go karaHttp()
	}
	if grpcPort != "" && grpcPort != "disable" {
		go karaGrpc()
	}

	select {}
}
func karaHttp() {
	r := gin.Default()
	r.POST("/kara/single/", src.SingleJob)
	r.POST("/kara/multiple/", src.MultipleJob)
	s := &http.Server{
		Addr:           httpPort,
		Handler:        cors.AllowAll().Handler(r),
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 21,
	}
	s.ListenAndServe()
}
func karaGrpc() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	s := grpc.NewServer()
	src.RegisterJobServiceServer(s, &src.JobService{})

	go func() {
		s.Serve(lis)
	}()
	fmt.Println(s.GetServiceInfo())
	select {}
}
