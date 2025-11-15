package models

import (
	"fmt"
	"slices"
	"sort"
	"strings"

	"github.com/mabd-dev/gira/api"
)

func FormatProject(
	projectsResponse api.ProjectResponse,
) ([]Project, error) {
	projects := make([]Project, len(projectsResponse.Projects))

	for i, p := range projectsResponse.Projects {
		project := Project{
			ID:             p.ID,
			Name:           p.Name,
			ProjectTypeKey: p.ProjectTypeKey,
		}
		projects[i] = project
	}

	return projects, nil
}

func FormatBoards(
	boardsApiResponse api.BoardsResponse,
) ([]Board, error) {
	boards := make([]Board, len(boardsApiResponse.Boards))

	for i, p := range boardsApiResponse.Boards {
		board := Board{
			ID:        p.ID,
			Name:      p.Name,
			IsPrivate: p.IsPrivate,
		}
		boards[i] = board
	}

	return boards, nil
}

func FormatSprint(
	sprintIssuesResponse api.SprintIssuesResponse,
) (Sprint, error) {

	if len(sprintIssuesResponse.Issues) == 0 {
		return Sprint{}, nil
	}

	firstIssue := sprintIssuesResponse.Issues[0]

	sprint := Sprint{
		Name:       firstIssue.Fields.Sprint.Name,
		StartDate:  firstIssue.Fields.Sprint.StartDate,
		EndDate:    firstIssue.Fields.Sprint.EndDate,
		Goal:       firstIssue.Fields.Sprint.Goal,
		Developers: []Developer{},
	}

	devNames := []string{}
	developersByName := map[string]Developer{}

	for _, issue := range sprintIssuesResponse.Issues {
		devName := issue.Fields.Assignee.Name
		if len(strings.TrimSpace(devName)) == 0 {
			devName = "Unassigned"
		}

		if !slices.Contains(devNames, devName) {
			devNames = append(devNames, devName)
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
			ID:          issue.ID,
			Summary:     issue.Fields.Summary,
			Description: issue.Fields.Description,
			StoryPoints: int(issue.Fields.StoryPoints),
			Components:  parseIssueComponents(issue.Fields.Components),
			FixVersions: parseIssueFixVersions(issue.Fields.FixVersions),
		})
		developer.TasksByStatus[taskStatus] = tasks

		developersByName[devName] = developer
	}

	sort.Strings(devNames)

	for _, devName := range devNames {
		sprint.Developers = append(sprint.Developers, developersByName[devName])
	}

	return sprint, nil
}

func parseIssueComponents(cmps []api.IssueComponent) []string {
	result := []string{}
	for _, cmp := range cmps {
		result = append(result, cmp.Name)
	}

	return result
}

func parseIssueFixVersions(versions []api.IssueFixVersion) []string {
	result := []string{}
	for _, version := range versions {
		result = append(result, version.Name)
	}

	return result
}
