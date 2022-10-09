package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

var PORT string = ":1000"

func run(c context.Context, port string) error {
	s := http.Server{
		Addr: port,
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
			},
		),
	}
	eg, c := errgroup.WithContext(c)
	eg.Go(func() error {
		log.Printf("server is starting,")
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("fail to close, %s", err.Error())
			return err
		}
		return nil
	})
	<-c.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		fmt.Printf("fail to Shutdown, %s", err.Error())
	}
	return eg.Wait()
}

func main() {
	if err := run(context.Background(), PORT); err != nil {
		log.Printf("fail to terminate, %s", err.Error())
	}
}
