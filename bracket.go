package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/codegangsta/cli"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	app := cli.NewApp()
	app.Action = func(c *cli.Context) {
		println("Calculating brackets")
		regions := []string{"Midwest", "East", "West", "South"}
		for _, region := range regions {
			fmt.Println(region + ":")
			teams := createTeams(16)
			for {
				length := len(teams)
				if length == 1 {
					fmt.Printf("Winner %s: %d\n\n", region, teams[0])
					break
				}

				winners := make([]int, 0)
				for index := 0; index < length/2; index++ {
					team1 := teams[index]
					team2 := teams[length-index-1]
					winner := playGame(team1, team2)
					winners = append(winners, winner)
					fmt.Printf("%d vs %d => %d\n", team1, team2, winner)
				}
				teams = winners
			}
		}
	}

	app.Run(os.Args)
}

func createTeams(n int) []int {
	teams := make([]int, n)
	for i := 0; i < n; i++ {
		teams[i] = i + 1
	}
	return teams
}

func playGame(team1 int, team2 int) int {
	// Make sure we have ascending seeds
	if team1 > team2 {
		temp := team1
		team1 = team2
		team2 = temp
	}

	diff := team2 - team1

	// Go with upsets 13% of the time, unless it's a 1 seed vs 11 or lower
	if !(team1 == 1 && diff > 10) && rand.Intn(100) < 13 {
		return team2
	}

	// 12 beats 5 23% of the time
	if team1 == 5 && team2 == 12 && rand.Intn(100) < 23 {
		return team2
	}

	// Higher seed wins 63% of the time
	if rand.Intn(100) < 63 {
		return team1
	}
	return team2
}
