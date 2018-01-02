// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/navossoc/gopl.io/ch4/github"
)

//!+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	var issuesM, issuesY, issuesO []*issue

	for _, v := range result.Items {
		item := issue(*v)

		switch item.Age() {
		case lessThanMonthOld:
			issuesM = append(issuesM, &item)
		case lessThanYearOld:
			issuesY = append(issuesY, &item)
		case moreThanYearOld:
			issuesO = append(issuesO, &item)
		}
	}

	fmt.Println("\nLess than a month old:")
	printIssues(issuesM)

	fmt.Println("\nLess than a year old:")
	printIssues(issuesY)

	fmt.Println("\nMore than a year old:")
	printIssues(issuesO)

}

func printIssues(issues []*issue) {
	sort.Slice(issues, func(i int, j int) bool {
		return issues[i].Number > issues[j].Number
	})

	for _, issue := range issues {
		fmt.Println(issue)
	}
}

//!-
