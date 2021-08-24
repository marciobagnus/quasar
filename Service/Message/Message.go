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
		return message, errors.New("array of messages is empty")
	}
	messageSize := getMessageSize(arrayMessages)
	/*Ahora comparo el tamaño del mensaje con el del array*/
	if isIncompleteArray(arrayMessages, messageSize) {
		return message, errors.New("one of the array of message is incomplete")
	}
	messagesWithouDelay := removeDelay(arrayMessages, messageSize)
	messageBuild := buildMessage(messagesWithouDelay, messageSize)
	message = messageBuild

	return message, nil

}

//De los satelites extraigo los mensajes
func getArrayMessages(satellites []Model.Satellite) [][]string {
	var messages [][]string

	for _, s := range satellites {
		messages = append(messages, s.Message)
	}

	return messages
}

/*Asume que el delay esta al inicio del mensaje. Recibe el tamaño del mensaje y
los mensajes. Luego
*/
func removeDelay(messages [][]string, messageSize int) [][]string {
	var messagesWithoutDelay [][]string

	for _, message := range messages {
		fromIndex := len(message) - messageSize
		messageWithoutDelay := message[fromIndex:]
		messagesWithoutDelay = append(messagesWithoutDelay, messageWithoutDelay)
	}
	return messagesWithoutDelay
}

//Recibe los mensajes sin el deelay
func buildMessage(messages [][]string, msgSize int) string {

	var message string
	//Voy a recorre el array de columna en columna. i son las filas y j las columnas
	for j := 0; j < msgSize; j++ {
		for i := 0; i < len(messages); i++ {

			word := messages[i][j]

			//No itero las palabras vacias
			if word != "" {
				//La primer palabra del mensaje
				if message != "" {
					message = message + " " + word
				} else {
					message = word
				}
				//Corto el for despues de la primer palabra que proceso
				break
			}
		}
	}
	return message
}

/*Calcula el tamaño del mensaje de acuerdo a la cantida de palabras distintas que hay en todo el array*/
func getMessageSize(messages [][]string) int {

	var uniqueWords []string

	for _, message := range messages {
		for _, word := range message {
			//No verifica las palabras vacias
			if word != "" {
				//Si la palabra bo esta en mi diccionario la agrego
				if !stringInSlice(word, uniqueWords) {
					uniqueWords = append(uniqueWords, word)
				}
			}
		}
	}
	//Devuelvo el tamaño de ese diccionario de palabras
	return len(uniqueWords)
}

//Recibe el string que quiero saber si esta en la lista y la lista
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//Valida si el array de mensajes esta vacío
func isEmptyArray(messages [][]string) bool {
	/*Rocorro todas las palabras del string*/
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
			return true
		}
	}
	return false
}
