package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

/**
* @description:
 基于 errgroup 实现一个 http server 的启动和关闭 ，
 以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。


* @param  {*}
* @return {*}
*/
func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	egroup, errCtx := errgroup.WithContext(ctx)
	server := &http.Server{Addr: "8080"}
	egroup.Go(func() error {
		return StartHttp(server)
	})

	defer cancel()

	egroup.Go(func() error {
		<-errCtx.Done()
		fmt.Println("server stop")
		return server.Shutdown(errCtx)
	})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	egroup.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-ch:
				cancel()
			}
		}
	})

	if err := egroup.Wait(); err != nil {
		fmt.Println("errgroup err:", err)
	}
	fmt.Println("Done")

}

func StartHttp(server *http.Server) error {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	})

	fmt.Println("server start")
	err := server.ListenAndServe()
	return err
}
