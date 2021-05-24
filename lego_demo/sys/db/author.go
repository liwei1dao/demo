package db

import (
	"lego_demo/pb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (this *DB) QueryAuthor(uid uint32) (result *pb.DB_AuthorData, err error) {
	result = &pb.DB_AuthorData{}
	err = this.mgo.FindOne(Sql_AuthorDataIdTable, bson.M{"_id": uid}).Decode(result)
	return
}

func (this *DB) CreateAuthor(data *pb.DB_AuthorData) (err error) {
	_, err = this.mgo.InsertOne(Sql_AuthorDataIdTable, data)
	return
}

//更新用户文章列表
func (this *DB) DB_UpDateAuthorByArticle(AuthorId uint32, Article uint32) (result *pb.DB_AuthorData, err error) {
	result = &pb.DB_AuthorData{}
	err = this.mgo.FindOneAndUpdate(Sql_AuthorDataIdTable,
		bson.M{"_id": AuthorId},
		bson.M{"$addToSet": bson.M{
			"articleids": Article,
		}},
		new(options.FindOneAndUpdateOptions).SetReturnDocument(options.After)).Decode(result)
	return
}

//更新用户文章列表
func (this *DB) DB_RemoveAuthorByArticle(AuthorId uint32, Article uint32) (result *pb.DB_AuthorData, err error) {
	result = &pb.DB_AuthorData{}
	err = this.mgo.FindOneAndUpdate(Sql_AuthorDataIdTable,
		bson.M{"_id": AuthorId},
		bson.M{"$pull": bson.M{
			"articleids": Article,
		}},
		new(options.FindOneAndUpdateOptions).SetReturnDocument(options.After)).Decode(result)
	return
}
