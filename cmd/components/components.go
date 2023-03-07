package components

import (
	"context"
	"fmt"
	"github.com/ServiceWeaver/weaver"
	"log"
	"net/http"
)

type Command struct{}

// Reverser component.
type Reverser interface {
	Reverse(context.Context, string) (string, error)
}

// Implementation of the Reverser component.
type reverser struct {
	weaver.Implements[Reverser]
}

func (r *reverser) Reverse(_ context.Context, s string) (string, error) {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes), nil
}

func (c *Command) Execute() error {
	instance := weaver.Init(context.Background())
	listenerOptions := weaver.ListenerOptions{
		LocalAddress: ":8081",
	}
	listener, err := instance.Listener("weaver-learn-component", listenerOptions)
	if err != nil {
		return err
	}

	reverser, err := weaver.Get[Reverser](instance)
	if err != nil {
		log.Fatal(err)
	}

	// Serve the /hello endpoint.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		reversed, err := reverser.Reverse(r.Context(), r.URL.Query().Get("name"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Hello, %s!\n", reversed)
	})
	err = http.Serve(listener, nil)
	if err != nil {
		return err
	}
	return nil
}
