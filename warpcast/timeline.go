package warpcast

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type LikeResponse struct {
	Result struct {
		Like struct {
			Type    string `json:"type"`
			Hash    string `json:"hash"`
			Reactor struct {
				Fid         int    `json:"fid"`
				Username    string `json:"username"`
				DisplayName string `json:"displayName"`
				Pfp         struct {
					URL      string `json:"url"`
					Verified bool   `json:"verified"`
				} `json:"pfp"`
				Profile struct {
					Bio struct {
						Text            string        `json:"text"`
						Mentions        []interface{} `json:"mentions"`
						ChannelMentions []interface{} `json:"channelMentions"`
					} `json:"bio"`
					Location struct {
						PlaceID     string `json:"placeId"`
						Description string `json:"description"`
					} `json:"location"`
				} `json:"profile"`
				FollowerCount     int  `json:"followerCount"`
				FollowingCount    int  `json:"followingCount"`
				ActiveOnFcNetwork bool `json:"activeOnFcNetwork"`
				ViewerContext     struct {
					Following  bool `json:"following"`
					FollowedBy bool `json:"followedBy"`
				} `json:"viewerContext"`
			} `json:"reactor"`
			Timestamp int64  `json:"timestamp"`
			CastHash  string `json:"castHash"`
		} `json:"like"`
	} `json:"result"`
}

func Like(accessToken string, castHash string) (*LikeResponse, error) {
	url := "https://client.warpcast.com/v2/cast-likes"
	method := "PUT"

	payload := strings.NewReader(`{"castHash":"` + castHash + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return &LikeResponse{}, err
	}
	req.Header.Add("authority", "client.warpcast.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("authorization", "Bearer "+accessToken)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("origin", "https://warpcast.com")
	req.Header.Add("referer", "https://warpcast.com/")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return &LikeResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &LikeResponse{}, err
	}

	statusCode := res.StatusCode

	if statusCode != 200 {
		return &LikeResponse{}, err
	}

	var likeResponse LikeResponse
	err = json.Unmarshal(body, &likeResponse)
	if err != nil {
		return &LikeResponse{}, err
	}

	return &likeResponse, nil
}

type RecastResponse struct {
	Result struct {
		CastHash string `json:"castHash"`
	} `json:"result"`
}

func Recast(accessToken string, castHash string) (*RecastResponse, error) {
	url := "https://client.warpcast.com/v2/recasts"
	method := "PUT"

	payload := strings.NewReader(`{"castHash":"` + castHash + `"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return &RecastResponse{}, err
	}
	req.Header.Add("authority", "client.warpcast.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("authorization", "Bearer "+accessToken)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("origin", "https://warpcast.com")
	req.Header.Add("referer", "https://warpcast.com/")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return &RecastResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &RecastResponse{}, err
	}

	statusCode := res.StatusCode

	if statusCode != 200 {
		return &RecastResponse{}, err
	}

	var recastResponse RecastResponse
	err = json.Unmarshal(body, &recastResponse)

	if err != nil {
		return &RecastResponse{}, err
	}

	return &recastResponse, nil
}

type CommentResponse struct {
	Result struct {
		Cast struct {
			Hash         string `json:"hash"`
			ThreadHash   string `json:"threadHash"`
			ParentHash   string `json:"parentHash"`
			ParentAuthor struct {
				Fid         int    `json:"fid"`
				Username    string `json:"username"`
				DisplayName string `json:"displayName"`
				Pfp         struct {
					URL      string `json:"url"`
					Verified bool   `json:"verified"`
				} `json:"pfp"`
				Profile struct {
					Bio struct {
						Text            string        `json:"text"`
						Mentions        []interface{} `json:"mentions"`
						ChannelMentions []interface{} `json:"channelMentions"`
					} `json:"bio"`
					Location struct {
						PlaceID     string `json:"placeId"`
						Description string `json:"description"`
					} `json:"location"`
				} `json:"profile"`
				FollowerCount     int  `json:"followerCount"`
				FollowingCount    int  `json:"followingCount"`
				ActiveOnFcNetwork bool `json:"activeOnFcNetwork"`
			} `json:"parentAuthor"`
			Author struct {
				Fid         int    `json:"fid"`
				Username    string `json:"username"`
				DisplayName string `json:"displayName"`
				Pfp         struct {
					URL      string `json:"url"`
					Verified bool   `json:"verified"`
				} `json:"pfp"`
				Profile struct {
					Bio struct {
						Text            string        `json:"text"`
						Mentions        []interface{} `json:"mentions"`
						ChannelMentions []interface{} `json:"channelMentions"`
					} `json:"bio"`
					Location struct {
						PlaceID     string `json:"placeId"`
						Description string `json:"description"`
					} `json:"location"`
				} `json:"profile"`
				FollowerCount     int  `json:"followerCount"`
				FollowingCount    int  `json:"followingCount"`
				ActiveOnFcNetwork bool `json:"activeOnFcNetwork"`
				ViewerContext     struct {
					Following bool `json:"following"`
				} `json:"viewerContext"`
			} `json:"author"`
			Text      string `json:"text"`
			Timestamp int64  `json:"timestamp"`
			Embeds    struct {
				Images   []interface{} `json:"images"`
				Urls     []interface{} `json:"urls"`
				Videos   []interface{} `json:"videos"`
				Unknowns []interface{} `json:"unknowns"`
			} `json:"embeds"`
			Replies struct {
				Count int `json:"count"`
			} `json:"replies"`
			Reactions struct {
				Count int `json:"count"`
			} `json:"reactions"`
			Recasts struct {
				Count     int           `json:"count"`
				Recasters []interface{} `json:"recasters"`
			} `json:"recasts"`
			Watches struct {
				Count int `json:"count"`
			} `json:"watches"`
			Tags []struct {
				Type     string `json:"type"`
				ID       string `json:"id"`
				Name     string `json:"name"`
				ImageURL string `json:"imageUrl"`
			} `json:"tags"`
			QuoteCount          int `json:"quoteCount"`
			CombinedRecastCount int `json:"combinedRecastCount"`
			ViewerContext       struct {
				Reacted bool `json:"reacted"`
				Recast  bool `json:"recast"`
				Watched bool `json:"watched"`
			} `json:"viewerContext"`
		} `json:"cast"`
	} `json:"result"`
}

func Comment(accessToken string, castHash string, text string) (*CommentResponse, error) {
	url := "https://client.warpcast.com/v2/casts"
	method := "POST"

	payload := strings.NewReader(`{"text":"` + text + `","parent":{"hash":"` + castHash + `"},"embeds":[]}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return &CommentResponse{}, err
	}
	req.Header.Add("authority", "client.warpcast.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("authorization", "Bearer "+accessToken)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("origin", "https://warpcast.com")
	req.Header.Add("referer", "https://warpcast.com/")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return &CommentResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &CommentResponse{}, err
	}

	statusCode := res.StatusCode

	if statusCode != 201 {
		return &CommentResponse{}, err
	}

	var commentResponse CommentResponse
	err = json.Unmarshal(body, &commentResponse)
	if err != nil {
		return &CommentResponse{}, err
	}

	return &commentResponse, nil
}

type GetFeedsItemsResponse struct {
	Result struct {
		Items []struct {
			ID        string `json:"id"`
			Timestamp int64  `json:"timestamp"`
			Cast      struct {
				Hash       string `json:"hash"`
				ThreadHash string `json:"threadHash"`
				Author     struct {
					Fid         int    `json:"fid"`
					Username    string `json:"username"`
					DisplayName string `json:"displayName"`
					Pfp         struct {
						URL      string `json:"url"`
						Verified bool   `json:"verified"`
					} `json:"pfp"`
					Profile struct {
						Bio struct {
							Text            string        `json:"text"`
							Mentions        []interface{} `json:"mentions"`
							ChannelMentions []interface{} `json:"channelMentions"`
						} `json:"bio"`
						Location struct {
							PlaceID     string `json:"placeId"`
							Description string `json:"description"`
						} `json:"location"`
					} `json:"profile"`
					FollowerCount     int  `json:"followerCount"`
					FollowingCount    int  `json:"followingCount"`
					ActiveOnFcNetwork bool `json:"activeOnFcNetwork"`
					ViewerContext     struct {
						Following bool `json:"following"`
					} `json:"viewerContext"`
				} `json:"author"`
				Text        string        `json:"text"`
				Timestamp   int64         `json:"timestamp"`
				Mentions    []interface{} `json:"mentions"`
				Attachments struct {
				} `json:"attachments"`
				Replies struct {
					Count int `json:"count"`
				} `json:"replies"`
				Reactions struct {
					Count int `json:"count"`
				} `json:"reactions"`
				Recasts struct {
					Count     int           `json:"count"`
					Recasters []interface{} `json:"recasters"`
				} `json:"recasts"`
				Watches struct {
					Count int `json:"count"`
				} `json:"watches"`
				Tags []struct {
					Type     string `json:"type"`
					ID       string `json:"id"`
					Name     string `json:"name"`
					ImageURL string `json:"imageUrl"`
				} `json:"tags"`
				QuoteCount          int `json:"quoteCount"`
				CombinedRecastCount int `json:"combinedRecastCount"`
				ViewerContext       struct {
					Reacted bool `json:"reacted"`
					Recast  bool `json:"recast"`
					Watched bool `json:"watched"`
				} `json:"viewerContext"`
			} `json:"cast"`
			OtherParticipants []interface{} `json:"otherParticipants"`
		} `json:"items"`
		LatestMainCastTimestamp int64 `json:"latestMainCastTimestamp"`
		FeedTopSeenAtTimestamp  int64 `json:"feedTopSeenAtTimestamp"`
	} `json:"result"`
}

func GetFeedsItems(accessToken string, feedKey string, lastTimestamp int64, excludeHash []string) (*GetFeedsItemsResponse, error) {
	url := "https://client.warpcast.com/v2/feed-items"
	method := "POST"

	payload := &strings.Reader{}
	if lastTimestamp == 0 {
		payload = strings.NewReader(`{"feedKey":"` + feedKey + `","viewedCastHashes":"","updateState":true}`)
	} else if lastTimestamp != 0 && len(excludeHash) > 0 {
		excludeHashJSON, err := json.Marshal(excludeHash)
		if err != nil {
			return &GetFeedsItemsResponse{}, err
		}

		payload = strings.NewReader(`{"feedKey":"` + feedKey + `","viewedCastHashes":"","updateState":true, "olderThan":` + strconv.FormatInt(lastTimestamp, 10) + `, "excludeItemIdPrefixes": ` + string(excludeHashJSON) + `}`)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return &GetFeedsItemsResponse{}, err
	}
	req.Header.Add("authority", "client.warpcast.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "id-ID,id;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("authorization", "Bearer "+accessToken)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("origin", "https://warpcast.com")
	req.Header.Add("referer", "https://warpcast.com/")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return &GetFeedsItemsResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &GetFeedsItemsResponse{}, err
	}

	var getFeedsItems GetFeedsItemsResponse
	err = json.Unmarshal(body, &getFeedsItems)
	if err != nil {
		return &GetFeedsItemsResponse{}, err
	}

	return &getFeedsItems, nil
}
