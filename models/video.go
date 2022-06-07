package models

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"` //omitempty实现嵌套结构体的省略输出
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
	Title         string `json:"title"`
}

type VideoDao struct {
	VideoId       int64  `gorm:"colum:video_id"`
	AuthorId      int64  `gorm:"colum:author_id"`
	PlayUrl       string `gorm:"colum:play_url"`
	CoverUrl      string `gorm:"colum:cover_url"`
	FavoriteCount int64  `gorm:"colum:favorite_count"`
	CommentCount  int64  `gorm:"colum:comment_count"`
	Title         string `gorm:"colum:title"`
}

func (v VideoDao) TableName() string {
	return "videos"
}

//gorm:"-"表示可以忽略该字段不参与数据表读写
