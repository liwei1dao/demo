package db

import (
	"lego_demo/pb"

	"github.com/liwei1dao/lego/core"
)

type (
	ISys interface {
		//创建作者
		CreateAuthor(data *pb.DB_AuthorData) (err error)
		//删除文章
		DeleteArticle(AuthoIrdId uint32, ArticlesId uint32) (err error)
		//查询作者
		QueryAuthor(uid uint32) (result *pb.DB_AuthorData, err error)
		//创建文章
		CreateOrUpDateArticle(data *pb.DB_ArticleData) (result *pb.DB_ArticleData, err error)
		//查询文章
		QueryArticle(aId uint32) (result *pb.DB_ArticleData, err error)
		//查询文章
		QueryArticles(aId []uint32) (result []*pb.DB_ArticleData, err error)
		//查询最新文章
		QueryArticlesByLast(start int64, num int64) (result []*pb.DB_ArticleData, err error)
	}
	ArticleId struct {
		Id uint32 `bson:"_id"`
	}
)

const (
	Sql_ArticleIdDataIdTable core.SqlTable = "articleid" //文章Id
	Sql_ArticleDataIdTable   core.SqlTable = "article"   //文章
	Sql_AuthorDataIdTable    core.SqlTable = "author"    //作者
	Sql_CommentDataIdTable   core.SqlTable = "comment"   //评论
)

var (
	defsys ISys
)

func OnInit(config map[string]interface{}, option ...Option) (err error) {
	defsys, err = newSys(newOptions(config, option...))
	return
}

func NewSys(option ...Option) (sys ISys, err error) {
	sys, err = newSys(newOptionsByOption(option...))
	return
}

/*								 相关接口
 * _______________#########_______________________
 * ______________############_____________________
 * ______________#############____________________
 * _____________##__###########___________________
 * ____________###__######_#####__________________
 * ____________###_#######___####_________________
 * ___________###__##########_####________________
 * __________####__###########_####_______________
 * ________#####___###########__#####_____________
 * _______######___###_########___#####___________
 * _______#####___###___########___######_________
 * ______######___###__###########___######_______
 * _____######___####_##############__######______
 * ____#######__#####################_#######_____
 * ____#######__##############################____
 * ___#######__######_#################_#######___
 * ___#######__######_######_#########___######___
 * ___#######____##__######___######_____######___
 * ___#######________######____#####_____#####____
 * ____######________#####_____#####_____####_____
 * _____#####________####______#####_____###______
 * ______#####______;###________###______#________
 * ________##_______####________####______________
 */
//删除文章
func DeleteArticle(AuthoIrdId uint32, ArticlesId uint32) (err error) {
	return defsys.DeleteArticle(AuthoIrdId, ArticlesId)
}

///创建作者
func CreateAuthor(data *pb.DB_AuthorData) (err error) {
	return defsys.CreateAuthor(data)
}
func QueryAuthor(uid uint32) (result *pb.DB_AuthorData, err error) {
	return defsys.QueryAuthor(uid)
}

///创建文章
func CreateOrUpDateArticle(data *pb.DB_ArticleData) (result *pb.DB_ArticleData, err error) {
	return defsys.CreateOrUpDateArticle(data)
}

func QueryArticle(aId uint32) (result *pb.DB_ArticleData, err error) {
	return defsys.QueryArticle(aId)
}

func QueryArticles(aId []uint32) (result []*pb.DB_ArticleData, err error) {
	return defsys.QueryArticles(aId)
}

//查询最新文章
func QueryArticlesByLast(start int64, num int64) (result []*pb.DB_ArticleData, err error) {
	return defsys.QueryArticlesByLast(start, num)
}
