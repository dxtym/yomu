package types

type GetMangaResponse struct {
	Title       string `json:"title"`
	CoverImage  string `json:"cover_image"`
	Description string `json:"description"`
	Chapters    []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"chapters"`
}

type SearchMangaResponse struct {
	Manga      string `json:"manga"`
	CoverImage string `json:"cover_image"`
}
