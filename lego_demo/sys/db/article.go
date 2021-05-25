package db

import (
	"context"
	"lego_demo/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//创建文章
func (this *DB) DB_CreateOrUpDateArticle(data *pb.DB_ArticleData) (result *pb.DB_ArticleData, err error) {
	result = &pb.DB_ArticleData{
		Id:           data.Id,
		Title:        data.Title,
		Content:      data.Content,
		ShortContent: data.ShortContent,
		Images:       data.Images,
		CreationTime: data.CreationTime,
		AuthorId:     data.AuthorId,
		AuthorName:   data.AuthorName,
		AuthorAvatar: data.AuthorAvatar,
		GreatNum:     data.GreatNum,
	}
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

func (this *DB) DB_DelArticles(aId uint32) (err error) {
	_, err = this.mgo.DeleteOne(Sql_ArticleDataIdTable, bson.M{"_id": aId})
	return
}

func (this *DB) QueryArticle(aId uint32) (result *pb.DB_ArticleData, err error) {
	result = &pb.DB_ArticleData{}
	err = this.mgo.FindOne(Sql_ArticleDataIdTable, bson.M{"_id": aId}).Decode(result)
	return
}

func (this *DB) QueryArticles(aId []uint32) (result []*pb.DB_ArticleData, err error) {
	var (
		cursor *mongo.Cursor
	)
	result = make([]*pb.DB_ArticleData, 0)
	if cursor, err = this.mgo.Find(Sql_ArticleDataIdTable, bson.M{"_id": bson.M{"$in": aId}}); err == nil {
		if err = cursor.Err(); err == nil {
			for cursor.Next(context.Background()) {
				temp := &pb.DB_ArticleData{}
				cursor.Decode(temp)
				result = append(result, temp)
			}
		}
	}
	return
}

//查询最新的文章
func (this *DB) QueryArticlesByLast(start int64, num int64) (result []*pb.DB_ArticleData, err error) {
	var (
		cursor *mongo.Cursor
	)
	result = make([]*pb.DB_ArticleData, 0)
	if cursor, err = this.mgo.Find(Sql_ArticleDataIdTable, bson.M{}, new(options.FindOptions).SetSkip(start).SetLimit(num).
		SetSort(bson.M{"creationtime": -1}).SetProjection(bson.M{})); err == nil {
		if err = cursor.Err(); err == nil {
			for cursor.Next(context.Background()) {
				temp := &pb.DB_ArticleData{}
				cursor.Decode(temp)
				result = append(result, temp)
			}
		}
	}
	return
}
