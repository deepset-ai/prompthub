package api

import (
	"context"
	fmt "fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/deepset-ai/prompthub/index"
	"github.com/deepset-ai/prompthub/output"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/spf13/viper"
)

// Serve starts the HTTP server and blocks
func Serve() {
	r := chi.NewRouter() // root router
	r.Use(render.SetContentType(render.ContentTypeJSON))

	promptsRouter := chi.NewRouter()
	promptsRouter.Get("/", ListPrompts)
	promptsRouter.Get("/*", GetPrompt)

	r.Mount("/prompts", promptsRouter)

	// Start the HTTP server
	srv := &http.Server{
		Addr: fmt.Sprintf("0.0.0.0:%s", viper.GetString("port")),
		// Set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			output.FATAL.Println(err)
		}
	}()

	output.INFO.Println("Prompthub running at", srv.Addr)

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	output.INFO.Println("shutting down")
}

func ListPrompts(w http.ResponseWriter, r *http.Request) {
	prompts := index.GetPrompts()
	if err := render.RenderList(w, r, NewPromptListResponse(prompts)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
}

func NewPromptListResponse(prompts []*index.Prompt) []render.Renderer {
	list := []render.Renderer{}
	for _, prompt := range prompts {
		list = append(list, NewPromptResponse(prompt))
	}
	return list
}

func GetPrompt(w http.ResponseWriter, r *http.Request) {
	promptName := chi.URLParam(r, "*")
	prompt, err := index.GetPrompt(promptName)
	if err != nil {
		render.Render(w, r, ErrNotFound)
		return
	}

	if err := render.Render(w, r, NewPromptResponse(prompt)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}
