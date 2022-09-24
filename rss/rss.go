package rss

import (
	"fmt"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
	"github.com/mmcdole/gofeed"
	"github.com/mmcdole/gofeed/rss"
)

type MyTranslator struct {
	defaultTranslator *gofeed.DefaultRSSTranslator
}

func NewMyTranslator() *MyTranslator {
	mt := &MyTranslator{}
	mt.defaultTranslator = &gofeed.DefaultRSSTranslator{}
	return mt
}

func (ct *MyTranslator) Translate(feed interface{}) (*gofeed.Feed, error) {
	rss, found := feed.(*rss.Feed)
	if !found {
		return nil, fmt.Errorf("Feed did not match expected type of *rss.Feed")
	}
	f, err := ct.defaultTranslator.Translate(rss)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetTwitterApi() *anaconda.TwitterApi {
	anaconda.SetConsumerKey(os.Getenv("CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("ACCESS_TOKEN"), os.Getenv("ACCESS_TOKEN_SECRET"))
	return api
}

func GetRss() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	fp := gofeed.NewParser()
	fp.RSSTranslator = NewMyTranslator()

	feed, _ := fp.ParseURL("https://www.sbbit.jp/rss/HotTopics.rss")
	api := GetTwitterApi()
	text := feed.Items[1].Title + feed.Items[1].Link

	tweet, err := api.PostTweet(text, nil)
	if err != nil {
		panic(err)
	}
}
