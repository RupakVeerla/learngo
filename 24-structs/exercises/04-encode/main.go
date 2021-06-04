// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

type game struct {
	ID, Price   int
	Name, Genre string
}

func main() {
	games := []game{
		{ID: 1, Name: "god of war", Price: 50, Genre: "action adventure"},
		{ID: 2, Name: "x-com 2", Price: 30, Genre: "strategy"},
		{ID: 3, Name: "minecraft", Price: 20, Genre: "sandbox"},
	}

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
> list - list all the games
> Id id - Game id 
> quit - to exit
> save - save as json

`)
		in.Scan()

		fmt.Println()

		input := strings.Fields(in.Text())
	loop:
		switch input[0] {
		case "list":
			fmt.Printf("#ID   %-15s %-18s Price\n", "Game", "Genre")

			for _, game := range games {
				fmt.Printf("#%d: %-15q %-20s $%d\n", game.ID, game.Name, "("+game.Genre+")", game.Price)
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
				if id == game.ID {
					fmt.Printf("  %-15s %-18s Price\n", "Game", "Genre")
					fmt.Printf("%-15q %-20s $%d\n", game.Name, "("+game.Genre+")", game.Price)
					break loop
				}
			}
			fmt.Println("No game found for ID: ", id)
		case "save":
			enc, err := json.MarshalIndent(games, "", "\t")
			if err != nil {
				fmt.Println("Error while saving: ", err)
			}
			fmt.Println(string(enc))
			return
		}
	}
}
