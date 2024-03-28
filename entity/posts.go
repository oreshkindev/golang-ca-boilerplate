package entity

type Posts struct {
	ID    int32  `json:"id"`
	Title string `gorm:"uniqueIndex:idx_post_title,sort:desc" json:"title"`
}

type (
	PostsUseCase interface {
		Get() ([]Posts, error)
		Post(*Posts) (*Posts, error)
	}

	PostsRepository interface {
		Get() ([]Posts, error)
		Post(*Posts) (*Posts, error)
	}
)
