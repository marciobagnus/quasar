package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"quasar/Model"
	"quasar/Service/Location"
	"quasar/Service/Message"
	"quasar/Store"
	"strings"

	"quasar/View"
)

/*Creo este store como variable global para los topSecret_split*/
var satellitesStore *Store.SatellitesStore

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
	/*Manera de utilizar un setter de un paquete. Tengo que declarar primero una variable a un puntero de ese paquete*/
	var satellitesReq View.SatellitesRequest
	satellitesReq.SetSatellitesRequest(satellites)

	/*Aca voy a ir a los dos metodos propuestos en el pdf*/
	message, err := Message.GetMessage(satellites)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}

	position, err := Location.GetLocation(satellites)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
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

func topSecret_splitGet(w http.ResponseWriter, r *http.Request) {

	/*En este caso, los satelites van a ser los del store*/
	satellites := satellitesStore.GetSatellites()

	/*Manera de utilizar un setter de un paquete. Tengo que declarar primero una variable a un puntero de ese paquete*/
	var satellitesReq View.SatellitesRequest
	satellitesReq.SetSatellitesRequest(satellites)

	/*Voy a los dos metodos propuestos en el chaleng*/
	position, err := Location.GetLocation(satellites)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
	message, err := Message.GetMessage(satellites)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
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
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)

}
func topSecret_splitPost(w http.ResponseWriter, r *http.Request) {

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

	/*Recibo de a un satelite y extraigo los datos, por un lado el distanceMessage y por otro tengo q sacar el nombre del satelide de la URL*/
	var distanceMessage View.DistanceMessageRequest
	err = json.Unmarshal(bodyBytes, &distanceMessage)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	/*Obtengo el nombre del request que viene en el path para guardarlo en el satelite luego*/
	name := strings.TrimPrefix(r.URL.Path, "/topsecret_split/")

	satellite := Model.Satellite{
		Name:     name,
		Distance: distanceMessage.Distance,
		Message:  distanceMessage.Message,
	}

	response := satellitesStore.SaveSatellite(satellite, name)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

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

type splitHandler struct {
}

func newSplitHandler() *splitHandler {
	return &splitHandler{}
}

func (a splitHandler) topSecret_split(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		topSecret_splitGet(w, r)
		w.Write([]byte(""))
		return
	case "POST":
		topSecret_splitPost(w, r)
		w.Write([]byte(""))
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return
	}
}

func main() {

	port := os.Getenv("PORT")
	/*Creo un store de satelites para los topSecret_split*/
	satellitesStore = Store.NewSatellitesStore()
	/*Creo un handler para poder redireccionar los metodos split*/

	handlerSplit := newSplitHandler()
	http.HandleFunc("/topsecret_split/", handlerSplit.topSecret_split)

	http.HandleFunc("/", indexRoute)
	http.HandleFunc("/topsecret/", topSecret)

	/*err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}*/
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
