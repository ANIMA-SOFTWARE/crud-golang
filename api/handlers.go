package api

import (
	"fmt"
	"main/templates"
	"main/types"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

func (s *Server) handleDataGetFirst(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	vars := mux.Vars(r)
	table := s.store.GetTableType(vars["table"])

	tabledata := s.store.GetPage(0, table)

	templates.TableBase(tabledata).Render(r.Context(), w)

}

func (s *Server) handleDataGetPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	vars := mux.Vars(r)
	table := s.store.GetTableType(vars["table"])
	pageNoInt, _ := strconv.Atoi(vars["PageNumber"])

	tabledata := s.store.GetPage(pageNoInt, table)

	templates.TableRows(tabledata).Render(r.Context(), w)
	templates.TableRowNew(tabledata).Render(r.Context(), w)

}

func (s *Server) handleDataGetBySearch(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	r.ParseForm()
	table := s.store.GetTableType(vars["table"])

	searchStr := r.PostForm.Get("search")

	var tabledata types.Table

	if searchStr == "" {
		tabledata = s.store.GetPage(0, table)
	} else {
		tabledata = s.store.GetBySearch(searchStr, table)
	}

	templates.TableRows(tabledata).Render(r.Context(), w)

	// data := util.GenerateTemplateData(types.User{}, users, "users")

	// executeTemplate("tablerows.gohtml", w, data, "templates/tablerows.gohtml")

}

func (s *Server) handleDataGetByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	table := s.store.GetTableType(vars["table"])
	id, _ := strconv.Atoi(vars["id"])

	tabledata := s.store.Get(id, table)

	fmt.Println(tabledata)

	templates.TableRows(tabledata).Render(r.Context(), w)

}

func (s *Server) handleDataEditByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	table := s.store.GetTableType(vars["table"])

	tabledata := s.store.Get(id, table)

	templates.TableRowEdit(tabledata.FirstRow()).Render(r.Context(), w)

}

func (s *Server) handleDataDeleteByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	// vars := mux.Vars(r)
	// id, _ := strconv.Atoi(vars["id"])

	// user := s.store.Delete(id)

	// json.NewEncoder(w).Encode(user)
}

func (s *Server) handleDataAppendByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	//id, _ := strconv.Atoi(vars["id"])
	table := s.store.GetTableType(vars["table"])

	decoder := schema.NewDecoder()

	appendUser := &types.User{}
	r.ParseForm()

	fmt.Println(r.PostForm)

	decoder.Decode(appendUser, r.PostForm)

	fmt.Println(appendUser)

	s.store.Append(appendUser, table)

	templates.TableRow(appendUser).Render(r.Context(), w)

}

func (s *Server) handleDataCreateByID(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	table := s.store.GetTableType(vars["table"])
	tabledata := s.store.Create(table)

	templates.TableRows(tabledata).Render(r.Context(), w)
}

func (s *Server) handleBase(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)

	// http.ServeFile(w, r, templ.Handler(base()))

	templates.Base().Render(r.Context(), w)

	// template := parseTemplate("index.html", "templates/index.html")

	// executeTemplate(template, w, nil)

}

func (s *Server) handleStylesheets(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)

	http.ServeFile(w, r, "stylesheets/"+vars["stylesheet"])

}

func (s *Server) handleScripts(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)

	http.ServeFile(w, r, "scripts/"+vars["script"])

}

func (s *Server) handleTemplates(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL)
	vars := mux.Vars(r)
	http.ServeFile(w, r, "templates/"+vars["template"])
}
