package models

import (
	"gin_blog/dao"
	"gin_blog/gb"
)

const (
	pageSize = 4
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	CreateTime int64 `db:"create_time"`
	Status     int
}

func AddArticle(article Article) (int64, error) {
	i, err := insertArticle(article)
	return i, err
}

func insertArticle(article Article) (int64, error) {
	return dao.ModifyDB("insert into article(title,tags,short,content,author,create_time,status) values(?,?,?,?,?,?,?)",
		article.Title, article.Tags, article.Short, article.Content, article.Author, article.CreateTime, article.Status)
}

func QueryArticleWithPage(pageNum int) (articleList []*Article, err error) {
	sqlStr := "select id, title, tags, short,content,author,create_time from article limit ?,?"
	articleList, err = queryArticleWithCon(pageNum, sqlStr)
	if err != nil {
		return nil, err
	}
	return articleList, nil
}
func QueryCurrUserArticleWithPage(username string, pageNum int) (articleList []*Article, err error) {
	sqlStr := "select id,title,tags,short,content,author,create_time from article where author=? limit ?,?"
	articleList, err = queryArticleWithCon(pageNum, sqlStr, username)
	if err != nil {
		gb.Logger.Error("queryArticleWithCon, ", err)
		return nil, err
	}
	gb.Logger.Debug("QueryCurrUserArticleWithPage,", articleList)
	return articleList, nil
}

func queryArticleWithCon(pageNum int, sqlStr string, args ...interface{}) (articleList []*Article, err error) {
	pageNum--
	args = append(args, pageNum*pageSize, pageSize)
	gb.Logger.Debug("queryArticleWithCon", sqlStr, args)
	err = dao.QueryRows(&articleList, sqlStr, args...)
	gb.Logger.Debug("dao .QueryRows,", articleList)
	return
}
