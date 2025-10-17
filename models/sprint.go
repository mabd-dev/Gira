package models

import (
	"fmt"
	"strings"

	"github.com/mabd-dev/gira/api"
)

type Sprint struct {
	Name       string
	StartDate  string
	EndDate    string
	Goal       string
	Developers []Developer
}

type Developer struct {
	Name          string
	TasksByStatus map[TaskStatus][]DeveloperTask
}

// TODO: add components, fixed-version
type DeveloperTask struct {
	Summary     string
	Description string
	StoryPoints int
}

func FormatSprint(
	sprintIssuesResponse api.SprintIssuesResponse,
) (Sprint, error) {

	if len(sprintIssuesResponse.Issues) == 0 {
		return Sprint{}, nil
	}

	firstIssue := sprintIssuesResponse.Issues[0]

	sprint := Sprint{
		Name:      firstIssue.Fields.Sprint.Name,
		StartDate: firstIssue.Fields.Sprint.StartDate,
		EndDate:   firstIssue.Fields.Sprint.EndDate,
		Goal:      firstIssue.Fields.Sprint.Goal,
	}

	developersByName := map[string]Developer{}

	for _, issue := range sprintIssuesResponse.Issues {
		devName := issue.Fields.Assignee.Name
		if len(strings.TrimSpace(devName)) == 0 {
			devName = "Unassigned"
		}

		_, ok := developersByName[devName]
		if !ok {
			developersByName[devName] = Developer{
				Name:          devName,
				TasksByStatus: make(map[TaskStatus][]DeveloperTask),
			}
		}

		developer := developersByName[devName]

		taskStatus, err := getTaskStatusFrom(issue.Fields.Status.Name)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		_, ok = developer.TasksByStatus[taskStatus]
		if !ok {
			developer.TasksByStatus[taskStatus] = []DeveloperTask{}
		}

		tasks := developer.TasksByStatus[taskStatus]
		tasks = append(tasks, DeveloperTask{
			Summary:     issue.Fields.Summary,
			Description: issue.Fields.Description,
			StoryPoints: int(issue.Fields.StoryPoints),
		})
		developer.TasksByStatus[taskStatus] = tasks

		developersByName[devName] = developer
	}

	developers := []Developer{}
	for _, dev := range developersByName {
		developers = append(developers, dev)
	}
	sprint.Developers = developers

	return sprint, nil
}
