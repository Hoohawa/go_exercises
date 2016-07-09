package main

import (
	"fmt"
	"github.com/hoohawa/go_exercises/ch4/github"
	"log"
	"os"
	"time"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Less than a month ago\n")
	PrintFilteredResults(result, daysOld)
	fmt.Println()

	fmt.Printf("More than a month ago\n")
	PrintFilteredResults(result, monthsOld)
	fmt.Println()

	fmt.Printf("More than a year ago\n")
	PrintFilteredResults(result, yearsOld)
	fmt.Println()
}

func PrintFilteredResults(result *github.IssuesSearchResult, f func(*github.Issue) bool) {
	var filteredItems [](*github.Issue)
	for _, item := range result.Items {
		if f(item) {
			filteredItems = append(filteredItems, item)
		}
	}

	fmt.Printf("%d issues:\n", len(filteredItems))
	for _, item := range filteredItems {
		fmt.Printf("#%-5d %9.9s %.55s %v\n",
			item.Number, item.User.Login, item.Title, item.CreatedAt)
	}
}

func daysOld(issue *github.Issue) bool {
	return time.Since(issue.CreatedAt).Hours()/24 < 30
}

func monthsOld(issue *github.Issue) bool {
	return time.Since(issue.CreatedAt).Hours()/24 >= 30 &&
		time.Since(issue.CreatedAt).Hours()/24/364 < 1
}

func yearsOld(issue *github.Issue) bool {
	return time.Since(issue.CreatedAt).Hours()/24/364 >= 1
}
