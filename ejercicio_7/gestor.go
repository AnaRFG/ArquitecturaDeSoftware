package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	Name  string `json: "name"`  //van en mayúscula para poder exportar el archivo
	Email string `json: "email"` //en formato json los datos van en minus
	Phone string `json: "phone"`
}

func saveContactsFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	//Formatear a JSON
	encoder := json.NewEncoder(file)
	err = encoder.Encode(contacts)

	if err != nil {
		return err
	}

	return nil
}

func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")

	if err != nil {
		return err
	}

	defer file.Close()

	//leer los datos en formato json
	decoder := json.NewDecoder(file)
	err = decoder.Decode(contacts)
	if err != nil {
		return err
	}

	return nil
}
func main() {
	var contacts []Contact //slice de contactos

	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos", err)
	}

	for {
		fmt.Print(
			"==== GESTOR DE CONTACTOS ==== \n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos \n",
			"3. Salir\n",
			"Elige una opción:",
		)
		var opcion int
		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Error al leer la opción:", err)
			return
		}

		switch opcion {
		case 1:
			//agregar contacto
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Nombre:")
			nombre, _ := reader.ReadString('\n')
			fmt.Print("Email:")
			email, _ := reader.ReadString('\n')
			fmt.Print("Telefono:")
			telefono, _ := reader.ReadString('\n')

			nombre = strings.TrimSpace(nombre)
			email = strings.TrimSpace(email)
			telefono = strings.TrimSpace(telefono)

			con := Contact{Name: nombre, Email: email, Phone: telefono}
			contacts = append(contacts, con) //agregar contacto al slice
			err := saveContactsFile(contacts)
			if err != nil {
				fmt.Println("Error al guardar el contacto")
			}
		case 2:
			//mostrar contactos
			fmt.Println("================")
			for index, conta := range contacts {
				fmt.Printf("%d. Nombre: %s  - Email: %s - Telefono: %s \n",
					index+1, conta.Name, conta.Email, conta.Phone)
			}
			fmt.Println("================")
		case 3:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Print("Opción incorrecta.")
		}
	}
}
