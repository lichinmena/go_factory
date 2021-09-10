package main

import (
	"example/factory"
	"fmt"
	"os"
)

func main() {
	var t int
	fmt.Print("Digite la conexión deseada:")
	_, err := fmt.Scan(&t)
	if err != nil {
		fmt.Printf("Error al leer la opcion: %v", err)
		os.Exit(1)
	}

	conn := factory.Factory(t)
	if conn == nil {
		fmt.Println("Motor no valido")
		os.Exit(1)
	}

	err = conn.Connect()
	if err != nil {
		fmt.Printf("No se pudo crear la conexion %v", err)
		os.Exit(1)
	}

	now, err := conn.GetNow()
	if err != nil {
		fmt.Printf("No se pudo consultar le fecha: %v", err)
		os.Exit(1)
	}

	fmt.Println(now.Format("2006-01-02"))

	err = conn.Close()
	if err != nil {
		fmt.Printf("No se pudo cerrar la conexion")
	}
}
