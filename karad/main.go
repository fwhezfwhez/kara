package karad

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	"kara/karad/src"
	"net/http"
	"time"
)

var httpPort string
var grpcPort string

func main() {
	flag.StringVar(&httpPort, "httpPort", ":8080", "kara -httpPort :8080")
	flag.StringVar(&grpcPort, "grpcPort", ":8081", "kara -grpcPort :8081")
	flag.Parse()

	go myHttp()
	go myGrpc()
}
func myHttp() {
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
func myGrpc() {

}
