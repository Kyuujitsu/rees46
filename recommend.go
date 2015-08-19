package rees46

import (
	"fmt"
	"net/url"

	qs "github.com/google/go-querystring/query"
)

type RecommendService struct {
	client *Client
}

func (r *RecommendService) Popular(req *PopularRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("popular", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) Similar(req *SimilarRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("similar", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) AlsoBought(req *AlsoBoughtRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("also_bought", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) SeeAlso(req *SeeAlsoRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("see_also", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) Interesting(req *InterestingRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("interesting", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) RecentlyViewed(req *RecentlyViewedRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("recently_viewed", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) BuyingNow(req *BuyingNowRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("buying_now", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) Dating(req *DatingRequest) (*RecommendationResult, *Response, error) {
	items := new([]Item)
	r.setSession(req.SessionID)
	resp, err := r.get("dating", req, items)
	result := &RecommendationResult{SessionID: r.client.SessionID, Items: *items}
	return result, resp, err
}

func (r *RecommendService) get(recommendationType string, opt interface{}, result interface{}) (*Response, error) {
	params, err := qs.Values(opt)
	if err != nil {
		return nil, err
	}

	if r.client.SessionID != "" {
		params.Add("ssid", r.client.SessionID)
	}

	params.Add("shop_id", r.client.ShopID)
	params.Add("recommender_type", recommendationType)
	u := fmt.Sprintf("recommend?%s", params.Encode())

	req, err := r.client.NewRequest(GET_METHOD, u, nil)
	if err != nil {
		return nil, err
	}

	return r.client.Do(req, result)
}

func (r *RecommendService) generateSession() (string, error) {
	params := url.Values{}
	params.Add("shop_id", r.client.ShopID)
	u := fmt.Sprintf("generate_ssid?%s", params.Encode())

	req, err := r.client.NewRequest(GET_METHOD, u, nil)
	sessionResponse, err := r.client.Do(req, nil)
	return sessionResponse.RawBody, err
}

func (r *RecommendService) setSession(sessionId string) (err error) {
	if sessionId != "" {
		r.client.SessionID = sessionId
	} else {
		r.client.SessionID, err = r.generateSession()
	}
	return
}
