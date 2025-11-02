package main

import (
	"fmt"
)

func main() {
	todos := &Todos{}

	for {
		fmt.Println("\n=== TODO LIST ===")
		fmt.Println("1. Aggiungi task")
		fmt.Println("2. Elimina task")
		fmt.Println("3. Completa task")
		fmt.Println("4. Mostra tutti i task")
		fmt.Println("5. Esci")

		var scelta int
		fmt.Print("Scegli un'opzione: ")
		fmt.Scan(&scelta)

		switch scelta {
		case 1:
			var title string
			fmt.Print("Inserisci il titolo del task: ")
			fmt.Scan(&title)
			todos.Add(title)

		case 2:
			var index int
			fmt.Print("Inserisci l'indice del task da eliminare: ")
			fmt.Scan(&index)
			todos.Delete(index)

		case 3:
			var index int
			fmt.Print("Inserisci l'indice del task da completare: ")
			fmt.Scan(&index)
			todos.Complete(index)

		case 4:
			todos.List()

		case 5:
			fmt.Println("Uscita...")
			return

		default:
			fmt.Println("Opzione non valida.")

		}
	}
}
