package main

import (
	"fmt"
	"os"

	poker "github.com/dre4success/players"
)

func main() {
	fmt.Println("let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store := poker.NewInMemoryPlayerStore()
	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.StdOutAlerter), store)
	cli := poker.NewCli(os.Stdin, os.Stdout, game)
	cli.PlayPoker()
}
