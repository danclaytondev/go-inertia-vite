package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/olivere/vite"
	inertia "github.com/romsar/gonertia"
	_ "modernc.org/sqlite"
)

var i *inertia.Inertia
var db *sql.DB

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
		"/{$}":       homeHandler,
		"/countries": usersHandler,
	}

	for endpoint, f := range endpoints {
		mux.Handle(endpoint, i.Middleware(http.HandlerFunc(f)))
	}

	// Open database
	db, err = sql.Open("sqlite", "countries.sqlite")
	if err != nil {
		panic(err)
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

	rows, err := db.Query("SELECT name FROM countries order by random() limit 5")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	countries := make([]string, 5)

	index := 0

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		countries[index] = name
		index = index + 1
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	err = i.Render(w, r, "Countries/Index", inertia.Props{
		"countries": countries,
	})
	if err != nil {
		panic(err)
	}
}

func serverStaticFolder(mux *http.ServeMux, path string, fs fs.FS) {
	mux.Handle(path, http.StripPrefix(path, http.FileServerFS(fs)))
}
