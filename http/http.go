package http

import (
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

type MacroForm struct {
	Player1 string
	Player2 string
}

var formTemplate *template.Template

func init() {
	tmpl, _err := template.New("formHtml").Parse(formHtml())
	if _err != nil {
		log.Println("Could not parse template.")
		log.Fatal(_err)
	}
	formTemplate = tmpl
}

func handle(requestHandler func(form *url.Values) string) (mux *http.ServeMux) {
	mux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()
			msg := requestHandler(&r.PostForm)
			if len(msg) > 0 {
				fmt.Fprintf(w, msg)
			} else {
				fmt.Fprintf(w, "OK")
			}
		} else {
			_err := formTemplate.Execute(w, MacroForm{"", ""})
			if _err != nil {
				log.Println("Could not execute template.")
				log.Println(_err)
			}
		}
	})

	return
}

func Server(requestHandler func(form *url.Values) string) {
	ch := make(chan error)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	go func() {
		mux := handle(requestHandler)
		ch <- http.Serve(listener, mux)
	}()

	fmt.Println("Server started at", listener.Addr())

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)
	go func() {
		log.Println(<-sig)
		listener.Close()
	}()

	log.Println(<-ch)
}
