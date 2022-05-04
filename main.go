package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Shipment model (Struct)
type Shipment struct {
	ID        int    `json:"id"`
	Manager   string `json:"manager"`
	Warehouse string `json:"warehouse"`
	Shipment  string `json:"shipment"`
	Product   string `json:"product"`
	Status    string `json:"status"`
	Approver  string `json:"approver"`
	Quantity  string `json:"quantity"`
	CreatedOn string `json:"created_on"`
}

// Initialize Shipment variable as slice (which nothing but an Array data structure)
var shipment []Shipment

// function to get all the shipments
func getShipments(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(shipment)
}

// function to get a single shipment
func getSingleShipment(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	for _, item := range shipment {
		if strconv.Itoa(item.ID) == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
	json.NewEncoder(res).Encode("Not found")

}

func main() {

	content, err := ioutil.ReadFile("./fakeShipmentData.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Now let's unmarshall the data into `shipment`

	err = json.Unmarshal(content, &shipment)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Println(string(shipment[0].ID))

	//Initialize router
	router := mux.NewRouter()

	// Handel the routes
	router.HandleFunc("/api/shipments", getShipments).Methods("GET")
	router.HandleFunc("/api/shipments/{id}", getSingleShipment).Methods("GET")

	// Initialize a server
	fmt.Printf("Starting server at port 8801\n")
	log.Fatal(http.ListenAndServe(":8801", router))

}
