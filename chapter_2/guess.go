// guess challenges players to guess a random number.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	success := false
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	for guesses := 1; guesses < 11; guesses++ {
		fmt.Println("You have", 11-guesses, "guess(es) left.")
		fmt.Print("Make a guess: ")
		guess, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		numericGuess, err := strconv.Atoi(strings.TrimSpace(guess))

		if err != nil {
			log.Fatal(err)
		}

		if numericGuess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else if numericGuess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			fmt.Println("Good job! You guessed it!")
			success = true
			break
		}
	}

	if !success {
		fmt.Println("Commiserations! The target was: ", target)
	}

}
