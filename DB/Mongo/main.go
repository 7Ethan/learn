// package main

// import (
// 	"gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// 	"github.com/astaxie/beego/logs"
// 	"fmt"
// )

// type Person struct{
// 	Name string
// 	Phone string
// }
// func main(){
// 	session,err := mgo.Dial("localhost:27017")
// 	if err != nil{
// 		logs.Error(err)
// 		panic(err)
// 	}
// 	defer session.Close()

// 	session.SetMode(mgo.Monotonic,true)session.SetMode(mgo.Monotonic,true)

// 	c := session.DB("test").C("person")

// 	err = c.Insert(&Person{"Sergey","8888888"},&Person{"shangan","0000000000"})

// 	if err != nil{
// 		logs.Error(err)
// 		panic(err)
// 	}

// 	result := Person{}

// 	err = c.Find(bson.M{"name" : "Sergey"}).One(&result)

// 	if err != nil{
// 		logs.Error(err)
// 		panic(err)
// 	}

// 	fmt.Println("name : ",result.Name)
// 	fmt.Println("phone : ",result.Phone)

// }

/*----------------------------------------
	** Microservice with MongoDB in Go **
----------------------------------------*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"goji.io"

	"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	ISBN    string   `json:"isbn"`
	Title   string   `json:"title"`
	Authors []string `json:"Authors"`
	Price   string   `json:"price"`
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message : %q}", message)
}

func ResponseWithJson(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/books"), allBooks(session))
	mux.HandleFunc(pat.Post("/books"), addBook(session))
	mux.HandleFunc(pat.Get("/books/:isbn"), bookByISBN(session))
	mux.HandleFunc(pat.Delete("/books/:isbn"), deleteBook(session))
	mux.HandleFunc(pat.Put("/books/:isbn"), updateBook(session))
	http.ListenAndServe("localhost:8080", mux)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("store").C("books")

	index := mgo.Index{
		Key:        []string{"isbn"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func allBooks(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("store").C("books")
		var books []Book
		err := c.Find(bson.M{}).All(&books)
		if err != nil {
			ErrorWithJSON(w, "Database Error ", http.StatusInternalServerError)
			log.Fatalln("Failed to find all books ", err)
			return
		}

		responseBody, err := json.MarshalIndent(books, "", " ")
		if err != nil {
			log.Fatalln(err)
		}
		ResponseWithJson(w, responseBody, http.StatusOK)
	}
}

func addBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Copy()

		// isbn := pat.Param(r, "isbn")
		var book Book
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&book)
		if err != nil {
			ErrorWithJSON(w, "Incrrect Body ", http.StatusBadRequest)
			return
		}

		c := session.DB("store").C("books")
		err = c.Insert(book)
		if err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "Book with this ISBN already exits", http.StatusBadRequest)
				return
			}
			ErrorWithJSON(w, "Database Error : ", http.StatusInternalServerError)
			log.Fatalln("Failed to insert book ", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path+"/"+book.ISBN)
		w.WriteHeader(http.StatusCreated)
	}
}

func bookByISBN(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Copy()

		isbn := pat.Param(r, "isbn")
		c := session.DB("store").C("books")
		var book Book
		err := c.Find(bson.M{"isbn": isbn}).One(&book)
		if err != nil {
			ErrorWithJSON(w, "Database Error", http.StatusInternalServerError)
			log.Fatalln("Failed to find the book ", err)
			return
		}

		if book.ISBN == "" {
			ErrorWithJSON(w, "Book not found ", http.StatusNotFound)
			return
		}

		responseBody, err := json.MarshalIndent(book, "", " ")
		if err != nil {
			log.Fatalln(err)
		}

		ResponseWithJson(w, responseBody, http.StatusOK)
	}
}

func updateBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")

		var book Book
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&book)
		if err != nil {
			ErrorWithJSON(w, "Incorrect book body", http.StatusBadRequest)
			return
		}

		c := session.DB("store").C("books")
		err = c.Update(bson.M{"isbn": isbn}, &book)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database Error ", http.StatusInternalServerError)
				log.Fatalln("Failed to update book ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Database not found", http.StatusNotFound)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteBook(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		isbn := pat.Param(r, "isbn")

		c := session.DB("store").C("books")
		err := c.Remove(bson.M{"isbn": isbn})
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error ", http.StatusInternalServerError)
				log.Panicln("Failed to delete book : ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Database not found ", http.StatusNotFound)
				return
			}
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
