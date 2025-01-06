package types

type AddLibraryRequest struct {
	Manga      string `json:"manga"`
	CoverImage string `json:"cover_image"`
}

type GetLibraryResponse struct {
	MangaUrl   string `json:"manga_url"`
	CoverImage string `json:"cover_image"`
}
