package db

import (
	"lego_demo/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//创建文章
func (this *DB) DB_CommentArticle(data *pb.DB_CommentData) (result *pb.DB_ArticleData, err error) {
	result = &pb.DB_ArticleData{}
	if result.Id == 0 {
		var newuId ArticleId
		if data.Id == 0 {
			if err = this.mgo.FindOneAndDelete(Sql_ArticleIdDataIdTable, bson.D{}).Decode(&newuId); err != nil {
				return
			}
		}
		result.Id = newuId.Id
	}
	_, err = this.mgo.UpdateOne(Sql_ArticleDataIdTable, bson.M{"_id": result.Id}, bson.M{
		"$set": bson.M{
			"title":        result.Title,
			"content":      result.Content,
			"shortcontent": result.ShortContent,
			"images":       result.Images,
			"creationtime": result.CreationTime,
			"authorId":     result.AuthorId,
			"authorname":   result.AuthorName,
			"authoravatar": result.AuthorAvatar,
			"greatnum":     result.GreatNum}},
		new(options.UpdateOptions).SetUpsert(true))
	return
}
