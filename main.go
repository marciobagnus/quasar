package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getMessage(w http.ResponseWriter, r *http.Request) {

	var arrayMessages = getArrayMessages(satellites)
	var messages [][]string

	messages = append(messages, arrayMessages...)

	messageSize := getMessageSize(messages[:])

	messagesWithoutDelay := removeDelay(messages[:], messageSize)

	message := buildMessage(messagesWithoutDelay, messageSize)

	json.NewEncoder(w).Encode(message)

}

func getLocation(w http.ResponseWriter, r *http.Request) {

	var arrayDistances = getArrayDistances(satellites)

	valorX, valorY := GetValues(arrayDistances)

	x := fmt.Sprintf("%f", valorX)
	y := fmt.Sprintf("%f", valorY)

	json.NewEncoder(w).Encode("X: " + x + " " + "Y: " + y)

}

func GetValues(distances []float64) (x, y float64) {

	var x1 float64 = -500
	var y1 float64 = -200
	var r1 float64 = distances[0]

	var x2 float64 = 100
	var y2 float64 = -100
	var r2 float64 = distances[1]

	var x3 float64 = 500
	var y3 float64 = 100
	var r3 float64 = distances[2]

	A := 2*x2 - 2*x1
	B := 2*y2 - 2*y1
	C := r1*r1 - r2*r2 - x1*x1 + x2*x2 - y1*y1 + y2*y2
	D := 2*x3 - 2*x2
	E := 2*y3 - 2*y2
	F := r2*r2 - r3*r3 - x2*x2 + x3*x3 - y2*y2 + y3*y3
	x = (C*E - F*B) / (E*A - B*D)
	y = (C*D - A*F) / (B*D - A*E)

	return x, y
}

// Types
type Satellite struct {
	name     string   `json:"name"`
	distance float64  `json:"distance"`
	message  []string `json:"message"`
}

type allSatellites []Satellite

// Persistence
var satellites = allSatellites{
	{
		name:     "kenobi",
		distance: -1595.89,
		message:  []string{"este", "", "", "mensaje", ""},
	},
	{
		name:     "skywalker",
		distance: -1575.89,
		message:  []string{"", "es", "", "", "secreto"},
	},
	{
		name:     "kenobi",
		distance: -1545.89,
		message:  []string{"este", "", "un", "", ""},
	},
}

func getArrayMessages(satellites []Satellite) [][]string {
	var messages [][]string

	for _, s := range satellites {
		messages = append(messages, s.message)
	}
	return messages
}

func getArrayDistances(satellites []Satellite) []float64 {
	var distances []float64

	for _, s := range satellites {
		distances = append(distances, s.distance)
	}
	return distances
}

func removeDelay(messages [][]string, messageSize int) [][]string {
	//asumimos que el delay esta al inicio del mensaje
	var messagesWithoutDelay [][]string

	for _, message := range messages {

		fromIndex := len(message) - messageSize
		messageWithoutDelay := message[fromIndex:]

		messagesWithoutDelay = append(messagesWithoutDelay, messageWithoutDelay)
	}
	return messagesWithoutDelay
}

//recibo los mensajes sin el deelay
func buildMessage(messages [][]string, msgSize int) string {

	var message string

	//los i son las filas y los j las columnas
	for j := 0; j < msgSize; j++ {
		for i := 0; i < len(messages); i++ {

			word := messages[i][j]

			if word != "" {
				if message != "" {
					message = message + " " + word
				} else {
					message = word
				}
				//corto despues de la primer palabra que proceso
				break
			}

		}
	}
	return message

}

func getMessageSize(messages [][]string) int {

	var uniqueWords []string

	for _, message := range messages {
		for _, word := range message {
			if word != "" {
				if !stringInSlice(word, uniqueWords) {
					uniqueWords = append(uniqueWords, word)
				}
			}
		}
	}

	return len(uniqueWords)

}

//recibe el string que quiero saber si esta en la lista y la lista
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)

	router.HandleFunc("/message", getMessage).Methods("GET")

	router.HandleFunc("/location", getLocation).Methods("GET")

	log.Fatal(http.ListenAndServe(":3001", router))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wecome the my GO API of MeLi Challenge!")
}
