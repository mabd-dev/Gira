package api

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProjectResponse struct {
	IsLast   bool      `json:"isLast"`
	Projects []Project `json:"values"`
}

type Board struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	IsPrivate bool   `json:"isPrivate"`
}

type BoardsResponse struct {
	IsLast bool    `json:"isLast"`
	Boards []Board `json:"values"`
}

type Sprint struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type SprintsResponse struct {
	IsLast  bool    `json:"isLast"`
	Springs []Board `json:"values"`
}
