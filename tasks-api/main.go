package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
}

var tasks []Task
var currentID int

type App struct {
	Router *mux.Router
}

func (app *App) handleRoutes() {
	app.Router.HandleFunc("/tasks", app.getTasks).Methods("GET")
	app.Router.HandleFunc("/task/{id}", app.readTask).Methods("GET")
	app.Router.HandleFunc("/task", app.createTask).Methods("POST")
	app.Router.HandleFunc("/task/{id}", app.updateTask).Methods("PUT")
	app.Router.HandleFunc("/task/{id}", app.deleteTask).Methods("DELETE")
}

func (app *App) Initialise(initialTasks []Task, id int) {
	tasks = initialTasks
	currentID = id
	app.Router = mux.NewRouter().StrictSlash(true)
	app.handleRoutes()
}
func main() {
	app := App{}
	tasks, id := CreateInitialTasks()
	app.Initialise(tasks, id)
	app.Run("localhost:10000")
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	//
	response, err := json.Marshal(payload)
	if err != nil {
		sendError(w, http.StatusInternalServerError, fmt.Sprintf("cant marshal payload, err: %v", err.Error()))
		return
	}
	//log.Println(string(response))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if _, err := w.Write(response); err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func sendError(w http.ResponseWriter, statusCode int, err string) {

	errorMsg := map[string]string{"error": err}
	sendResponse(w, statusCode, errorMsg)

}

func (app *App) getTasks(writer http.ResponseWriter, request *http.Request) {
	//
	sendResponse(writer, http.StatusOK, tasks)

}

func (app *App) createTask(writer http.ResponseWriter, r *http.Request) {
	//
	//var reqBody map[string]interface{}
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	_ = t.createTask() // just sets currentID + 1 to t.ID and appends t to tasks slice
	sendResponse(writer, http.StatusCreated, t)

}

func (app *App) readTask(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request) // mux.Vars(r) returns route vars for current request, if any
	key := vars["id"]         // var is a map[string]string, values are always string, localhost:10000/tasks/2, type of 2 would be string in vars["id"]
	//log.Printf("%v, %T", key, key)
	id, err := strconv.Atoi(key)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "invalid task ID")
		return
	}

	//if id > currentID {
	//	sendError(writer, http.StatusNotFound, "task not found")
	//	return
	//} redundant,coverd by model.go/getTask, context: id can be even less than tasks[0].id i.e. 0, or -1 or 45,
	// simple logic: iterate through tasks and look for id, if not found return "task not found"

	var t Task
	t.ID = id
	if err := t.getTask(); err != nil {
		sendError(writer, http.StatusNotFound, err.Error())
		return
	}

	sendResponse(writer, http.StatusOK, t)

}

func (app *App) updateTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	id, err := strconv.Atoi(key)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "invalid task ID")
		return
	}
	var t Task
	t.ID = id

	err = json.NewDecoder(request.Body).Decode(&t)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "Invalid request payload")
		return
	}
	err = t.updateTask()
	if err != nil {
		sendError(writer, http.StatusNotFound, err.Error())
		return
	}
	sendResponse(writer, http.StatusOK, t)
}

func (app *App) deleteTask(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	id, err := strconv.Atoi(key)
	if err != nil {
		sendError(writer, http.StatusBadRequest, "invalid task ID")
		return
	}
	var t Task
	t.ID = id
	err = t.deleteTask()
	if err != nil {
		sendError(writer, http.StatusNotFound, err.Error())
		return
	}
	var successResponse = map[string]interface{}{"result": "successful deletion"}
	sendResponse(writer, http.StatusOK, successResponse)

}
