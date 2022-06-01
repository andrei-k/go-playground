package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

const (
	ROCK     = 0 // Beats scissors, (scissors + 1) % 3 = 0
	PAPER    = 1 // Beats rock, (rock + 1) % 3 = 1
	SCISSORS = 2 // Beats paper, (paper + 1) % 3 = 2
)

type Game struct {
	DisplayChan chan string
	RoundChan   chan int
	Round       Round
}

type Round struct {
	RoundNumber   int
	PlayerScore   int
	ComputerScore int
}

var reader = bufio.NewReader(os.Stdin)

func (g *Game) Rounds() {
	// Use select to process input in channels
	for {
		select {
		case round := <-g.RoundChan: // Receive from channel
			g.Round.RoundNumber = g.Round.RoundNumber + round
			// Send info back into the channel to indicate
			// that it's finished doing what it needs to do
			g.RoundChan <- 1
		case msg := <-g.DisplayChan: // Receive from channel
			fmt.Print(msg)
			// Send info back into the channel to indicate
			// that it's finished doing what it needs to do
			g.DisplayChan <- ""
		}
	}
}

func (g *Game) PrintIntro() {
	// Use channels and select statement to print the information to the screen
	g.DisplayChan <- `
Rock Paper Scissors
Best out of 3 wins!
-------------------

`
	// Wait for the channel to send something back
	// because we don't want to try using this
	// channel before it's finished printing
	<-g.DisplayChan
}

func (g *Game) PlayRound() bool {
	rand.Seed(time.Now().UnixNano())
	playerValue := -1

	g.DisplayChan <- fmt.Sprintf("Round %d\n", g.Round.RoundNumber)
	<-g.DisplayChan
	g.DisplayChan <- "Please enter rock, paper, or scissors -> "
	<-g.DisplayChan

	playerChoice, _ := reader.ReadString('\n')
	playerChoice = strings.Replace(playerChoice, "\n", "", -1)

	computerValue := rand.Intn(3)

	if playerChoice == "rock" {
		playerValue = ROCK
	} else if playerChoice == "paper" {
		playerValue = PAPER
	} else if playerChoice == "scissors" {
		playerValue = SCISSORS
	}

	g.DisplayChan <- fmt.Sprintf("\nPlayer chose: %s\n", strings.ToUpper(playerChoice))
	<-g.DisplayChan

	switch computerValue {
	case ROCK:
		g.DisplayChan <- "Computer chose: ROCK\n"
		<-g.DisplayChan
	case PAPER:
		g.DisplayChan <- "Computer chose: PAPER\n"
		<-g.DisplayChan
	case SCISSORS:
		g.DisplayChan <- "Computer chose: SCISSORS\n"
		<-g.DisplayChan
	default:
	}

	if playerValue == computerValue {
		g.DisplayChan <- "It's a draw!\n\n"
		<-g.DisplayChan
		return false
	} else if playerValue == -1 {
		g.DisplayChan <- "Invalid choice!\n\n"
		<-g.DisplayChan
	} else if playerValue == (computerValue+1)%3 {
		// We can use the modulo operator because there's a cyclical relationship
		// Three (Rock) beats two (Scissors)
		// Two (Scissors) beats one (Paper), and
		// One (Paper) beats three (Rock)
		g.PlayerWins()
	} else {
		g.ComputerWins()
	}

	return true
}

func (g *Game) ComputerWins() {
	g.Round.ComputerScore++
	g.DisplayChan <- "Computer wins!\n\n"
	<-g.DisplayChan
}

func (g *Game) PlayerWins() {
	g.Round.PlayerScore++
	g.DisplayChan <- "Player wins!\n\n"
	<-g.DisplayChan
}

func (g *Game) PrintSummary() {
	g.DisplayChan <- fmt.Sprintf(`
-------------------
Player score: %d
Computer score: %d

`, g.Round.PlayerScore, g.Round.ComputerScore)
	<-g.DisplayChan

	if g.Round.PlayerScore > g.Round.ComputerScore {
		g.DisplayChan <- "Player wins the game!\n\n"
		<-g.DisplayChan
	} else {
		g.DisplayChan <- "Computer wins the game!\n\n"
		<-g.DisplayChan
	}
}

// Clears the Terminal screen
func (g *Game) ClearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// Windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// Linux and Mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
