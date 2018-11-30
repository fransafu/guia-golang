# Variables

No solo terminamos revisando las variables, tambien como hacer comenarios, algo tan fundamental como poder dejar una explicacion, guia o referencia de una funcion o variable que sea de muchisima utilidad.

Volviendo a las variables, hay 2 formar de declarar:

* var nombreVar tipo = valor
* nombreVar := valor
* var nombreVar = valor

La primera de una forma mas explicita pero la segunda es igual de aceptada y se suele utilizar mucho. Un punto importante de esto es que la segunda forma tiene una inferencia de tipo sobre la variable pero esto no significa que se pueda cambiar el valor de la variable de tipo string a integer, esto es porque Go es un lenguaje fuertemente tipado. La tercera es una simplificacion de la primera, se omite el tipo de dato

Otro punto a ver es la forma de importar los paquetes o librerias externas, pueden notar que es con parentesis y se pueden agregar hacia abajo siempre y cuando sean utilizados. Una caracteristica de Golang es que remueve los paquetes que no son utilizados (no te asustes si desaparecen lineas de codigo al momento de compilar).

