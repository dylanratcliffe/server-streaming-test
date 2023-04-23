package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/dylanratcliffe/server-streaming-test/sst"
	"github.com/dylanratcliffe/server-streaming-test/sst/sstconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type CountServer struct{}

func (c *CountServer) CountTo(ctx context.Context, req *connect_go.Request[sst.CountToRequest], res *connect_go.ServerStream[sst.CountToResponse]) error {
	fmt.Printf("Counting to %v\n", req.Msg.GetTo())

	var count int32

	// Create a ticker
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			count++

			fmt.Printf("Sending count %v\n", count)

			// Send the current count
			res.Send(&sst.CountToResponse{
				Count: count,
			})

			if count == req.Msg.GetTo() {
				return nil
			}
		}
	}
}

func main() {
	fmt.Printf("Starting count server...\n")

	counter := &CountServer{}
	mux := http.NewServeMux()
	path, handler := sstconnect.NewCounterServiceHandler(counter)
	mux.Handle(path, handler)
	http.ListenAndServe(":8080", h2c.NewHandler(mux, &http2.Server{}))
}
