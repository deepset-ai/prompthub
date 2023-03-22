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
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	// RESTy routes for "articles" resource
	r.Route("/prompts", func(r chi.Router) {
		r.With(paginate).Get("/", ListPrompts)
		// r.Get("/search", SearchPrompts)

		r.Route("/{promptID}", func(r chi.Router) {
			r.Use(PromptCtx)      // Load the *Article on the request context
			r.Get("/", GetPrompt) // GET /prompt/my-prompt
		})
	})

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
	if err := render.RenderList(w, r, NewPromptListResponse(prompts)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// PromptCtx middleware is used to load an Article object from
// the URL parameters passed through as the request. In case
// the Article could not be found, we stop here and return a 404.
func PromptCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var prompt *index.Prompt
		var err error

		if promptName := chi.URLParam(r, "promptName"); promptName != "" {
			prompt, err = index.GetPrompt(promptName)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "prompt", prompt)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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
	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the ArticleCtx
	// middleware. The worst case, the recoverer middleware will save us.
	prompt := r.Context().Value("prompt").(*index.Prompt)

	if err := render.Render(w, r, NewPromptResponse(prompt)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

// Article fixture data
var prompts = []*index.Prompt{
	{Name: "foo"},
	{Name: "bar"},
	{Name: "baz"},
}
