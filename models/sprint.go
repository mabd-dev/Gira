package models

type Sprint struct {
	ID         string
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
	ID          string
	Summary     string
	Description string
	StoryPoints int
	Components  []string
	FixVersions []string
}
