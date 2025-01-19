package main

import (
	"bufio"
	"cmd_client/rest"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Starting up. Type 'quit' to exit.")

	rest.JoinLobby(scanner)
}
