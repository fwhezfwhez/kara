// Should run karad first.
//
/*
	cd /path/to/kara/karad
	go run main.go
*/
// or
/*
    cd /path/to/kara/karad/bin
    karad -httpPort :8080
*/

package test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"io/ioutil"
	"kara/karad/src"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"testing"
)

func TestHTTPSingle(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var spot = "spot_1"
	var key = "single_job_x"
	var url = "http://localhost:8080/kara/single/"
	srv := func(wg *sync.WaitGroup) {
		defer wg.Done()
		var mp = map[string]interface{}{
			"spot": spot,
			"key":  key,
		}
		buf, _ := json.Marshal(mp)
		rsp, e := http.Post(url, "application/json", bytes.NewReader(buf))
		if e != nil {
			fmt.Println(e.Error(), "\n", string(debug.Stack()))
			os.Exit(-1)
			return
		}
		if rsp.StatusCode != 200 {
			defer rsp.Body.Close()
			rs, e := ioutil.ReadAll(rsp.Body)
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(-1)
				return
			}
			fmt.Println(string(rs))
			os.Exit(-1)
			return
		}

		var result struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}
		if rsp != nil && rsp.Body != nil {
			defer rsp.Body.Close()
			rs, e := ioutil.ReadAll(rsp.Body)
			if e != nil {
				fmt.Println(e.Error(), "\n", string(debug.Stack()))
				os.Exit(-1)
				return
			}
			if e := json.Unmarshal(rs, &result); e != nil {
				fmt.Println(e.Error(), "\n", string(debug.Stack()))
				os.Exit(-1)
				return
			}
		}

		ok := result.Status
		if ok {
			fmt.Println("hello, ketty")
		}
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go srv(&wg)
	}

	wg.Wait()
}

func TestHTTPMulti(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var spot = "spot_2"
	var key = "multi_job_x"
	var url = "http://localhost:8080/kara/multiple/"
	var limit = 10
	srv := func(wg *sync.WaitGroup) {
		defer wg.Done()
		var mp = map[string]interface{}{
			"spot":  spot,
			"key":   key,
			"limit": limit,
		}
		buf, _ := json.Marshal(mp)
		rsp, e := http.Post(url, "application/json", bytes.NewReader(buf))
		if e != nil {
			fmt.Println(e.Error(), "\n", string(debug.Stack()))
			os.Exit(-1)
			return
		}
		if rsp.StatusCode != 200 {
			defer rsp.Body.Close()
			rs, e := ioutil.ReadAll(rsp.Body)
			if e != nil {
				fmt.Println(e.Error())
				os.Exit(-1)
				return
			}
			fmt.Println(string(rs))
			os.Exit(-1)
			return
		}

		var result struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}
		if rsp != nil && rsp.Body != nil {
			defer rsp.Body.Close()
			rs, e := ioutil.ReadAll(rsp.Body)
			if e != nil {
				fmt.Println(e.Error(), "\n", string(debug.Stack()))
				os.Exit(-1)
				return
			}
			if e := json.Unmarshal(rs, &result); e != nil {
				fmt.Println(e.Error(), "\n", string(debug.Stack()))
				os.Exit(-1)
				return
			}
		}

		ok := result.Status
		if ok {
			fmt.Println("hello, ketty")
		}
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go srv(&wg)
	}

	wg.Wait()
}

func TestGRPCSingle(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var spot = "spot_3"
	var key = "single_job_x"
	var url = "localhost:8081"

	conn, e := grpc.Dial(url, grpc.WithInsecure())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer conn.Close()
	c := src.NewJobServiceClient(conn)

	srv := func(wg *sync.WaitGroup) {
		defer wg.Done()

		r, e := c.SingleTimesJob(context.Background(), &src.SingleTimesJobRequest{
			SpotId: spot,
			Key:    key,
		})
		if e != nil {
			fmt.Println(e.Error())
			return
		}

		ok := r.Status
		if ok {
			fmt.Println("hello, ketty")
		}
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 11; i++ {
		wg.Add(1)
		go srv(&wg)
	}

	wg.Wait()
}

func TestGRPCMulti(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var spot = "spot_4"
	var key = "multiple_job_y"
	var url = "localhost:8081"

	conn, e := grpc.Dial(url, grpc.WithInsecure())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer conn.Close()
	c := src.NewJobServiceClient(conn)

	srv := func(wg *sync.WaitGroup) {
		defer wg.Done()

		r, e := c.MultipleTimesJob(context.Background(), &src.MultipleTimesJobRequest{
			SpotId: spot,
			Key:    key,
			Limit:  10,
		})
		if e != nil {
			fmt.Println(e.Error())
			return
		}

		ok := r.Status
		if ok {
			fmt.Println("hello, ketty")
		}
	}

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go srv(&wg)
	}

	wg.Wait()
}
