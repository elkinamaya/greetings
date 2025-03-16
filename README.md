# Saludos en Go

Este paquete proporciona una forma de simple de obtener saludos personales de forma aleatoria

## Instalacion
Ejectua el siguiente comando para instalar el paquete:
``` bash
go get -u github.com/elkinamaya/greetings
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo:

```go
package main

import (
	"fmt"	
	"github.com/elkinamaya/greetings"
)

func main() {
	
	messages, err := greetings.Hello("Elkin")

	if err!= nil{
	    fmt.Println("Ha ocurrido un error:", err)
        return 
	}
	fmt.Println(messages)
}
```

Este ejemplo importa el paquete github.com/elkinamaya/greetings y llama la funcion Hello que retorna saludos aleatorios con el nombre