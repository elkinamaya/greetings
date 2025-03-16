package greetings

import (
	"errors"
	"fmt"
	"math/rand/v2"

)

// hello devuelve un saludo

func Hello(name string)(string, error){
	if name == "" {
		return name, errors.New("Nombre vacio")
	}
	message := fmt.Sprintf(randomFormat(),name)
	return message , nil
}

func Hellos(names []string)(map[string]string,error){
	messages := make(map[string]string)
	for _, name := range names {
		message, err  := Hello(name)
		if err != nil{
			return nil , err 
		}

		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []  string{
		"!Hola, %v!, bienvenido a esta prueba", 
		"!Hi, %v!, welcome",
		"!Que bueno verte, %v!",
		"!Saludo, %v!",
	}

	return formats[rand.IntN(len(formats))]
}