package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"quasar/Model"

	"quasar/Service/Location"
	"quasar/Service/Message"

	"quasar/View"
)

func topSecret(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var satellites map[string][]Model.Satellite
	err = json.Unmarshal(bodyBytes, &satellites)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	/*Asi se peude utilizar un set de un paquete . Tengo que declara primero una variable a un puntero de ese paquete*/
	var satellitesReq View.SatellitesRequest
	satellitesReq.SetSatellites(satellites)

	/*Aca voy a ir a los dos metodos propuestos en el pdf*/

	position, err := Location.GetLocation(satellites)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	message, err := Message.GetMessage(satellites)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	var messageAndLocationResponse View.MessagePositionResponse
	messageAndLocationResponse.Message = message
	messageAndLocationResponse.Position = position

	response := messageAndLocationResponse

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Write(jsonBytes)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome the my GO API of MeLi Challenge!")
}

func main() {
	/*NewLocationStore y NewSatelliteStore deben hacerce aqui */
	/*Asi creo un store*/

	http.HandleFunc("/", indexRoute)
	http.HandleFunc("/topsecret", topSecret)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
