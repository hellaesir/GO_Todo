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
	
	json.NewEncoder(w).Encode("OK")
}

func CheckTaskEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
		
	var id string = params["id"]
	
	task := database.GetTask(id)
	
	task.Checked = true
	
	database.AddTask(task)
	
	json.NewEncoder(w).Encode("OK")
}

func UncheckTaskEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
		
	var id string = params["id"]
	
	task := database.GetTask(id)
	
	task.Checked = false
	
	database.AddTask(task)
	
	json.NewEncoder(w).Encode("OK")
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasksEndpoint).Methods("GET")
	router.HandleFunc("/addtask", AddTasksEndpoint).Methods("POST")
	router.HandleFunc("/checktask/{id}", CheckTaskEndpoint).Methods("POST")
	router.HandleFunc("/unchecktask/{id}", UncheckTaskEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":12345", router))
}