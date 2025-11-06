package api

type EmptyResponse struct{}

type Project struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ProjectTypeKey string `json:"projectTypeKey"`
}

type ProjectResponse struct {
	IsLast   bool      `json:"isLast"`
	Projects []Project `json:"values"`
}

type Board struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"isPrivate"`
}

type BoardsResponse struct {
	IsLast bool    `json:"isLast"`
	Boards []Board `json:"values"`
}

type Sprint struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Goal      string `json:"goal"`
}

type SprintsResponse struct {
	IsLast  bool     `json:"isLast"`
	Sprints []Sprint `json:"values"`
}

type IssueAssignee struct {
	AccountID string `json:"accountId"`
	Name      string `json:"displayName"`
}

type IssueStatus struct {
	Name string `json:"name"`
}
type IssueFixVersion struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type IssueComponent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IssueFields struct {
	Assignee    IssueAssignee     `json:"assignee"`
	Status      IssueStatus       `json:"status"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
	Sprint      Sprint            `json:"sprint"`
	FixVersions []IssueFixVersion `json:"fixVersions"`
	StoryPoints float32           `json:"customfield_10024"`
	Components  []IssueComponent  `json:"components"`
}

type Issue struct {
	ID     string      `json:"id"`
	Fields IssueFields `json:"fields"`
}

type SprintIssuesResponse struct {
	Issues []Issue `json:"issues"`
}
