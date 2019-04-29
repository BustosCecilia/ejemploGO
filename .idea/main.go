package main

import (
	"./operaciones"
	"fmt"
)
/*
// los metodos que son exportados desde afuera son de mayuscula public, si lo escribo con minusculo es private,
// las sentensias se separan por enter, los for con ; y el if si queres separar una de otra
// para declarar una variable , setea el cero del tipo de la variable
asignacion en tiempo de compilacion y cuidado en tiempo de ejecucion puede ocurrir valor de tipo,
go hace reasignacion de tipo dinamica una vez que le asigne no lo puedo cambiar
go es fuertemente tipado
*/
func main() {
	//tipos personales
fmt.Println(operaciones.Sumar(10,30))
fmt.Println(operaciones.Potencia(20,30))
}



