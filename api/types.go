package api

type Project struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProjectResponse struct {
	IsLast   bool      `json:"isLast"`
	Projects []Project `json:"values"`
}
