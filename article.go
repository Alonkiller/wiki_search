package wikisearch

type ArticleStorager interface {
	GetArticlesBy(keyWords string) ([]Article, error)
}

type Article struct {
	Link string
}
