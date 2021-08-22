package Message

import (
	"errors"
	"quasar/Model"
)

func GetMessage(mapSatellites map[string][]Model.Satellite) (message string, err error) {

	var satellites []Model.Satellite
	for _, v := range mapSatellites {
		satellites = v
	}

	arrayMessages := getArrayMessages(satellites)

	//si me viene vacio
	if isEmptyArray(arrayMessages) {
		return message, errors.New("Array of messages is empty")
	}
	messageSize := getMessageSize(arrayMessages)
	/*Ahora comparo el tamaño del mensaje con el del array*/
	if isIncompleteArray(arrayMessages, messageSize) {
		return message, errors.New("One of the array of message is incomplete")
	}
	messagesWithouDelay := removeDelay(arrayMessages, messageSize)
	messageBuild := buildMessage(messagesWithouDelay, messageSize)
	message = messageBuild

	return message, nil

}

func getArrayMessages(satellites []Model.Satellite) [][]string {
	var messages [][]string

	for _, s := range satellites {
		messages = append(messages, s.Message)
	}
	return messages
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

			//No itelo la palabras vacias
			if word != "" {
				//La primer palabra del mensaje
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

func isEmptyArray(messages [][]string) bool {
	/*rocorro todas las palabras del string*/
	for _, message := range messages {
		for _, word := range message {
			if word != "" {
				return false
			}
		}

	}
	return true
}

func isIncompleteArray(messages [][]string, msgSize int) bool {
	for _, message := range messages {
		if len(message) < msgSize {
			return false
		}
	}
	return true
}
