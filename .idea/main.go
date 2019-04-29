package main

import (
	"bytes"         // lo que yo le mando se lo mando como stream ojo
	"encoding/json" // necesito manejar json quiero usar la misma http para hacer un psot pegarle, necesito crear un mapa
	"fmt"
	"io/ioutil" // necesito para poder leer InputString
	"net/http"  // hago un http conexion recibe el reponse y un error
)

/*
api
Vamos a conectarnos, necesitamos http
es obligatorio manejar los errores
*/
func main() {
	//tipos personales
	fmt.Println("Iniciando aplicacion")
	response, err := http.Get("https://httpbin.org/ip") // devuelve dos valores lo vamos a capturar en dos variables
	if err != nil {
		fmt.Println("Ocurrio un error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body) // me devuelve dos elementos, tomo lo que necesito _ es el otro valor,solo muestra contenido no puedo hacer nada
		fmt.Println(string(data))
	}

	jsonData := map[string]string{ // el tipo de key y afuera del corchete el valor
		"nombre":   "pipo",
		"apellido": "pepo"}

	// al mapa lo conviero en json, todos los metodos me deberian devolver un variable error
	jsonValue, _ := json.Marshal(jsonData) // convierto ese mapa en un json si me da un error de parseo lo tengo que capturar
	// como ya la use , ya existe no uso := sino =, cuidado cuando haces un psot
	response, err = http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Ocurrio un error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body) // me devuelve dos elementos, tomo lo que necesito _ es el otro valor,solo muestra contenido no puedo hacer nada
		fmt.Println(string(data))
	}

	fmt.Println("Terminando aplicacion")
}
