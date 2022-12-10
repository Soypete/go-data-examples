package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
)

func main() {
	// open file with list of urls
	f, err := os.Open("data/player/player_stats.csv")
	if err != nil {
		log.Fatalf("unable to access csv: %s", err)
		return
	}

	// build dataframe from csv
	df := dataframe.ReadCSV(f)
	// check for 0 row and 0 columns
	if df.Nrow() == 0 || df.Ncol() == 0 {
		log.Fatal("your csv is empty")
		return
	}

	// print out summaru of data
	fmt.Printf("Summary of the Player Stats dataframe\n%v+", df.Describe())

	// from player stats group the players by country
	dfTeams := df.GroupBy("team")
	teamsByGroup := dfTeams.GetGroups()

	// print @soypete top 5 teams USMNT, Uruguay, Argentina, Brazil, Mexico
	brazil := teamsByGroup["Brazil"].Filter(
		dataframe.F{
			Colname:    "player",
			Comparator: series.Neq,
			Comparando: "",
		},
	)
	fmt.Println("Brazil National Team")
	fmt.Println(brazil.Select([]string{"player", "age", "position"}))

	argentina := teamsByGroup["Argentina"].Filter(
		dataframe.F{
			Colname:    "player",
			Comparator: series.Neq,
			Comparando: "",
		},
	)
	fmt.Println("Argentina National Team")
	fmt.Println(argentina.Select([]string{"player", "age", "position"}))

	usmnt := teamsByGroup["United States"].Filter(
		dataframe.F{
			Colname:    "player",
			Comparator: series.Neq,
			Comparando: "",
		},
	)
	fmt.Println("US Mens National Team")
	fmt.Println(usmnt.Select([]string{"player", "age", "position"}))

	uruguay := teamsByGroup["Uruguay"].Filter(
		dataframe.F{
			Colname:    "player",
			Comparator: series.Neq,
			Comparando: "",
		},
	)
	fmt.Println("Uruguay National Team")
	fmt.Println(uruguay.Select([]string{"player", "age", "position"}))

	mexico := teamsByGroup["Mexico"].Filter(
		dataframe.F{
			Colname:    "player",
			Comparator: series.Neq,
			Comparando: "",
		},
	)
	fmt.Println("Mexico National Team")
	fmt.Println(mexico.Select([]string{"player", "age", "position"}))

	// players with top 5 scoring percentage
	scorers := df.Arrange(dataframe.RevSort("goals_per90"))
	fmt.Println(scorers.Select([]string{"player", "team", "goals", "goals_per90"}))

}
