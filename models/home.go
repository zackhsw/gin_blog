package models

import (
	"fmt"
	"gin_blog/utils"
	"strings"
)

type HomeBlockParam struct {
	Article *Article
	TagLinks []*TagLink
	CreateTimeStr string
	Link string

	UpdateLink string
	DeleteLink string

	IsLogin bool
}

type TagLink struct{
	TagName string
	TagUrl string
}

func createTagsLinks(tagStr string)[]*TagLink{
	var tagLinks = make([]*TagLink,0,strings.Count(tagStr,"&"))
	tagList :=strings.Split(tagStr,"&")
	for _, tag := range tagList{
		tagLinks = append(tagLinks, &TagLink{tag,"/?tag="+tag})
	}
	return tagLinks
}

func GenHomeBlocks(articleList []*Article, isLogin bool)(ret []*HomeBlockParam){
	ret = make([]*HomeBlockParam,0,len(articleList))
	for _, art := range articleList{
		homeParam := HomeBlockParam{
			Article:art,
			IsLogin:isLogin,
		}
		homeParam.TagLinks = createTagsLinks(art.Tags)
		homeParam.CreateTimeStr = utils.SwitchTimeStampToStr(art.CreateTime)

		homeParam.Link = fmt.Sprintf("/show/%d", art.Id)
		homeParam.UpdateLink = fmt.Sprintf("/article/update?id=%d", art.Id)
		homeParam.DeleteLink = fmt.Sprintf("/article/delete?id=%d", art.Id)
		ret = append(ret, &homeParam) // 不再需要动态扩容
	}
	return
}