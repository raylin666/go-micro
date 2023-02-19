package service

import (
	"context"
	v1 "github.com/raylin666/go-micro-protoc/api/article/v1"
	"mt/internal/biz"
	"mt/internal/constant/defined"
	"strconv"
)

type ArticleService struct {
	v1.UnimplementedArticleServer

	uc *biz.ArticleUsecase
}

func NewArticleService(uc *biz.ArticleUsecase) *ArticleService {
	return &ArticleService{uc: uc}
}

// List 文章列表
func (s *ArticleService) List(ctx context.Context, req *v1.ListRequest) (*v1.ListResponse, error) {
	result, err := s.uc.List(ctx, &biz.DataSetArticleListParams{Page: req.GetPage(), Size: req.GetSize()})
	if err != nil {
		return nil, err
	}

	var resp = new(v1.ListResponse)
	resp.Count = result.Count
	resp.Page = req.GetPage()
	resp.Size = req.GetSize()
	for _, item := range result.List {
		var category []*v1.ArticleCategoryItem
		for _, cm := range item.Category {
			category = append(category, &v1.ArticleCategoryItem{
				Id:    cm.Id,
				Name:  cm.Name,
				Color: cm.Color,
			})
		}

		resp.List = append(resp.List, &v1.ArticleListItem{
			Id:              item.Id,
			Title:           item.Title,
			Author:          item.Author,
			Avatar:          item.Avatar,
			Summary:         item.Summary,
			Cover:           item.Cover,
			Time:            item.Time,
			Category:        category,
			ViewCount:       int32(item.ViewCount),
			CollectionCount: int32(item.CollectionCount),
			ZanCount:        int32(item.ZanCount),
			ShareCount:      int32(item.ShareCount),
			CommentCount:    int32(item.CommentCount),
		})
	}
	return resp, nil
}

// Info 文章详情
func (s *ArticleService) Info(ctx context.Context, req *v1.InfoRequest) (*v1.InfoResponse, error) {
	id, err := strconv.Atoi(req.GetId())
	if err != nil {
		return nil, defined.ErrorIdInvalidValueError
	}

	result, err := s.uc.Info(ctx, id)
	if err != nil {
		return nil, err
	}

	var category []*v1.ArticleCategoryItem
	for _, cm := range result.Category {
		category = append(category, &v1.ArticleCategoryItem{
			Id:    cm.Id,
			Name:  cm.Name,
			Color: cm.Color,
		})
	}

	var resp = &v1.InfoResponse{
		Id:                 result.Id,
		Title:              result.Title,
		Author:             result.Author,
		Avatar:             result.Avatar,
		Summary:            result.Summary,
		Cover:              result.Cover,
		Time:               result.Time,
		Date: 				result.Date,
		Category:           category,
		ViewCount:          int32(result.ViewCount),
		CollectionCount:    int32(result.CollectionCount),
		CommentCount:       int32(result.CommentCount),
		ShareCount:         int32(result.ShareCount),
		Source:             result.Source,
		SourceUrl:          result.SourceUrl,
		Content:            result.Content,
		CopyrightAuthor:    result.CopyrightAuthor,
		CopyrightArticleId: result.CopyrightArticleId,
		CopyrightStatement: result.CopyrightStatement,
		CopyrightLink:      result.CopyrightLink,
	}

	if result.PrevArticle != nil {
		resp.PrevArticle = &v1.BasicIntroduceArticleItem{ Id: result.PrevArticle.Id, Title: result.PrevArticle.Title }
	}
	if result.NextArticle != nil {
		resp.NextArticle = &v1.BasicIntroduceArticleItem{ Id: result.NextArticle.Id, Title: result.NextArticle.Title }
	}
	return resp, nil
}
