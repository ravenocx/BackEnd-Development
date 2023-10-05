package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type Server struct {
	*mux.Router

	shoppingItem []Item
}

func NewServer() *Server {
	s := &Server{
		Router : mux.NewRouter(), // create new http mux router instance
		//  a "route" refers to the URL path and the associated HTTP method that a client
		shoppingItem: []Item{},
	}
	s.routes()
	return s
}
func (s *Server) routes(){
	s.HandleFunc("/shopping-items", s.listShoppingItem()).Methods("GET") // REGISTER NEW ROUTE
	s.HandleFunc("/shopping-items", s.addShoppingItem()).Methods("POST")
	s.HandleFunc("/shopping-items/{id}", s.removeShoppingItem()).Methods("DELETE")
	// 1st parameter -> path for handler
	// 2nd parameter -> handler function / signature
}

func(s *Server) addShoppingItem() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil { // http request -> input untuk add shopping items
			// json -> data
			// NewDecoder(x) (create a new JSON decoder) -> x akan dilakukan decode (input untuk di decode ke target)
			// Decode(y) -> data dari x di parse(decode) dan disimpan di y
			http.Error(w, err.Error(), http.StatusBadRequest) // request -> StatusBadRequest
			return
		}
		i.ID = uuid.New()
		s.shoppingItem = append(s.shoppingItem, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil { // http response -> yang ditampilin 
			// NewEncoder(x) -> x adalah target yang ingin dilakukan Encode
			// Encode(i) -> i adalah yang akan di Encode
			http.Error(w , err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listShoppingItem() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.shoppingItem);err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError) // response -> StatusInternalServerError
			return
		}
	}
}

func (s *Server) removeShoppingItem() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		idStr:= mux.Vars(r)["id"] // returns the route variables for the current request
		id, err := uuid.Parse(idStr) // decode uuid (json -> data)
		if err != nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.shoppingItem{
			if item.ID == id {
				s.shoppingItem = append(s.shoppingItem[:i],s.shoppingItem[i+1:]... )
				break
			}
		}
	}
}