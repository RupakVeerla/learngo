// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

type game struct {
	ID, Price   int
	Name, Genre string
}

func main() {
	var games []game

	err := json.Unmarshal([]byte(data), &games)
	if err != nil {
		fmt.Println("Error decoding data: ", err)
		return
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

