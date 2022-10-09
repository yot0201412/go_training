package main

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func run(c context.Context) error {
	s := http.Server{
		Addr: ":1000",
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "Hello, %s", r.URL.Path)
			},
		),
	}
	eg, c := errgroup.WithContext(c)
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("fail to close, %s", err.Error())
			return err
		}
		return nil
	})
	<-c.Done()
	if err := s.Shutdown(c); err != nil {
		fmt.Printf("fail to Shutdown, %s", err.Error())
	}
	return eg.Wait()
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Printf("fail to terminate, %s", err.Error())
	}
}
