package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = "8181"

// Hotel represents an hotel
type Hotel struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Timezone    string         `json:"timezone"`
	Address     *HotelAddress  `json:"address"`
	Location    *HotelLocation `json:"location"`
}

// HotelAddress represents the address of an hotel
type HotelAddress struct {
	LineOne string `json:"lineOne"`
	LineTwo string `json:"lineTwo"`
	ZipCode string `json:"zipCode"`
	City    string `json:"city"`
	Country string `json:"country"`
}

// HotelLocation represents the location of an hotel
type HotelLocation struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// ContractAddress represents the address of a smart contract
type ContractAddress struct {
	Address string `json:"address"`
}

// RetrieveHotelResponse represents the answer to the method "RetrieveHotel"
type RetrieveHotelResponse struct {
	Hotel

	HotelOwner string   `json:"owner"`
	UnitTypes  []string `json:"unitTypes"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/hotels/{hotelAddress}", getHotelsHandler).Methods("GET")
	r.HandleFunc("/hotels", createHotelHandler).Methods("POST")
	r.HandleFunc("/hotels/{hotelAddress}", modifyHotelHandler).Methods("PATCH")
	r.HandleFunc("/hotels/{hotelAddress}", deleteHotelHandler).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":"+port, r))

}

func getHotelsHandler(w http.ResponseWriter, r *http.Request) {
	// Get the hotel address

	// Retrieve the hotel info using WT API

	fakeHotel := Hotel{
		Name:        "FakeHotel",
		Description: "This is a fake hotel",
		Address:     &HotelAddress{},
		Location:    &HotelLocation{},
	}

	// For now let send back a fake hotel
	if err := json.NewEncoder(w).Encode(fakeHotel); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back the correct status code
	// 200: Success
	// 404: Hotel Contract not found

}

func createHotelHandler(w http.ResponseWriter, r *http.Request) {

	var h Hotel

	if err := json.NewDecoder(r.Body).Decode(&h); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Call the WT API to create the hotel

	// Return hotel address - fake for now
	address := ContractAddress{Address: "0x46506d900559ad005feb4645dcbb2dbbf65e19cc"}
	if err := json.NewEncoder(w).Encode(address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func modifyHotelHandler(w http.ResponseWriter, r *http.Request) {
	// Get the list of hotel properties needing to be changed

	// Call the WT API to do the changes

	// Return hotel address - fake for now
	address := ContractAddress{Address: "0x46506d900559ad005feb4645dcbb2dbbf65e19cc"}
	if err := json.NewEncoder(w).Encode(address); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func deleteHotelHandler(w http.ResponseWriter, r *http.Request) {
	// Get the hotel address

	// Delete the hotel

	// Send back the correct status code

	// 200: Success
	// 401: Not Authorized to delete the Hotel, operation restricted to the owner.
	// 404: Hotel Contract not found
	return
}
