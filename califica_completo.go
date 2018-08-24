// aplicación que informa si se ha superado o no la materia
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Escribe una puntuación: ")
	reader := bufio.NewReader(os.Stdin)
	entrada, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}
	entrada = strings.TrimSpace(entrada)
	puntuacion, error := strconv.ParseFloat(entrada, 64)
	if error != nil {
		log.Fatal(error)
	}

	var status string
	if puntuacion >= 50 {
		status = "aprobado"
	} else {
		status = "suspenso"
	}
	fmt.Println("una puntuación de", puntuacion, "es", status)

}
