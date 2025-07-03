package api

import (
	"net/http"
	"github.com/AdityaAnandCodes/GoStats/collector"
	"encoding/json"
	"fmt"
)


func getStats(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	stats, err := collector.GetSystemStats()
	if err != nil {
		fmt.Println("Error Fetching Data for API: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}



func RunServer(){
	http.HandleFunc("/stats", getStats)
	fmt.Println("Server Running at - http://localhost:8080")
	err := http.ListenAndServe(":8080",nil)	
	if err != nil {
		fmt.Println("Error Starting Server :", err)
		return
	}
}