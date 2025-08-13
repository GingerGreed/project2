package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	shoppingList := []string{}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n === MENU ===")
		fmt.Println("1 Add good")
		fmt.Println("2 Listing")
		fmt.Println("3 Delete good")
		fmt.Println("4 Exit")
		fmt.Println("Choose:")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Print("Enter good")
			item, _ := reader.ReadString('\n')
			item = strings.TrimSpace(item)
			shoppingList = append(shoppingList, item)
			fmt.Println("U add good")
		case "2":
			fmt.Println("List")
			for i, item := range shoppingList {
				fmt.Printf("%d. %s\n", i+1, item)
			}
		case "3":
			fmt.Print("Enter number of good to DELETE: ")
			var index int
			fmt.Scan(&index)
			if index > 0 && index <= len(shoppingList) {
				shoppingList = append(shoppingList[:index-1], shoppingList[index:]...)
				fmt.Println("Good deleted")
			} else {
				fmt.Println("Invalid index")
			}
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again")
		}
	}
}
