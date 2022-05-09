package main

import (
	"fmt"
	"strings"
)

type AddTimeZoneArguments struct {
	Label    string
	Location string
}

func HandleAddTimeZoneCommandArguments(args string) (AddTimeZoneArguments, error) {
	argsArray := strings.Fields(args)

	if len(argsArray) != 2 {
		return AddTimeZoneArguments{}, fmt.Errorf("Please provide a label and a time zone as in the tz database.\n")
	}

	loc := argsArray[1]
	if _, err := GetLocation(loc); err != nil {
		return AddTimeZoneArguments{}, err
	}

	commandArgs := AddTimeZoneArguments{
		Label:    argsArray[0],
		Location: loc,
	}

	return commandArgs, nil
}

type RemoveTimeZoneArguments struct {
	Label string
}

func HandleRemoveTimeZoneCommandArguments(args string) (RemoveTimeZoneArguments, error) {
	argsArray := strings.Fields(args)

	if len(argsArray) != 1 {
		return RemoveTimeZoneArguments{}, fmt.Errorf("Please provide a label.\n")
	}

	commandArgs := RemoveTimeZoneArguments{
		Label: argsArray[0],
	}

	return commandArgs, nil
}