package simple

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ServiceWeaver/weaver"
)

type Command struct{}

func (c *Command) Execute() error {
	instance := weaver.Init(context.Background())
	listenerOptions := weaver.ListenerOptions{
		LocalAddress: ":8080",
	}
	listener, err := instance.Listener("weaver-learn", listenerOptions)
	if err != nil {
		return err
	}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, err = fmt.Fprintf(w, "Hello, %s!\n", r.URL.Query().Get("name"))
		if err != nil {
			log.Fatal(err)
			// PASS
		}
	})
	err = http.Serve(listener, nil)
	if err != nil {
		return err
	}

	return nil
}
