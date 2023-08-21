// package main

// import (
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"text/template"
// 	"time"
// )

// const version = "1.0.0"
// const cssVersion = "1"

// type config struct {
// 	port int
// 	env  string
// 	api  string
// 	db   struct {
// 		dsn string
// 	}
// 	stripe struct {
// 		secret string
// 		key    string
// 	}
// }

// type application struct {
// 	config        config
// 	infoLog       *log.Logger
// 	errorLog      *log.Logger
// 	templateCache map[string]*template.Template
// 	version       string
// }

// func (app *application) serve() error {
// 	srv := &http.Server{
// 		Addr:        fmt.Sprintf(":%d", app.config.port),
// 		Handler:     app.routes(),
// 		IdleTimeout: 30 * time.Second,
// 		ReadTimeout: 30 * time.Second,
// 	}
// }

// func main() {
// 	var cfg config
// 	flag.IntVar(&cfg.port, "port", 4000, "Server port in listen on")
// 	flag.StringVar(&cfg.env, "env", "development", "Application env {development | production}")
// 	flag.Parse()

// 	cfg.stripe.key = os.Getenv("STRIPE_KEY")
// 	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

// 	infoLog := log.New(os.Stdout, "INFO\t ", log.Ldate|log.Ltime)
// 	errorLog := log.New(os.Stdout, "ERROR\t ", log.Ldate|log.Ltime|log.Lshortfile)

// }
