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
	
	json.NewEncoder(w).Encode(results)
	
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
	
	ret := database.UpdateTask(id, true)
	
	if ret {
		json.NewEncoder(w).Encode("OK")
	}else{
		json.NewEncoder(w).Encode("NOK")
	}
}

func UncheckTaskEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
		
	var id string = params["id"]
	
	ret := database.UpdateTask(id, false)
	
	if ret {
		json.NewEncoder(w).Encode("OK")
	}else{
		json.NewEncoder(w).Encode("NOK")
	}
}

func main(){
	router := mux.NewRouter()
	router.HandleFunc("/tasks", GetTasksEndpoint).Methods("GET")
	router.HandleFunc("/addtask", AddTasksEndpoint).Methods("POST")
	router.HandleFunc("/checktask/{id}", CheckTaskEndpoint).Methods("PUT")
	router.HandleFunc("/unchecktask/{id}", UncheckTaskEndpoint).Methods("PUT")
	log.Fatal(http.ListenAndServe(":12345", router))
}