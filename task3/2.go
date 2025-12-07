package main

import (
	"gorm/global"
	"gorm/models"
	"log"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Age      int
	PostList []Post `gorm:"foreignKey:UserId"`
}

type Post struct {
	gorm.Model
	Title       string
	Content     string
	UserId      int
	User        User      `gorm:"foreignKey:UserId"`
	CommentList []Comment `gorm:"foreignKey:PostId"`
}

type Comment struct {
	gorm.Model
	Comment string
	PostId  int
	Post    Post `gorm:"foreignKey:PostId"`
}

func main() {
	global.Connect()
	// global.Migrate()
	// user := models.User{
	// 	Name: "张三",
	// 	Age:  20,
	// 	PostList: []models.Post{
	// 		{Title: "post1",
	// 			Content: "content1",
	// 			CommentList: []models.Comment{
	// 				{Comment: "pinglun1"},
	// 				{Comment: "pinglun2"},
	// 				{Comment: "pinglun3"},
	// 			}},
	// 		{Title: "post2",
	// 			Content: "content2",
	// 			CommentList: []models.Comment{
	// 				{Comment: "pinglun4"},
	// 				{Comment: "pinglun5"},
	// 				{Comment: "pinglun6"},
	// 			}},
	// 		{Title: "post3",
	// 			Content: "content3",
	// 			CommentList: []models.Comment{
	// 				{Comment: "pinglun7"},
	// 				{Comment: "pinglun8"},
	// 				{Comment: "pinglun9"},
	// 			},
	// 		},
	// 	},
	// }
	// global.DB.Create(&user)
	// comment := models.Comment{Comment: "pinglun10", PostId: 3}
	// global.DB.Create(&comment)
	var user models.User
	global.DB.Preload("PostList.CommentList").First(&user, 1)

	log.Printf("========== 用户 %s 的所有文章及评论 ==========\n", user.Name)
	for _, post := range user.PostList {
		log.Printf("文章ID：%d | 标题：%s\n", post.ID, post.Title)
		log.Printf("文章内容：%s\n", post.Content)
		log.Println("评论列表：")
		if len(post.CommentList) == 0 {
			log.Println("  暂无评论")
		} else {
			for _, comment := range post.CommentList {
				log.Printf("  - %s\n", comment.Comment)
			}
		}
		log.Println("----------------------------------------")
	}

}
