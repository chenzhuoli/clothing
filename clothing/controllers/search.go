package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
	"io"
	"github.com/astaxie/beego"
	//"clothing/models"
	"fmt"
)

type SearchController struct {
	beego.Controller
}

type SearchResponse struct {
	Code int `json:"code"`
	Data []HitDoc `json:"data"`
	Msg string `json:"msg"`
}
/*
{
  "query": {
    "multi_match": {
      "query": "连衣裙",
      "fields": ["title", "style", "designer", "colour"]
    }
  }
}
*/

type ElasticSearchRequest struct {
	Query ElasticSearchQuery `json:"query"`
}

type ElasticSearchQuery struct {
	MultiMatch ElasticSearchMultiMatch `json:"multi_match"`
}

type ElasticSearchMultiMatch struct {
	MultiQuery string `json:"query"`
	Fields []string `json:"fields"`
}

/*
{
    "took": 1,
    "timed_out": false,
    "_shards": {
        "total": 1,
        "successful": 1,
        "skipped": 0,
        "failed": 0
    },
    "hits": {
        "total": {
            "value": 2,
            "relation": "eq"
        },
        "max_score": 1.3112575,
        "hits": [
            {
                "_index": "clothing",
                "_type": "_doc",
                "_id": "5",
                "_score": 1.3112575,
                "_source": {
                    "title": "红色连衣裙",
                    "style": "连衣裙",
                    "designer": "小丽",
                    "colour": "红色",
                    "picture": "http://101.42.109.110:8081/6,047c4188ab"
                }
            },
            {
                "_index": "clothing",
                "_type": "_doc",
                "_id": "2",
                "_score": 0.6931471,
                "_source": {
                    "title": "桔红色裙装",
                    "style": "连衣裙",
                    "designer": "小丽",
                    "colour": "桔红色",
                    "picture": "http://101.42.109.110:8081/7,02d1134eb8"
                }
            }
        ]
    }
}
*/

type ElasticSearchResponse struct {
	Took int `json:"took"`
	TimeOut bool `json:"timed_out"`
	Shards ElasticSearchShards `json:"_shards"`
	Hits TotalHits `json:"hits"`
}

type ElasticSearchShards struct {
	Total int `json:"total"`
	Successful int `json:"successful"`
	Skipped int `json:"skipped"`
	Failed int `json:"failed"`
}

type TotalHits struct {
	Total TotalAttribute `json:"total"`	
	MaxScore float32 `json:"max_score"`
	Hits []HitDoc `json:"hits"`
}

type TotalAttribute struct {
	Value int `json:"value"`
	Relation string `json:"relation"`
}

type HitDoc struct {
	Index string `json:"_index"`
	Type string `json:"_type"`
	ID string `json:"_id"`
	Score float32 `json:"_score"`
	Source SourceDoc `json:"_source"`
}

type SourceDoc struct {
	Title string `json:"title"`
	Style string `json:"style"`
	Designer string `json:"designer"`
	Colour string `json:"colour"`
	Picture string `json:"picture"`
}

func (c *SearchController) GetSearchList() {
        c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")  //允许跨域
	// URL是 /search?query=连衣裙
	query := c.GetString("query")
	fmt.Println("Received: query=",query)
	fmt.Println(query)
	esResp, err := GetElasticSearchResult(query)
	if err != nil {
		resp := SearchResponse{
			Code: 500,
			Data: nil,
			Msg: err.Error(),
		}
		c.Data["json"]  = resp
		c.ServeJSON()
		return
	}
	resp := SearchResponse{
		Code: 200,
		Data: esResp.Hits.Hits,
		Msg: "OK",
	}
	c.Data["json"]  = resp
	c.ServeJSON()
}

func GetElasticSearchResult(query string) (*ElasticSearchResponse, error) {
	req_str := "http://101.42.109.110:9200/clothing/_search"
	req := ElasticSearchRequest{
		Query: ElasticSearchQuery{
			MultiMatch: ElasticSearchMultiMatch{
				MultiQuery: query,
				Fields: []string{"title", "style", "designer", "colour"},
			},
		},
	}
	dataByte, err := json.Marshal(req)
	if err !=nil {
		fmt.Println(err)
		return nil, err
	}
	bodyReader := bytes.NewReader(dataByte)
    	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
    	defer cancelFunc()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, req_str, bodyReader)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	defer resp.Body.Close()
	var elasticSearchResp ElasticSearchResponse
	err = json.Unmarshal(body, &elasticSearchResp)
	return &elasticSearchResp, err
}
