package models

type CommentModel struct {
	Model
	SubComments        []*CommentModel `gorm:"foreignkey:ParentCommentID;comment:父评论" json:"sub_comments"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID;comment:子评论" json:"comment_model"`
	ParentCommentID    *uint           `gorm:"column:parent_comment_id;type:int;comment:父评论ID" json:"parent_comment_id"`

	Content      string       `gorm:"column:content;type:varchar(256);not null;comment:内容" json:"content"`
	DiggCount    int          `gorm:"column:digg_count;type:int;not null;default:0;comment:点赞次数" json:"digg_count"`
	CommentCount int          `gorm:"column:comment_count;type:int;not null;default:0;comment:评论次数" json:"comment_count"`
	Article      ArticleModel `gorm:"foreignKey:ArticleID;comment:关联文章" json:"-"`
	ArticleID    uint         `json:"article_id"`
	User         UserModel    `gorm:"foreignKey:UserID;comment:关联用户" json:"user_model"`
	UserID       uint         `json:"user_id"`
}
