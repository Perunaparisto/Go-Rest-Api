
//Basic preparations
package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "fmt"
    "strconv"
)

// Defining the objects
type Person struct {
    ID        int   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// Display a single person
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    //Checking the chart for the given ID
    for _, item := range people {
    	//Converting the given string to an integer
        i, _ := strconv.Atoi(params["id"])
        if item.ID == i {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// create a new person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    var i int
    //Checking the existing IDs
    _ = json.NewDecoder(r.Body).Decode(&person)  
    i = 0
    for _, item := range people {
        if item.ID >= i {
            i = item.ID
        }
    }
    //Making the ID one bigger than the last
    person.ID = i + 1
    
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Delete a person
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    //Checking the chart for the given ID
    for index, item := range people {
        i, _ := strconv.Atoi(params["id"])
        if item.ID == i {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

//Edit an existing person
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var index int
    //Finding the person to be edited
    i, _ := strconv.Atoi(params["id"])
    for _, item := range people {
        if item.ID == i {
            //jzz = item
            index = i
            break
        }
    }

    //Replacing the existing information
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)

    //Checking if something has changed and replacing old information
    if person.Firstname != "" {
        people[index].Firstname = person.Firstname

    if person.Lastname != "" {
        people[index].Lastname = person.Lastname
    }

    if person.Address.City != "" {
        people[index].Address.City = person.Address.City
    }

    if person.Address.State != "" {
        people[index].Address.State = person.Address.State
    }
    
    json.NewEncoder(w).Encode(people)
}



// main function to boot up everything
func main() {
    fmt.Println("Starting...");
    router := mux.NewRouter()
    fmt.Println("Router started");
    //Creating the base information
    people = append(people, Person{ID: 1, Firstname: "____", Lastname: "____", Address: &Address{City: "_____", State: "_____"}})
    people = append(people, Person{ID: 2, Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    people = append(people, Person{ID: 3, Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
    fmt.Println("Data defined");
    //Defining the commands for the program
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    fmt.Println("End-points defined");
    //Final preparations
    log.Fatal(http.ListenAndServe(":8080", router))
    fmt.Println("Server started");
}

//{"id":"___","firstname":"___","lastname":"___","address":{"city":"___","state":"___"}}