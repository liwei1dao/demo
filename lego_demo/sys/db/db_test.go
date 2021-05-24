package db

import (
	"lego_demo/pb"
	"testing"
)

func Test_CreateArticle(t *testing.T) {
	if err := OnInit(map[string]interface{}{
		"MongodbUrl":      "mongodb://127.0.0.1:27017",
		"MongodbDatabase": "demo",
	}); err == nil {
		CreateOrUpDateArticle(&pb.DB_ArticleData{
			Title:    "测试",
			Content:  "测试内容",
			AuthorId: 142250,
		})
	}
}
