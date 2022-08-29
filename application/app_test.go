package application

import (
	"testing"
	"time"

	"github.com/go-warrior/pkg/transport/grpc"
	"github.com/go-warrior/pkg/transport/http"
)

func TestApp(t *testing.T) {
	hs := http.NewServer()
	gs := grpc.NewServer()
	app := New(
		Name("warrior"),
		Version("v0.0.1"),
		Server(hs, gs),
	)
	time.AfterFunc(time.Second, func() {
		app.Stop()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
