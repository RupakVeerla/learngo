// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------

type item struct {
	id, price int
	name      string
}

type game struct {
	item
	genre string
}

func main() {
	games := []game{
		{item: item{id: 1, name: "god of war", price: 50}, genre: "action adventure"},
		{item: item{id: 2, name: "x-com 2", price: 30}, genre: "strategy"},
		{item: item{id: 3, name: "minecraft", price: 20}, genre: "sandbox"},
	}

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> list - list all the games
> Id id - Game id 
> quit - to exit

`)
		in.Scan()

		fmt.Println()

		input := strings.Fields(in.Text())
	loop:
		switch input[0] {
		case "list":
			fmt.Printf("#ID   %-15s %-18s Price\n", "Game", "Genre")

			for _, game := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
			}
		case "quit":
			fmt.Println("Exiting!")
			return
		case "id":
			if len(input) != 2 {
				fmt.Println("Wrong input")
				break
			}

			id, err := strconv.Atoi(input[1])
			if err != nil {
				fmt.Println("Wrong input: ", err)
				break
			}

			for _, game := range games {
				if id == game.id {
					fmt.Printf("  %-15s %-18s Price\n", "Game", "Genre")
					fmt.Printf("%-15q %-20s $%d\n", game.name, "("+game.genre+")", game.price)
					break loop
				}
			}
			fmt.Println("No game found for ID: ", id)
		}
	}
}

