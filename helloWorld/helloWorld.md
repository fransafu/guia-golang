# Hello World

Hola, este es el inicio de mi guia personal de Golang.

Primero que todo quiero dejar claro que anotare lo que recuerde o encuentre importante de documentar, el resto asumo que pondan ganas de investigar o probar.

## Programa

Un programa en Go primero que todo debe terminar con la extension ".go". Ok, demasiado simple? Hay que anotar todo lo que se puede olvidar.

La estructura principal es la siguiente:

```
package main

import "lib"

func main() {
  // program
}
```

Como pueden ver siempre parte con la declaracion de un paquete, en Go los paquetes agrupan funcionalidades y tambien definen alcance.

Siempre debe existir el paquete main y es una buena practica llamar al paquete de la forma en que se llama la carpeta, por ejemplo. Si tiene un modulo llamado "basededatos" debes crear una carpeta llamada "basededatos" y los programas dentro deben decir "package basededatos".

Respecto al ejemplo que tengo en este sitio, corresponde un "Hola Mundo" y sigue la estructura basica de Go sin embargo el paquete se llama main para poder ejecutar este simple programa. Más adelante se vera cual es la estructura de un proyecto (en realidad las recomendaciones).