// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
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
> quit - to exit
`)
		in.Scan()

		switch in.Text() {
		case "list":
			fmt.Printf("#ID   %-15s %-18s Price\n", "Game", "Genre")

			for _, game := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.id, game.name, "("+game.genre+")", game.price)
			}
		case "quit":
			fmt.Println("Exiting!")
			return
		}
	}
}
