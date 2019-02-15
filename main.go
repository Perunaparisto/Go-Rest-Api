package main

// Importing necessary packages
import (
    // Used to convert data
    "encoding/json"
    // Used for routing requests to routes and router management
    "github.com/gorilla/mux"
    //Used for output to the server
    "log"
    // Provides client and server implementations
    "net/http"
    // Used for printing text to the terminal
    "fmt"
    // Used for converting strings to integers
    "strconv"
)

// The person Type (more like an object)
// Defines what a Person structure is made of
type Person struct {
    // The ID is used for identifying the person
    ID        int   `json:"id,omitempty"`
    // The firstname and lastname are strings
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`

    // New info slots can be added simply by duplicating an existing part of the structure
    //Phonenumber string   `json:"phonenumber,omitempty"`

    // The address is a seperate structure defined under the person
    Address   *Address `json:"address,omitempty"`
}
// Defines the Address structure located inside of the Person structure
type Address struct {
    // Just like the firstname and lastname the city and state are two strings
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    // The JSON encodes the people list
    json.NewEncoder(w).Encode(people)
}

// Display a single person info
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    // The for loop checks the people list for the given Id
    for _, item := range people {
        // The given ID is converted from a string to an integer
        tmpID, _ := strconv.Atoi(params["id"])
        // If the current list item's ID matches the given ID, the loop stops
        if item.ID == tmpID {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// create a new person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    // Setting variables for the creation process
    var person Person
    var tmpID int
    tmpID = 0

    _ = json.NewDecoder(r.Body).Decode(&person)

    // The for loop checks the person list
    for _, item := range people {
        // Using this process, the loop ensures that the next ID is one bigger than the last

        // If the current list item is equal to or bigger than the temporary ID, the temporary
        //ID will be set to the same value as the current list item's ID
        if item.ID >= tmpID {
            tmpID = item.ID
        }
    }
    // After the for loop has made the temporary ID as big as the biggest currently existing ID
    // The upcoming ID will be one bigger than that
    person.ID = tmpID + 1
    
    // A new item is added to the people list
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Deleting an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    // The for loop checks the person list
    for index, item := range people {
        // The given ID is converted to an integer
        tmpID, _ := strconv.Atoi(params["id"])
        // If the given ID is equal to the current list item's ID, the item will be deleted from the list
        if item.ID == tmpID {
            people = append(people[:index], people[index+1:]...)
            break
        }
        json.NewEncoder(w).Encode(people)
    }
}

//Edit an existing item
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var index int

    // The temporary ID is defined outside of the for loop because it will be used by two different loops    
    tmpID, _ := strconv.Atoi(params["id"])
    // The for loop checks the person list
    for _, item := range people {
        // If a matching ID is found, the index integer will be set to it's value
        if item.ID == tmpID {
            index = tmpID
            break
        }
    }

    // Replacing old info with any possible new info
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)

    // If the firstname slot isn't empty, the firstname will be set to the given string
    if person.Firstname != "" {
        people[index].Firstname = person.Firstname
        fmt.Printf("Changing firstname to: %#v\n", person.Firstname)
    }

    // If the lastname slot isn't empty, the lastname will be set to the given string
    if person.Lastname != "" {
        people[index].Lastname = person.Lastname
        fmt.Printf("Changing lastname to: %#v\n", person.Lastname)
    }

    // If the city slot isn't empty, the city will be set to the given string
    if person.Address.City != "" {
        people[index].Address.City = person.Address.City
        fmt.Printf("Changing city to: %#v\n", person.Address.City)
    }

    // If the state slot isn't empty, the state will be set to the given string
    if person.Address.State != "" {
        people[index].Address.State = person.Address.State
        fmt.Printf("Changing state to: %#v\n", person.Address.State)
    }
    
    json.NewEncoder(w).Encode(people)
}



// The main function is used to boot up everything
func main() {
    // Starting the router with mux
    fmt.Println("Starting...")
    router := mux.NewRouter()
    fmt.Println("Router started");
    // Creating three people for the list
    people = append(people, Person{ID: 1, Firstname: "Aria", Lastname: "Turner", Address: &Address{City: "Fayette", State: "Alabama"}})
    people = append(people, Person{ID: 2, Firstname: "John", Lastname: "Doe", Address: &Address{City: "Columbus", State: "Ohio"}})
    people = append(people, Person{ID: 3, Firstname: "Koko", Lastname: "Evans", Address: &Address{City: "Tampa", State: "Florida"}})
    people = append(people, Person{ID: 2, Firstname: "Joel", Lastname: "Ford", Address: &Address{City: "San Fransisco", State: "California"}})
    people = append(people, Person{ID: 3, Firstname: "Troy", Lastname: "Carson", Address: &Address{City: "Bellingham", State: "Washington"}})
    fmt.Println("Data defined")
    // Defining the end-points
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", UpdatePerson).Methods("PUT")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    fmt.Println("End-points defined")
    // Starting the server
    log.Fatal(http.ListenAndServe(":8080", router))
    fmt.Println("Server started")
}
