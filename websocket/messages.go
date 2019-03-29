package websocket

import (
	"math/big"
	"time"
)

// EventType - data structure for parsing events
type EventType struct {
	Event string `json:"event"`
}

// Subscription - data structure of subscription entity
type Subscription struct {
	Name     string `json:"name"`
	Interval int64  `json:"interval,omitempty"`
	Depth    int64  `json:"depth,omitempty"`
}

// SubscriptionRequest - data structure for subscription request
type SubscriptionRequest struct {
	ReqID string `json:"reqid,omitempty"`
	Event string `json:"event"`

	Pairs        []string     `json:"pair"`
	Subscription Subscription `json:"subscription"`
}

// UnsubscribeRequest - data structure for unsubscription request
type UnsubscribeRequest struct {
	Event        string       `json:"event"`
	Pairs        []string     `json:"pair"`
	Subscription Subscription `json:"subscription"`
}

// SubscriptionStatus - data structure for subscription status event
type SubscriptionStatus struct {
	ChannelID    int64        `json:"channelID"`
	Event        string       `json:"event"`
	Status       string       `json:"status"`
	Pair         string       `json:"pair"`
	ReqID        string       `json:"reqid,omitempty"`
	Error        string       `json:"errorMessage,omitempty"`
	Subscription Subscription `json:"subscription"`
}

// PingRequest - data structure for ping request
type PingRequest struct {
	Event string `json:"event"`
	ReqID int    `json:"reqid,omitempty"`
}

// PongResponse - data structure for ping response
type PongResponse struct {
	Event string `json:"event"`
	ReqID int    `json:"reqid,omitempty"`
}

// SystemStatus - data structure for system status event
type SystemStatus struct {
	Event        string  `json:"event"`
	ConnectionID big.Int `json:"connectionID"`
	Status       string  `json:"status"`
	Version      string  `json:"version"`
}

// TickerUpdate - data structure for ticker update
type TickerUpdate struct {
	Ask                Level
	Bid                Level
	Close              Values
	Volume             Values
	VolumeAveragePrice Values
	TradeVolume        Values
	Low                Values
	High               Values
	Open               Values
	Pair               string
}

// Level - data structure for ticker data ask/bid
type Level struct {
	Price          float64
	Volume         float64
	WholeLotVolume int
}

// Values - data structure for ticker others data
type Values struct {
	Today  interface{}
	Last24 interface{}
}

// CandleUpdate - data structure for candles update
type CandleUpdate struct {
	Time      time.Time
	EndTime   time.Time
	Open      float64
	High      float64
	Low       float64
	Close     float64
	VolumeWAP float64
	Volume    float64
	Count     int
	Pair      string
}

// TradeUpdate - data structure for trade update
type TradeUpdate struct {
	Price     float64
	Volume    float64
	Time      time.Time
	Side      string
	OrderType string
	Misc      string
	Pair      string
}

// SpreadUpdate - data structure for spread update
type SpreadUpdate struct {
	Ask  float64
	Bid  float64
	Time time.Time
	Pair string
}

// OrderBookItem - data structure for order book item
type OrderBookItem struct {
	Price  float64
	Volume float64
	Time   time.Time
}

// OrderBookUpdate - data structure for order book update
type OrderBookUpdate struct {
	Asks       []OrderBookItem
	Bids       []OrderBookItem
	IsSnapshot bool
	Pair       string
}
