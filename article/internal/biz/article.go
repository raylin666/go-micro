package biz

import (
	"context"
	"mt/internal/repositories/dbrepo/model"
	"mt/pkg/logger"
	"mt/pkg/utils"
	"strconv"
)

// DataSetArticleListParams 文章列表参数
type DataSetArticleListParams struct {
	Page int32 `form:"page" json:"page"`
	Size int32 `form:"size" json:"size"`
}

// DataSetArticleListResult 文章列表结果集
type DataSetArticleListResult struct {
	List  []*DataSetArticleListItem `json:"list"`
	Count int64                     `json:"count"`
}

// DataSetArticleListItem 文章列表 Item
type DataSetArticleListItem struct {
	Id                string                              `json:"id"`
	Title             string                              `json:"title"`
	Author            string                              `json:"author"`
	Summary           string                              `json:"summary"`
	Cover             string                              `json:"cover"`
	Time              string                              `json:"time"`
	Category          []*DataSetArticleListToCategoryItem `json:"category"`
	RecommendFlag     bool                                `json:"recommend_flag"`
	ViewCount         int                                 `json:"view_count"`
	CommentCount      int                                 `json:"comment_count"`
	CollectionCount   int                                 `json:"collection_count"`
	ZanCount          int                                 `json:"zan_count"`
	ShareCount        int                                 `json:"share_count"`
	UserId            string                              `json:"user_id"`
	Avatar            string                              `json:"avatar"`
	PublisherUsername string                              `json:"publisher_username"`
}

type DataSetArticleListToCategoryItem struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}

type DataSetArticleInfoResult struct {
	Id                 string                              `json:"id"`
	Title              string                              `json:"title"`
	Author             string                              `json:"author"`
	Summary            string                              `json:"summary"`
	Cover              string                              `json:"cover"`
	Time               string                              `json:"time"`
	Date               string                              `json:"date"`
	Category           []*DataSetArticleListToCategoryItem `json:"category"`
	RecommendFlag      bool                                `json:"recommend_flag"`
	ViewCount          int                                 `json:"view_count"`
	CommentCount       int                                 `json:"comment_count"`
	CollectionCount    int                                 `json:"collection_count"`
	ZanCount           int                                 `json:"zan_count"`
	ShareCount         int                                 `json:"share_count"`
	UserId             string                              `json:"user_id"`
	Avatar             string                              `json:"avatar"`
	PublisherUsername  string                              `json:"publisher_username"`
	Content            string                              `json:"content"`
	Source             string                              `json:"source"`
	SourceUrl          string                              `json:"source_url"`
	PrevArticle        *DataSetBasicIntroduceArticle       `json:"prev_article"`
	NextArticle        *DataSetBasicIntroduceArticle       `json:"next_article"`
	CopyrightAuthor    string                              `json:"copyright_author"`
	CopyrightArticleId string                              `json:"copyright_article_id"`
	CopyrightLink      string                              `json:"copyright_link"`
	CopyrightStatement string                              `json:"copyright_statement"`
	ContentLength      string                              `json:"content_length"`
}

type DataSetBasicIntroduceArticle struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type ArticleRepo interface {
	// List 文章列表
	List(ctx context.Context, page, size int32) (m []*model.Article, count int64, err error)
	// CategoryList 文章分类列表
	CategoryList(ctx context.Context, status int32) (m []*model.ArticleCategory, count int64, err error)
	// ArticleByCategory 获取文章对应分类
	ArticleByCategory(ctx context.Context, aid int) ([]*model.ArticleCategory, error)
	// Info 文章详情
	Info(ctx context.Context, id int) (m *model.Article, em *model.ArticleExtend, err error)
	// GetPrevArticle 获取上一篇文章
	GetPrevArticle(ctx context.Context, id int) *model.Article
	// GetNextArticle 获取下一篇文章
	GetNextArticle(ctx context.Context, id int) *model.Article
}

type ArticleUsecase struct {
	repo ArticleRepo
	log  *logger.Logger
}

func NewArticleUsecase(repo ArticleRepo, logger *logger.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo, log: logger}
}

// List 文章列表
func (uc *ArticleUsecase) List(ctx context.Context, params *DataSetArticleListParams) (result *DataSetArticleListResult, err error) {
	m, count, err := uc.repo.List(ctx, params.Page, params.Size)
	if err != nil {
		return nil, err
	}

	result = new(DataSetArticleListResult)
	result.Count = count
	for _, article := range m {
		var category []*DataSetArticleListToCategoryItem
		categoryModel, _ := uc.repo.ArticleByCategory(ctx, article.ID)
		for _, cm := range categoryModel {
			category = append(category, &DataSetArticleListToCategoryItem{
				Id:    strconv.Itoa(cm.ID),
				Name:  cm.Name,
				Color: cm.Color,
			})
		}

		result.List = append(result.List, &DataSetArticleListItem{
			Id:              strconv.Itoa(article.ID),
			Title:           article.Title,
			Author:          article.Author,
			Avatar:          "http://cdn.ls331.com/micro/2022/1129/474f0018-a88e-44f3-a6b3-c6e6922af1fc.png",
			Cover:           article.Cover,
			Time:            utils.HowLongAgo(article.CreatedAt),
			Category:        category,
			Summary:         article.Summary,
			RecommendFlag:   article.RecommendFlag == 1,
			ViewCount:       article.ViewCount,
			CommentCount:    article.CommentCount,
			CollectionCount: article.CollectionCount,
			ZanCount:        article.ZanCount,
			ShareCount:      article.ShareCount,
		})
	}
	return result, nil
}

// Info 文章详情
func (uc *ArticleUsecase) Info(ctx context.Context, id int) (*DataSetArticleInfoResult, error) {
	m, em, err := uc.repo.Info(ctx, id)
	if err != nil {
		return nil, err
	}

	var category []*DataSetArticleListToCategoryItem
	categoryModel, _ := uc.repo.ArticleByCategory(ctx, m.ID)
	for _, cm := range categoryModel {
		category = append(category, &DataSetArticleListToCategoryItem{
			Id:    strconv.Itoa(cm.ID),
			Name:  cm.Name,
			Color: cm.Color,
		})
	}

	var result = new(DataSetArticleInfoResult)
	var prevArticle = uc.repo.GetPrevArticle(ctx, m.ID)
	var nextArticle = uc.repo.GetNextArticle(ctx, m.ID)
	result.Id = strconv.Itoa(m.ID)
	result.Cover = m.Cover
	result.Title = m.Title
	result.Time = utils.HowLongAgo(m.CreatedAt)
	result.Date = "2020年01月23日"
	result.Category = category
	result.Summary = m.Summary
	result.Avatar = "http://cdn.ls331.com/micro/2022/1129/474f0018-a88e-44f3-a6b3-c6e6922af1fc.png"
	result.Author = m.Author
	result.ZanCount = m.ZanCount
	result.CollectionCount = m.CollectionCount
	result.CommentCount = m.CommentCount
	result.ViewCount = m.ViewCount
	result.Content = em.Content
	result.Source = em.Source
	result.SourceUrl = em.SourceUrl
	if prevArticle != nil {
		result.PrevArticle = &DataSetBasicIntroduceArticle{ Id: strconv.Itoa(prevArticle.ID), Title: prevArticle.Title }
	}
	if nextArticle != nil {
		result.NextArticle = &DataSetBasicIntroduceArticle{ Id: strconv.Itoa(nextArticle.ID), Title: nextArticle.Title }
	}
	result.CopyrightAuthor = m.Author
	result.CopyrightArticleId = result.Id
	result.CopyrightStatement = "本博客所有文章除特別声明外，均采用 CC BY 4.0 许可协议。转载请注明来源 林山 !"
	result.CopyrightLink = "http://127.0.0.1:3000/article/info/1"
	return result, nil
}
