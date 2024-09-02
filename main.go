package main

import (
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"time"

	"github.com/olivere/vite"
	inertia "github.com/romsar/gonertia"
)

var i *inertia.Inertia

func main() {
	var (
		isDev = flag.Bool("dev", false, "run in development mode")
	)
	flag.Parse()

	viteFragment, err := vite.HTMLFragment(vite.Config{
		FS:      os.DirFS("frontend/dist"),
		IsDev:   *isDev,
		ViteURL: "http://localhost:5173",
	})
	if err != nil {
		panic(err)
	}

	i, err = inertia.NewFromFile("frontend/index.tmpl")
	if err != nil {
		panic(err)
	}

	i.ShareTemplateData("Vite", viteFragment)

	mux := http.NewServeMux()

	if *isDev {
		serverStaticFolder(mux, "/src/assets/", os.DirFS("frontend/src/assets"))
	} else {
		serverStaticFolder(mux, "/assets/", os.DirFS("frontend/dist/assets"))
	}

	endpoints := map[string]http.HandlerFunc{
		"/{$}":   homeHandler,
		"/users": usersHandler,
	}

	for endpoint, f := range endpoints {
		mux.Handle(endpoint, i.Middleware(http.HandlerFunc(f)))
	}

	// Start a listener.
	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(300 * time.Millisecond)
	err := i.Render(w, r, "Home", inertia.Props{
		"user": "data",
	})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(300 * time.Millisecond)
	err := i.Render(w, r, "Users/Index", inertia.Props{
		"users": []string{"Dan", "Mark"},
	})

	if err != nil {
		panic(err)
	}
}

func serverStaticFolder(mux *http.ServeMux, path string, fs fs.FS) {
	mux.Handle(path, http.StripPrefix(path, http.FileServerFS(fs)))
}
