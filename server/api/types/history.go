package types

type GetHistoryResponse struct {
	Id     uint64 `json:"id"`
	Manga  string `json:"manga"`
	ReadAt string `json:"read_at"`
}
