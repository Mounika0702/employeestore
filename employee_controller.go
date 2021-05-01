package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var employeeDB = make(map[string]Employee)

func getDemployeeDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if _, ok := employeeDB[id]; ok {
		fmt.Println("checked the ID, it is present")
		w.Header().Set("Content-Type", "application/json")
		j, err := json.Marshal(Details{id, employeeDB[id].Name, employeeDB[id].Address})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	} else {
		fmt.Println("Not present")
		http.Error(w, "Data not found", http.StatusNotFound)
		return
	}

}

func updateEmployeeDetails(w http.ResponseWriter, r *http.Request) {

	var employee Details
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("gonna check the ID")
	if _, ok := employeeDB[employee.ID]; ok {
		fmt.Println("checked the ID")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		fmt.Println("I am going")
		employeeDB[employee.ID] = Employee{
			Name:    employee.Name,
			Address: employee.Address,
		}
		fmt.Println("Added the employee details")
		fmt.Println(employeeDB)
		w.WriteHeader(http.StatusCreated)
	}
}
