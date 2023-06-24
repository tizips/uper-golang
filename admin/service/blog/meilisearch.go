package blog

import (
	"github.com/herhe-com/framework/facades"
	"github.com/meilisearch/meilisearch-go"
	"github.com/tizips/uper-go/model"
)

func NewMeilisearch() *meilisearch.Client {

	return meilisearch.NewClient(meilisearch.ClientConfig{
		Host:   facades.Cfg.GetString("database.meilisearch.host"),
		APIKey: facades.Cfg.GetString("database.meilisearch.key"),
	})
}

func SearchIndexForArticle() *meilisearch.Index {

	prefix := facades.Cfg.GetString("app.name")

	client := NewMeilisearch()

	index := client.Index(prefix + "_" + model.TableBlgArticle)

	return index
}
