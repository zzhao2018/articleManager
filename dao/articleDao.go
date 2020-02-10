package dao

import (
	"wxGroupSend/model"
)

//插入数据
func InsertArticle(article *model.ArticleInfo){
	db.Create(article)
}

/*******************查询数据******************/
//查询所有文章
func SearchArticle()[]*model.ArticleInfo{
	var articleList []*model.ArticleInfo
	db.Find(&articleList)
	return articleList
}

//使用id查询文章
func SearchArticleById(id int)*model.ArticleInfo{
	var articleInfo model.ArticleInfo
	db.Where("id=?",id).First(&articleInfo)
	return &articleInfo
}

//随机取一条数据
func SearchArticleByRandom(typeid int)*model.ArticleInfo{
	var articelInfo model.ArticleInfo
	db.Order("rand()").Where("type=?",typeid).First(&articelInfo)
	return &articelInfo
}
