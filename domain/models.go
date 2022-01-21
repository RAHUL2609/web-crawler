package domain

type UrlTemplate struct {
	UrlMap map[string][]string `json:"urlMap"`
}

type Error struct {
	ErrorMessage string `json:"errorMessage"`
}

type FetchUrl struct {
	Url	string `json:"url"`
}

type TaskResponse struct {
	TaskId        string `json:"taskId"`
}

type State string



