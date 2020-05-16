package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/bloveless/tweetgo"
)

func statusesUpdate(c config, status string) {
	tc := getTwitterClient(c)

	input := tweetgo.StatusesUpdateInput{
		Status: tweetgo.String(status),
	}

	output, err := tc.StatusesUpdatePost(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%+v\n", output)
}

func listsList(c config) {
	tc := getTwitterClient(c)

	input := tweetgo.ListsListInput{}

	output, err := tc.ListsListGet(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%+v\n", output)
}

func listsMembers(c config) {
	tc := getTwitterClient(c)

	input := tweetgo.ListsMembersInput{
		ListID:     tweetgo.Int64(1130185227375038465),
		Cursor:     tweetgo.Int(-1),
		SkipStatus: tweetgo.Bool(true),
	}

	output, err := tc.ListsMembersGet(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%+v\n", output)
}

func listsMembersShow(c config) {
	tc := getTwitterClient(c)

	input := tweetgo.ListsMembersShowInput{
		ListID:     tweetgo.Int64(1130185227375038465),
		SkipStatus: tweetgo.Bool(true),
		ScreenName: tweetgo.String("twitterdev"),
	}

	output, err := tc.ListsMembersShowGet(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n\n%+v\n", output)
}

func streamTweets(c config, hashtag string) {
	fmt.Printf("Beginning to stream #%s tweets\n", hashtag)
	tc := getTwitterClient(c)

	input := tweetgo.StatusesFilterInput{
		Track: tweetgo.String(hashtag),
	}

	output, err := tc.StatusesFilterPostRaw(input)
	if err != nil {
		panic(err)
	}

	for {
		tweet := tweetgo.StatusesFilterOutput{}
		err := json.NewDecoder(output.Body).Decode(&tweet)
		if err == io.EOF {
			fmt.Println("End of file")
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("%#v\n\n", tweet)
	}
}
