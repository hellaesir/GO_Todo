package main
 
import (
	"github.com/hellaesir/entities"
	"github.com/hellaesir/database"
	
    "encoding/json"
    "log"
    "net/http"
 
    "github.com/gorilla/mux"
)

func GetTasksEndpoint(w http.ResponseWriter, req *http.Request){

	var results []entities.Task = database.GetAllTasks()
	
	for _,item := range results{
		json.NewEncoder(w).Encode(item)
	}
	
	return
}

func AddTasksEndpoint(w http.ResponseWriter, req *http.Request){
	var task entities.Task
	_ = json.NewDecoder(req.Body).Decode(&task)
	
	database.AddTask(task)
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasksEndpoint).Methods("GET")
	router.HandleFunc("/addtask", AddTasksEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}