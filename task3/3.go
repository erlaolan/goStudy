import (
	"gorm/global"
	"gorm/models"
	"log"
)

func main() {
	global.Connect()

	var postWithCommentCount models.PostWithCommentCount

	result := global.DB.Table("posts").
		Joins("left join comments on posts.id = comments.post_id").
		Select("posts.*, count(comments.id) as comment_count").
		Group("posts.id").
		Order("comment_count desc").
		Limit(1).
		Scan(&postWithCommentCount)

	if result.Error != nil {
		log.Fatalf("查询失败：%v", result.Error)
	}

	log.Printf("========== 评论数最多的文章 ==========\n")
	log.Printf("文章ID：%d\n", postWithCommentCount.ID)
	log.Printf("标题：%s\n", postWithCommentCount.Title)
	log.Printf("内容：%s\n", postWithCommentCount.Content)
	log.Printf("发布时间：%s\n", postWithCommentCount.CreatedAt)
	log.Printf("评论数量：%d\n", postWithCommentCount.CommentCount)
}