package main

import (
	"fmt"
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

var DB *gorm.DB

func Migrate() {
	err := DB.AutoMigrate(&models.Comment{}, &models.Post{}, &models.User{})
	if err != nil {
		log.Fatalf("数据库迁移失败 %s", err)
		return
	}
	log.Fatalf("数据库迁移成功")
}

func Connect() {

	dsn := "root:123456@tcp(localhost:3306)/gorm"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}

func main() {
	Connect()
	Migrate()
}
