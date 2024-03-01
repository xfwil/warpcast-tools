package degen

import (
	"encoding/json"
	"net/http"
	"time"
)

type GetPointsResponse struct {
	DisplayName string `json:"display_name"`
	Points      string `json:"points"`
}

func GetPoints(address string) ([]GetPointsResponse, error) {
	url := "https://www.degen.tips/api/airdrop2/season2/points?address=" + address
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Referer", "https://warpcast.com/")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response []GetPointsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

type GetTipAllowanceResponse struct {
	SnapshotDate       time.Time `json:"snapshot_date"`
	UserRank           string    `json:"user_rank"`
	WalletAddress      string    `json:"wallet_address"`
	AvatarURL          string    `json:"avatar_url"`
	DisplayName        string    `json:"display_name"`
	TipAllowance       string    `json:"tip_allowance"`
	RemainingAllowance string    `json:"remaining_allowance"`
}

func GetTipAllowance(address string) ([]GetTipAllowanceResponse, error) {
	url := "https://www.degen.tips/api/airdrop2/tip-allowance?address=" + address
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Referer", "https://warpcast.com/")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response []GetTipAllowanceResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
