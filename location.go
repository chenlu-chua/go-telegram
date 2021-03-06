package telegram

import (
	"context"
)

type Location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	LivePeriod           int     `json:"live_period"`
	Heading              int     `json:"heading"`
	ProximityAlertRadius int     `json:"proximity_alert_radius"`
}

type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareId    string    `json:"foursquare_id"`
	FoursquareType  string    `json:"foursquare_type"`
	GooglePlaceId   string    `json:"google_place_id"`
	GooglePlaceType string    `json:"google_place_type"`
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

func (c *BotClient) SendLocation(ctx context.Context, options SendLocationOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiSendLocation, options, &message)
	return &message, err
}

type SendLocationOptions struct {
	ChatId                   *int        `json:"chat_id,omitempty"`
	Latitude                 *float64    `json:"latitude,omitempty"`
	Longitude                *float64    `json:"longitude,omitempty"`
	HorizontalAccuracy       *float64    `json:"horizontal_accuracy,omitempty"`
	LivePeriod               *int        `json:"live_period,omitempty"`
	Heading                  *int        `json:"heading,omitempty"`
	ProximityAlertRadius     *int        `json:"proximity_alert_radius,omitempty"`
	DisableNotification      *bool       `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int        `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool       `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              interface{} `json:"reply_markup,omitempty"`
}

func (c *BotClient) EditMessageLiveLocation(ctx context.Context, options EditMessageLiveLocationOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiEditMessageLiveLocation, options, &message)
	return &message, err
}

type EditMessageLiveLocationOptions struct {
	ChatId               *int                         `json:"chat_id,omitempty"`
	MessageId            *int                         `json:"message_id,omitempty"`
	InlineMessageId      *string                      `json:"inline_message_id,omitempty"`
	Latitude             *float64                     `json:"latitude,omitempty"`
	Longitude            *float64                     `json:"longitude,omitempty"`
	HorizontalAccuracy   *float64                     `json:"horizontal_accuracy,omitempty"`
	Heading              *int                         `json:"heading,omitempty"`
	ProximityAlertRadius *int                         `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}

func (c *BotClient) StopMessageLiveLocation(ctx context.Context, options StopMessageLiveLocationOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiStopMessageLiveLocation, options, &message)
	return &message, err
}

type StopMessageLiveLocationOptions struct {
	ChatId          *int                         `json:"chat_id,omitempty"`
	MessageId       *int                         `json:"message_id,omitempty"`
	InlineMessageId *string                      `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}

func (c *BotClient) SendVenue(ctx context.Context, options SendVenueOptions) (*Message, error) {
	var message Message
	err := c.postJson(ctx, apiSendVenue, options, &message)
	return &message, err
}

type SendVenueOptions struct {
	ChatId                   *int                         `json:"chat_id,omitempty"`
	Latitude                 *float64                     `json:"latitude,omitempty"`
	Longitude                *float64                     `json:"longitude,omitempty"`
	Title                    *string                      `json:"title,omitempty"`
	Address                  *string                      `json:"address,omitempty"`
	FoursquareId             *string                      `json:"foursquare_id,omitempty"`
	FoursquareType           *string                      `json:"foursquare_type,omitempty"`
	GooglePlaceId            *string                      `json:"google_place_id,omitempty"`
	GooglePlaceType          *string                      `json:"google_place_type,omitempty"`
	DisableNotification      *bool                        `json:"disable_notification,omitempty"`
	ReplyToMessageId         *int                         `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply *bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup              *InlineKeyboardMarkupOptions `json:"reply_markup,omitempty"`
}
