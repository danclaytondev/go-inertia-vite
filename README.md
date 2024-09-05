# Go + Inertia.js Demo (with Vite)

This is a demo application using a golang backend, with React frontend connected using Inertia.js. Vite is used to build and bundle the frontend assets. The Inertia.js logic is provided by [romsar/gonertia](https://github.com/romsar/gonertia).
## How to run

### Development mode

With the use of Vite, we need to run one process for Vite, and another to run the Go backend. That is actually very useful as we get all the benefits of Hot Module Reloading with Vite which makes developing the frontend code a breeze.

To start, first run the development tasks for frontend:
```bash
$ cd frontend
$ pnpm dev
```
Vite should start running at `http://localhost:5173`. If this is running at another address (e.g. because you've configured it to), you'll need to update the Go code that generates the Vite script tags.

Then you can start the Go server **in a second terminal**:
```bash
$ go run main.go -dev
```
You could use [air](https://github.com/air-verse/air) for live reloading of the Go code. 

The `dev` flag is used to instruct the code that generated Vite tags, whether to link to the development Vite server or whether to link to the built assets that Go will serve.

### Building for Production 

First you'll want to build the frontend assets.
```bash
$ cd frontend
$ pnpm build
```

Then you can build you Go app. There are different ways to deploy the Go app; if you are producing a single binary, it would be easy to embed the `frontend/dist/` folder using the Go `embed` package.


## How Vite is integrated

Vite is integrated using the [olivere/vite](https://github.com/olivere/vite) package, which acts as a 'helper function' to generate HTML tags for links to JS/CSS/static assets. We must tell this library where our vite files are located, what URL the Vite dev server will run on, and whether we want it to link to the dev server or our production files.