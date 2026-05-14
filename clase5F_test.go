package clase5f

import (
	"regexp"
	"testing"
)

func TestSaludoNombre(t *testing.T) {
	// Guardamos el nombre que vamos a probar.
	nombre := "juan"
	// Construimos una expresión regular que busque la palabra completa "juan".
	want := regexp.MustCompile(`\b` + nombre + `\b`)

	// Llamamos a la función Saludo con el nombre "juan".
	mensaje, err := Saludo("juan")
	// Si el mensaje no contiene el nombre esperado o hubo un error, el test falla.
	if !want.MatchString(mensaje) || err != nil {
		// Mostramos el valor obtenido, el error y lo que esperábamos encontrar.
		t.Fatalf(`Saludo("juan") = %q, %v, quiere coincidencia para%#q, nil`, mensaje, err, want)
	}
}

func TestSaludoVacio(t *testing.T) {
	mensaje, err := Saludo("")
	if mensaje != "" || err == nil {
		t.Fatalf(`Saludo("") = %q, %v, quiere "", error`, mensaje, err)
	}
}
