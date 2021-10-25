package elasticsearch

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/olivere/elastic/config"
	"log"
)

func ConnectIndex(index string, ctx context.Context, mapping string) *elastic.Client {
	//cfg := &config.Config{
	//	URL: "http://localhost:9200",
	//	//Username: ,
	//	//Password: password,
	//	Index: "sniff=false",
	//}
	cfg, _ := config.Parse("http://localhost:9200/index?sniff=false")
	client, err := elastic.NewClientFromConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}
	exists, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex(index).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Fatal(err)
		}
		if !createIndex.Acknowledged {
		}
	}
	return client
}