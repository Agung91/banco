package repo

import (
	"banco/common/config"
	"banco/common/infra/orm"
	"banco/domain/article"

	"gorm.io/gorm"
)

func NewRepoArticle(conf config.Config) article.ArticleRepo {
	db, err := orm.DbCon(conf)
	if err != nil {
		panic(err)
	}
	return &repo{
		db: db,
	}
}

type repo struct {
	db *gorm.DB
}

// Create implements artikel.ArticleRepo
func (*repo) Create(article.Artikel) (*string, error) {
	panic("unimplemented")
}

// Delete implements artikel.ArticleRepo
func (*repo) Delete(string) (*string, error) {
	panic("unimplemented")
}

// Read implements artikel.ArticleRepo
func (*repo) Read(string) (*string, error) {
	panic("unimplemented")
}

// Update implements artikel.ArticleRepo
func (*repo) Update(article.Artikel) (*string, error) {
	panic("unimplemented")
}