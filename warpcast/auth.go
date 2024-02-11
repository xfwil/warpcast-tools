package warpcast

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetMyProfileResponse struct {
	Result struct {
		State struct {
			ID    string `json:"id"`
			Email string `json:"email"`
			User  struct {
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
			} `json:"user"`
			HasOnboarding                 bool `json:"hasOnboarding"`
			HasConfirmedEmail             bool `json:"hasConfirmedEmail"`
			HandledConnectAddress         bool `json:"handledConnectAddress"`
			CanRegisterUsername           bool `json:"canRegisterUsername"`
			NeedsRegistrationPayment      bool `json:"needsRegistrationPayment"`
			HasFid                        bool `json:"hasFid"`
			HasFname                      bool `json:"hasFname"`
			HasDelegatedSigner            bool `json:"hasDelegatedSigner"`
			HasSetupProfile               bool `json:"hasSetupProfile"`
			HasCompletedRegistration      bool `json:"hasCompletedRegistration"`
			HasStorage                    bool `json:"hasStorage"`
			HandledPushNotificationsNudge bool `json:"handledPushNotificationsNudge"`
			HandledContactsNudge          bool `json:"handledContactsNudge"`
			HandledInterestsNudge         bool `json:"handledInterestsNudge"`
			HasValidPaidInvite            bool `json:"hasValidPaidInvite"`
			HasPhone                      bool `json:"hasPhone"`
			NeedsPhone                    bool `json:"needsPhone"`
			SponsoredRegisterEligible     bool `json:"sponsoredRegisterEligible"`
		} `json:"state"`
	} `json:"result"`
}

func GetMyProfile(accessToken string) (*GetMyProfileResponse, error) {
	url := "https://client.warpcast.com/v2/onboarding-state"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return &GetMyProfileResponse{}, err
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
		return &GetMyProfileResponse{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &GetMyProfileResponse{}, err
	}
	var response GetMyProfileResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return &GetMyProfileResponse{}, err
	}

	return &response, nil
}
