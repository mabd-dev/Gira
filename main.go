package main

import (
	"fmt"

	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/config"
	"github.com/mabd-dev/gira/internal/logger"
	"github.com/mabd-dev/gira/internal/ui"
)

func main() {

	if err := createMockAPIClient(); err != nil {
		panic(err)
	}

	// if err := createRealApiClient(); err != nil {
	// 	panic(err)
	// }

	// getProjectsResponse, err := client.GetProjects()
	// if err != nil {
	// 	fmt.Printf("error getting projects, err=%s", err.Error())
	// }
	//
	// for _, project := range getProjectsResponse.Projects {
	// 	fmt.Println(project)
	// }
	//
	// softwareTeamProjectID := "10313"
	// getBoardsResponse, err := client.GetBoards(softwareTeamProjectID)
	// if err != nil {
	// 	fmt.Printf("error getting boards, err=%s", err.Error())
	// 	return
	// }
	//
	// for _, board := range getBoardsResponse.Boards {
	// 	fmt.Println(board)
	// }

	// softwareSprintBoardID := "54"
	// getSprintsResponse, err := client.GetSprints(softwareSprintBoardID)
	// if err != nil {
	// 	fmt.Printf("error getting sprints, err=%s", err.Error())
	// 	return
	// }
	//
	// for _, sprint := range getSprintsResponse.Sprints {
	// 	fmt.Println(sprint)
	// }

	// activeSprint, err := client.GetActiveSprint(softwareSprintBoardID)
	// if err != nil {
	// 	fmt.Printf("error getting active sprint, err=%s", err.Error())
	// 	return
	// }
	// fmt.Printf("active sprint: %v\n", activeSprint)

	// getSprintIssuesResponse, err := client.GetSprintIssues(1853)
	// if err != nil {
	// 	fmt.Printf("error getting sprint issues, err=%s", err.Error())
	// 	return
	// }
	//

	logger.Init(true, "/.config/gira/logs/")

	// fileData, err := os.ReadFile("samples/mockApiResponses/sprintIssues.json")
	// if err != nil {
	// 	fmt.Printf("error reading json file, err=%s", err.Error())
	// 	return
	// }
	//
	// var getSprintIssuesResponse api.SprintIssuesResponse
	// err = json.Unmarshal(fileData, &getSprintIssuesResponse)
	// if err != nil {
	// 	fmt.Printf("error unmarshal json data, err=%s", err.Error())
	// 	return
	// }
	//
	// sprint, err := models.FormatSprint(getSprintIssuesResponse)
	// if err != nil {
	// 	fmt.Printf("error formatting sprint, err=%s", err.Error())
	// 	return
	// }

	if err := ui.Render(); err != nil {
		fmt.Printf("failed to render using bubbletea, err=%s", err.Error())
		return
	}
}

func createMockAPIClient() error {
	_, err := api.NewMockClient()
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

func createRealApiClient() error {
	cred, err := config.Load()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err = api.NewClient(
		cred.Email,
		cred.Secret,
		cred.Domain,
	)
	if err != nil {
		fmt.Printf("error creating new client, err=%s", err.Error())
		return err
	}
	return nil
}
