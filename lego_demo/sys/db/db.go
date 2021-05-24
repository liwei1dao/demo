package db

import (
	"context"
	"fmt"
	"lego_demo/pb"
	"math/rand"
	"time"

	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/mgo"
	"go.mongodb.org/mongo-driver/bson"
)

func newSys(options Options) (sys *DB, err error) {
	sys = &DB{options: options}
	if sys.mgo, err = mgo.NewSys(
		mgo.SetMongodbUrl(options.MongodbUrl),
		mgo.SetMongodbDatabase(options.MongodbDatabase),
		mgo.SetTimeOut(options.TimeOut),
	); err == nil {
		err = sys.checkDbInit()
	}
	return
}

type DB struct {
	options Options
	mgo     mgo.IMongodb
}

//校验数据库初始化工作是否完成
func (this DB) checkDbInit() (err error) {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*60)
	count, err := this.mgo.CountDocuments(Sql_ArticleIdDataIdTable, bson.M{})
	if err != nil || count == 0 {
		//批量插入数据
		leng := 1000000
		cIds := make([]interface{}, leng)
		for i, _ := range cIds {
			cIds[i] = 1000000 + i
		}
		data := make([]interface{}, leng)
		r := rand.New(rand.NewSource(time.Now().Unix()))
		n := 0
		for _, i := range r.Perm(leng) {
			data[n] = bson.M{"_id": i}
			n++
		}
		var (
			err error
		)
		begin := time.Now()
		if _, err = this.mgo.InsertManyByCtx(Sql_ArticleIdDataIdTable, ctx, data); err != nil {
			return fmt.Errorf("CheckDbInit  err=%s", err.Error())
		}
		log.Debugf("CheckDbInit succ time consuming:%v", time.Now().Sub(begin))
	}
	return
}

//创建文章
func (this *DB) CreateOrUpDateArticle(data *pb.DB_ArticleData) (result *pb.DB_ArticleData, err error) {
	if result, err = this.DB_CreateOrUpDateArticle(data); err == nil {
		_, err = this.DB_UpDateAuthorByArticle(result.AuthorId, result.Id)
	}
	return
}

//删除文章
func (this *DB) DeleteArticle(AuthoIrdId uint32, ArticlesId uint32) (err error) {
	if err = this.DB_DelArticles(ArticlesId); err == nil {
		_, err = this.DB_RemoveAuthorByArticle(AuthoIrdId, ArticlesId)
	}
	return
}
