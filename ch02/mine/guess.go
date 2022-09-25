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
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin)
	success := false
	max_asked := 100
	min_asked := 1
	for guesses := 0; guesses < 10; {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess >= max_asked {
			fmt.Println("As you already know, the number is low than", input)
			continue
		}
		if guess <= min_asked {
			fmt.Println("As you already know, the number is greater than", input)
			continue
		}
		guesses++
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
			if guess > min_asked {
				min_asked = guess
			}
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
			if guess < max_asked {
				max_asked = guess
			}
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
