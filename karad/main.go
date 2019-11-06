package main

import (
	"flag"
	"fmt"
	"github.com/fwhezfwhez/kara"
	"github.com/fwhezfwhez/kara/karad/src"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"time"
)

var httpPort string
var grpcPort string

var (
	version = kara.Version
)

const (
	help = `
# Kara
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
	fmt.Println(os.Args)
	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "version", "--version":
			fmt.Println(version)
			return
		case "help":
			fmt.Println(help)
			return
		case "cli":
			if args[1] == "ping" {
				//ping()
			}
			return
		}
	}

	if httpPort != "" && httpPort != "disable" {
		go karaHttp()
	}
	if grpcPort != "" && grpcPort != "disable" {
		go karaGrpc()
	}
	fmt.Println(kara.GetLogo(httpPort, grpcPort))
	select {}
}
func karaHttp() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
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
	select {}
}
