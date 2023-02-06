package graph

import (
	"context"
	"reflect"
	"testing"

	"github.com/ManuelLecaro/erpcore/internal/core/domain"
	serviceMock "github.com/ManuelLecaro/erpcore/internal/core/ports/services/mocks"
	model1 "github.com/ManuelLecaro/erpcore/internal/infra/graph/model"
	"github.com/stretchr/testify/mock"
)

// go run github.com/99designs/gqlgen generate

func Test_mutationResolver_CreateArticle(t *testing.T) {
	test_image := "image.image.jpg"

	returnedArticle := domain.Article{}

	articleMock := serviceMock.IArticleService{}
	articleMock.Mock.On("Create", mock.Anything, mock.Anything).Return(&returnedArticle, nil)

	type fields struct {
		Resolver *Resolver
	}
	type args struct {
		ctx   context.Context
		input model1.NewArticle
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model1.Article
		wantErr bool
	}{
		{
			name: "should work correctly",
			fields: fields{
				Resolver: &Resolver{
					ArticleService: &articleMock,
				},
			},
			args: args{
				ctx: context.Background(),
				input: model1.NewArticle{
					Name:        "test_article",
					Ean:         "test_ean",
					ImgURI:      &test_image,
					Description: "test_description",
					Categories: []*model1.CategoriesArticle{
						{
							Name: "father",
							Children: []*model1.CategoriesArticle{
								{
									Name:     "child_1",
									Children: []*model1.CategoriesArticle{},
								},
								{
									Name: "child_2",
									Children: []*model1.CategoriesArticle{
										{
											Name:     "grand_child_1",
											Children: []*model1.CategoriesArticle{},
										},
									},
								},
							},
						},
					},
					Images: []*model1.NewImage{
						{
							Name:        "test_image.jpg",
							Description: "test_image_description",
							URL:         "test_image_url.jpg",
						},
					},
				},
			},
			want: &model1.Article{
				ID:          "",
				Name:        "test_article",
				Ean:         "test_ean",
				Description: "test_description",
				Category: []*model1.Category{
					{
						Name:        "father",
						Description: "test_category_description",
						Category: []*model1.Category{
							{
								Name:     "child_1",
								Category: []*model1.Category{},
							},
							{
								Name: "child_2",
								Category: []*model1.Category{
									{
										Name:     "grand_child_1",
										Category: []*model1.Category{},
									},
								},
							},
						},
					}},
				Images: []*model1.Image{{
					ID:          "",
					Name:        "test_image.jpg",
					Description: "test_image_description",
					URL:         "test_image_url.jpg",
				}},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &mutationResolver{
				Resolver: tt.fields.Resolver,
			}
			got, err := r.CreateArticle(tt.args.ctx, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("mutationResolver.CreateArticle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mutationResolver.CreateArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
