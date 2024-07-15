package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rs/cors"
)

// 　店舗情報の定義
type ShopInfo struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Tel     string `json:"tel"`
}

func main() {
	// APIエンドポイントとパラメータの設定
	url := "https://webservice.recruit.co.jp/hotpepper/gourmet/v1/"
	apiKey := "API Key"
	queryParams := map[string]string{
		"key":        apiKey,
		"large_area": "Z011",
		"count":      "10",
		"format":     "json",
	}

	// URLの作成
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("Failed to create request:", err)
		return
	}

	// パラメータの設定
	query := req.URL.Query()
	for key, value := range queryParams {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	// HTTPリクエストの送信
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	// 応答の確認
	if resp.StatusCode == http.StatusOK {
		var data map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			log.Println("Failed to decode JSON response:", err)
			return
		}

		// 店舗情報の抽出
		shops := extractShops(data)

		// HTTPハンドラの設定
		http.HandleFunc("/shops", func(w http.ResponseWriter, r *http.Request) {
			// JSON形式でレスポンスを返す
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			if err := json.NewEncoder(w).Encode(shops); err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
				return
			}
		})
		// CORSを有効にする
		handler := cors.Default().Handler(http.DefaultServeMux)

		//	サーバーの起動
		if err := http.ListenAndServe(":8080", handler); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	} else {
		fmt.Println("Failed to retrieve data:", resp.Status)
	}

}

// 店舗情報を抽出
func extractShops(data map[string]interface{}) []ShopInfo {
	var shops []ShopInfo

	results, ok := data["results"].(map[string]interface{})
	if !ok {
		log.Println("Results not found in JSON")
		return shops
	}

	shopsArray, ok := results["shop"].([]interface{})
	if !ok {
		log.Println("Shop array not found in JSON")
		return shops
	}

	for _, item := range shopsArray {
		shopData, ok := item.(map[string]interface{})
		if !ok {
			log.Println("Shop data format error")
			continue
		}

		shop := ShopInfo{
			Name:    fmt.Sprintf("%v", shopData["name"]),
			Address: fmt.Sprintf("%v", shopData["address"]),
			Tel:     fmt.Sprintf("%v", shopData["tel"]),
		}

		shops = append(shops, shop)
	}

	return shops
}
