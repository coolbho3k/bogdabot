package bogdabot

import (
	"bogdabot/store/commands"
	"context"
	"fmt"
	"go.uber.org/fx"
	"net/http"
	"path"
	"strings"
)

// Service serves responses to commands.
type Service interface {
	Handle(w http.ResponseWriter, req *http.Request)
}

type bogdabotService struct {
	store commands.Store
}

func (p *bogdabotService) Handle(w http.ResponseWriter, req *http.Request) {
	requestPath := strings.Split(path.Clean(req.URL.Path), "/")
	if len(requestPath) != 2 {
		http.Error(w, "could not parse path", 400)
		return
	}

	if response, err := p.store.GetResponseByPath(requestPath[1]); err != nil {
		http.Error(w, "could not find path", 404)
	} else {
		if req.Method != "POST" {
			http.Error(w, "method not allowed", 405)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if _, err := fmt.Fprintf(w, response); err != nil {
			http.Error(w, "could not write body", 400)
		}
	}
}


func New(lc fx.Lifecycle, store commands.Store) (Service, error) {
	bogdabotService := &bogdabotService{store}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			bogdabotService.Serve()
			return nil
		},
	})

	return bogdabotService, nil
}

func (p *bogdabotService) Serve() {
	http.HandleFunc("/", p.Handle)
	go func() {
		_ = http.ListenAndServe(":8080", nil)
	}()
}
