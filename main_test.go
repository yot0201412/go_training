package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func Test_run(t *testing.T) {
	ctx, _ := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx, ":2000")
	})
	in := "message"

	rsp, err := http.Get("http://localhost:2000/" + in)
	if err != nil {
		t.Errorf("fail to get, %+v", err)
	}
	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Errorf("fail to get, %+v", err)
	}
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("diffrent response, want: %s, got: %s", want, got)
	}
	defer rsp.Body.Close()
}
