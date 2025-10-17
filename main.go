package main

import (
	"fmt"

	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/config"
)

func main() {
	cred, err := config.Load()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client, err := api.NewClient(
		cred.Email,
		cred.Secret,
		cred.Domain,
	)
	if err != nil {
		fmt.Printf("error creating new client, err=%s", err.Error())
		return
	}

	// getProjectsResponse, err := client.GetProjects()
	// if err != nil {
	// 	fmt.Printf("error getting projects, err=%s", err.Error())
	// }
	//
	// for _, project := range getProjectsResponse.Projects {
	// 	fmt.Println(project)
	// }

	getBoardsResponse, err := client.GetBoards("10313")
	if err != nil {
		fmt.Printf("error getting boards, err=%s", err.Error())
	}

	for _, board := range getBoardsResponse.Boards {
		fmt.Println(board)
	}
}
