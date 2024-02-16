package warpcast

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type GetProfileResponse struct {
	Result struct {
		User struct {
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
				Following            bool `json:"following"`
				FollowedBy           bool `json:"followedBy"`
				CanSendDirectCasts   bool `json:"canSendDirectCasts"`
				HasUploadedInboxKeys bool `json:"hasUploadedInboxKeys"`
			} `json:"viewerContext"`
		} `json:"user"`
		InviterIsReferrer bool          `json:"inviterIsReferrer"`
		CollectionsOwned  []interface{} `json:"collectionsOwned"`
		Extras            struct {
			Fid            int    `json:"fid"`
			CustodyAddress string `json:"custodyAddress"`
		} `json:"extras"`
	} `json:"result"`
}

func GetProfile(accessToken string, username string) (*GetProfileResponse, error) {
	url := "https://client.warpcast.com/v2/user-by-username?username=" + username
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return &GetProfileResponse{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Referer", "https://warpcast.com/")

	res, err := client.Do(req)
	if err != nil {
		return &GetProfileResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &GetProfileResponse{}, err
	}

	var profile GetProfileResponse
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return &GetProfileResponse{}, err
	}

	return &profile, nil
}

type GetFollowersResponse struct {
	Result struct {
		Users []struct {
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
		} `json:"users"`
	} `json:"result"`
	Next struct {
		Cursor string `json:"cursor"`
	} `json:"next"`
}

func GetProfileInformation(types string, accessToken string, fid string, cursor string) (*GetFollowersResponse, error) {
	url := "https://client.warpcast.com/v2/" + types + "?fid=" + fid + "&limit=200&cursor=" + cursor
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return &GetFollowersResponse{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Referer", "https://warpcast.com/")

	res, err := client.Do(req)
	if err != nil {
		return &GetFollowersResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &GetFollowersResponse{}, err
	}

	var followers GetFollowersResponse
	err = json.Unmarshal(body, &followers)

	if err != nil {
		return &GetFollowersResponse{}, err
	}

	return &followers, nil
}

type FollowResponse struct {
	Result struct {
		Success        bool `json:"success"`
		UserAppContext struct {
			CanAddLinks bool `json:"canAddLinks"`
		} `json:"userAppContext"`
	} `json:"result"`
}

func Follow(accessToken string, fid string) (*FollowResponse, error) {
	url := "https://client.warpcast.com/v2/follows"
	method := "PUT"

	payload := strings.NewReader(`{"targetFid":` + fid + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return &FollowResponse{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Referer", "https://warpcast.com/")

	res, err := client.Do(req)
	if err != nil {
		return &FollowResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &FollowResponse{}, err
	}

	statusCode := res.StatusCode
	if statusCode != 200 {
		return &FollowResponse{}, errors.New("error")
	}

	var follow FollowResponse
	err = json.Unmarshal(body, &follow)
	if err != nil {
		return &FollowResponse{}, err
	}

	return &follow, nil
}

func Unfollow(accessToken string, fid string) (*FollowResponse, error) {
	url := "https://client.warpcast.com/v2/follows"
	method := "DELETE"

	payload := strings.NewReader(`{"targetFid":` + fid + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return &FollowResponse{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Referer", "https://warpcast.com/")

	res, err := client.Do(req)
	if err != nil {
		return &FollowResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &FollowResponse{}, err
	}

	statusCode := res.StatusCode
	if statusCode != 200 {
		return &FollowResponse{}, errors.New("error")
	}

	var follow FollowResponse
	err = json.Unmarshal(body, &follow)
	if err != nil {
		return &FollowResponse{}, err
	}

	return &follow, nil
}

type GetAddressVerifiedResponse struct {
	Result struct {
		Verifications []struct {
			Fid       int    `json:"fid"`
			Address   string `json:"address"`
			Timestamp int64  `json:"timestamp"`
			Version   string `json:"version"`
		} `json:"verifications"`
	} `json:"result"`
}

func GetAddressVerified(accessToken string, fid string) (*GetAddressVerifiedResponse, error) {
	url := "https://client.warpcast.com/v2/verifications?fid=" + fid + "&limit=5"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return &GetAddressVerifiedResponse{}, err
	}
	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Referer", "https://warpcast.com/")

	res, err := client.Do(req)
	if err != nil {
		return &GetAddressVerifiedResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &GetAddressVerifiedResponse{}, err
	}

	var profile GetAddressVerifiedResponse
	err = json.Unmarshal(body, &profile)
	if err != nil {
		return &GetAddressVerifiedResponse{}, err
	}

	return &profile, nil
}
