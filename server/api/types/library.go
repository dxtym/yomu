package types

type AddLibraryRequest struct {
	Manga      string `json:"manga" binding:"required"`
	CoverImage string `json:"cover_image" binding:"required"`
}

type GetLibraryResponse struct {
	Manga      string `json:"manga"`
	CoverImage string `json:"cover_image"`
}

type RemoveLibraryRequest struct {
	Manga string `json:"manga" binding:"required"`
}
