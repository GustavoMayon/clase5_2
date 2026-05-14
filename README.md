# Saludos en GO

## Instalación
ejecuta el siguiente comando para instalar el paquete
```bash
go get -u github.com/GustavoMayon/clase5_2
```
## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu código:
```Go
package main
import
"fmt"
"github.com/GustavoMayon/clase5_2"

func main() {
mensaje, err := clase5.Saludo("Gustavo")
	if err != nil {
		fmt.Println("Ocurrió un error: ",err)
    return
	}
	fmt.Println(mensaje)
}
```
Este ejemplo importa el paquete github.com/GustavoMayon/clase5 y llama a la función Saludo() saludo personalizado. Si ocurre un error, se imprime un mensaje de error.
