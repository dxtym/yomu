package types

type GetMangaResponse struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	CoverImage  string `json:"cover_image"`
	Description string `json:"description"`
	Chapters    []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"chapters"`
}

type SearchMangaResponse struct {
	MangaUrl   string `json:"manga_url"`
	CoverImage string `json:"cover_image"`
}
