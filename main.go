package main

import (
	"fmt"
	"strconv"

	"github.com/mabd-dev/gira/api"
	"github.com/mabd-dev/gira/config"
	"github.com/mabd-dev/gira/models"
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
	//
	// getBoardsResponse, err := client.GetBoards("10313")
	// if err != nil {
	// 	fmt.Printf("error getting boards, err=%s", err.Error())
	// 	return
	// }
	//
	// for _, board := range getBoardsResponse.Boards {
	// 	fmt.Println(board)
	// }

	// getSprintsResponse, err := client.GetSprints("54")
	// if err != nil {
	// 	fmt.Printf("error getting sprints, err=%s", err.Error())
	// 	return
	// }
	//
	// for _, sprint := range getSprintsResponse.Sprints {
	// 	fmt.Println(sprint)
	// }
	//

	// activeSprint, err := client.GetActiveSprint("54")
	// if err != nil {
	// 	fmt.Printf("error getting active sprint, err=%s", err.Error())
	// 	return
	// }
	// fmt.Printf("active sprint: %v\n", activeSprint)

	getSprintIssuesResponse, err := client.GetSprintIssues(1853)
	if err != nil {
		fmt.Printf("error getting sprint issues, err=%s", err.Error())
		return
	}

	// for _, issue := range getSprintIssuesResponse.Issues {
	// 	fmt.Println(issue)
	// }

	sprint, err := models.FormatSprint(getSprintIssuesResponse)
	if err != nil {
		fmt.Printf("error formatting sprint, err=%s", err.Error())
		return
	}

	printSprint(sprint)

}

func printSprint(sprint models.Sprint) {
	fmt.Println(sprint.Name)
	fmt.Println(sprint.StartDate + " | " + sprint.EndDate)
	fmt.Println(sprint.Goal)
	fmt.Println("----------------------")

	for _, dev := range sprint.Developers {
		fmt.Println(dev.Name + ":")

		for _, status := range models.TasksInOrder {
			tasks := dev.TasksByStatus[status]
			if len(tasks) == 0 {
				continue
			}

			fmt.Println("  " + string(status))
			for i, task := range tasks {
				fmt.Println("    " + strconv.Itoa(i+1) + " - " + task.Summary)
			}
		}

		fmt.Println("***************")
	}
}
