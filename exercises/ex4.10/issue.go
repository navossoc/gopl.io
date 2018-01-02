package main

import (
	"fmt"
	"time"

	"github.com/navossoc/gopl.io/ch4/github"
)

type age int

const (
	lessThanMonthOld age = iota
	lessThanYearOld
	moreThanYearOld
)

type issue github.Issue

func (is *issue) Age() age {
	now := time.Now()
	switch {
	case now.AddDate(0, -1, 0).Before(is.CreatedAt):
		return lessThanMonthOld
	case now.AddDate(-1, 0, 0).Before(is.CreatedAt):
		return lessThanYearOld
	default:
		return moreThanYearOld
	}
}

func (is issue) String() string {
	return fmt.Sprintf("%s\t#%-5d %9.9s %.55s",
		is.CreatedAt.Format("2006-01-02 15:04:05"), is.Number, is.User.Login, is.Title)
}
