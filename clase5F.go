package clase5f

import (
	"errors"
	"fmt"
	"math/rand"
)

func Saludo(nombre string) (string, error) {
	//manejando error de string vacio
	if nombre == "" {
		return "", errors.New("Nombre vacío")
	}

	mensaje := fmt.Sprintf(SaludoAleaotrio(), nombre)
	return mensaje, nil
}

func SaludoVarios(nombres []string) (map[string]string, error) {
	mensajes := make(map[string]string)
	for _, nombre := range nombres {
		mensaje, err := Saludo(nombre)
		if err != nil {
			return nil, err
		}
		mensajes[nombre] = mensaje
	}
	return mensajes, nil
}

func SaludoAleaotrio() string {
	saludoRandom := []string{
		"¡Hola %v, te deseo lo mejor!",
		"¡Que bueno verte de nuevo, %v!",
		"¡%v, ánimo sigue así!",
	}
	return saludoRandom[rand.Intn(len(saludoRandom))]
}
