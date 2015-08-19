package rees46

type Item string

type RecommendationResult struct {
	SessionID string `json:"ssid"`
	Items     []Item `json:"items"`
}

type PopularRequest struct {
	SessionID        string            `json:"ssid" url:"-"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}

type SimilarRequest struct {
	SessionID        string `json:"ssid" url:"-"`
	Item             `json:"item_id" url:"item_id"`
	CartItems        []Item            `json:"cart_item_id" url:"cart_item_id[]"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Categories       string            `json:"categories,omitempty" url:"categories,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}

type AlsoBoughtRequest struct {
	SessionID        string `json:"ssid" url:"-"`
	Item             `json:"item_id" url:"item_id"`
	CartItems        []Item            `json:"cart_item_id" url:"cart_item_id[]"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Categories       string            `json:"categories,omitempty" url:"categories,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}

type SeeAlsoRequest struct {
	SessionID        string            `json:"ssid" url:"-"`
	CartItems        []Item            `json:"cart_item_id" url:"cart_item_id[]"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}

type InterestingRequest struct {
	SessionID        string `json:"ssid" url:"-"`
	Item             `json:"item_id,omitempty" url:"item_id,omitempty"`
	CartItems        []Item            `json:"cart_item_id,omitempty" url:"cart_item_id[],omitempty"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}

type RecentlyViewedRequest struct {
	SessionID        string            `json:"ssid" url:"-"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}

type BuyingNowRequest struct {
	SessionID        string `json:"ssid" url:"-"`
	Item             `json:"item_id,omitempty" url:"item_id,omitempty"`
	CartItems        []Item            `json:"cart_item_id,omitempty" url:"cart_item_id[],omitempty"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}

type DatingRequest struct {
	SessionID        string            `json:"ssid" url:"-"`
	Category         string            `json:"category,omitempty" url:"category,omitempty"`
	Limit            int               `json:"limit,omitempty" url:"limit,omitempty"`
	Locations        string            `json:"locations,omitempty" url:"locations,omitempty"`
	Items            string            `json:"items,omitempty" url:"items,omitempty"`
	Exclude          string            `json:"exclude,omitempty" url:"exclude,omitempty"`
	Modification     string            `json:"modification,omitempty" url:"modification,omitempty"`
	CustomAttributes map[string]string `json:"custom_attributes,omitempty" url:"custom_attributes,omitempty"`
}
