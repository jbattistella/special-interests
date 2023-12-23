package server

import (
	// "encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/jbattistella/special-interests/engine"
	_"gorm.io/gorm"
)


func qCauseHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/homepage.html")
	fmt.Println("HOMEPAGE")
}

func getPromptMsgHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("/prompt called")

	type Reply struct {
		Prompt       string
	}

	type ErrPage struct {
		Message error
	}

	res, preposition, err := engine.Engine()
	if err != nil {
		er := ErrPage{Message: err}
		t, _ := template.ParseFiles("html/errpage.html")
		if err := t.Execute(w, er); err != nil {
			log.Fatal(err)
		}
		return
	}

	var rep Reply
	msg := res.Group + " "+ preposition + " " + res.Cause

	rep = Reply{Prompt: msg}

	t, _ := template.ParseFiles("html/prompt.html")
	if err := t.Execute(w, rep); err != nil {
		log.Fatal(err)
	}
}


func QuesitonableCauseAPI() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// DB, err := database.ConnectDB()
	// if err != nil {
	// 	log.Fatal()
	// }

	// a := DataStore{db: DB}

	r := mux.NewRouter()

	r.HandleFunc("/prompt", getPromptMsgHandler).Methods("GET")
	r.HandleFunc("/", qCauseHome).Methods("GET")

	port := os.Getenv("PORT")

	srv := &http.Server{
        Handler:      r,
        Addr:         "0.0.0.0:" + port,
        // Good practice: enforce timeouts for servers you create!
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }


	//database
	// r.HandleFunc("/vegetables/all", a.getVegetables).Methods("GET")
	// r.HandleFunc("/vegetables/{name}", a.getVegetable).Methods("GET")
	// r.HandleFunc("/vegetables", a.createVegetable).Methods("POST")
	// r.HandleFunc("/vegetables", a.updateVegetable).Methods("PUT")
	// r.HandleFunc("/vegetables/{name}", a.deleteVegetable).Methods("DELETE")

	log.Println("Listening on port 8080")
	log.Fatal(srv.ListenAndServe())
}
