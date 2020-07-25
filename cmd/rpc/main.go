package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"

	"github.com/public-awesome/stakewatcher/models"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "server")
		os.Exit(1)
	}
	serverAddress := os.Args[1]

	client, err := rpc.DialHTTP("tcp", serverAddress+":1666")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	upvote := models.Upvote{
		ID:            "123",
		Height:        1,
		VendorID:      1,
		PostID:        "postid",
		Creator:       "creator",
		RewardAddress: "reward",
		VoteNumber:    1,
		VoteAmount:    1000,
		VoteDenom:     "ufuel",
		DepositAmount: 100,
		DepositDenom:  "ufuel",
		Body:          "body body",
	}

	var reply int

	err = client.Call("StakeWatcher.Upvoted", upvote, &reply)
	if err != nil {
		log.Fatal("StakeWatcher error:", err)
	}
	fmt.Println(reply)
}
