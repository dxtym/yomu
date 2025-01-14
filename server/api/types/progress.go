package types

type UpdateProgressRequest struct {
	Manga   string `json:"manga"`
	Chapter string `json:"chapter"`
	Page    uint   `json:"page"`
}
