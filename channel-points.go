package twitchpubsub

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const channelPointsTopicPrefix = "channel-points-channel-v1."

// ChannelPoints describes an incoming "ChannelPoints" redemption coming from Twitch's PubSub servers
type ChannelPoints struct {
	Redemption struct {
		ChannelID  string `json:"channel_id"`
		ID         string `json:"id"`
		RedeemedAt string `json:"redeemed_at"`
		Reward     struct {
			BackgroundColor string `json:"background_color"`
			ChannelID       string `json:"channel_id"`
			Cost            int64  `json:"cost"`
			DefaultImage    struct {
				URL1x string `json:"url_1x"`
				URL2x string `json:"url_2x"`
				URL4x string `json:"url_4x"`
			} `json:"default_image"`
			ID    string `json:"id"`
			Image struct {
				URL1x string `json:"url_1x"`
				URL2x string `json:"url_2x"`
				URL4x string `json:"url_4x"`
			} `json:"image"`
			IsEnabled           bool `json:"is_enabled"`
			IsInStock           bool `json:"is_in_stock"`
			IsPaused            bool `json:"is_paused"`
			IsSubOnly           bool `json:"is_sub_only"`
			IsUserInputRequired bool `json:"is_user_input_required"`
			MaxPerStream        struct {
				IsEnabled    bool  `json:"is_enabled"`
				MaxPerStream int64 `json:"max_per_stream"`
			} `json:"max_per_stream"`
			Prompt                            string `json:"prompt"`
			ShouldRedemptionsSkipRequestQueue bool   `json:"should_redemptions_skip_request_queue"`
			Title                             string `json:"title"`
		} `json:"reward"`
		Status string `json:"status"`
		User   struct {
			DisplayName string `json:"display_name"`
			ID          string `json:"id"`
			Login       string `json:"login"`
		} `json:"user"`
		UserInput string `json:"user_input"`
	} `json:"redemption"`
	Timestamp string `json:"timestamp"`
}

type outerChannelPoints struct {
	Data ChannelPoints `json:"data"`
}

func parseChannelPoints(bytes []byte) (*ChannelPoints, error) {
	data := &outerChannelPoints{}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		return nil, err
	}

	return &data.Data, nil
}

func parseChannelIDFromChannelPointsTopic(topic string) (string, error) {
	parts := strings.Split(topic, ".")
	if len(parts) != 2 {
		return "", errors.New("Unable to parse channel ID from channel points topic")
	}

	return parts[1], nil
}

// ChannelPointsTopic returns a properly formatted channel points topic string with the given user and channel ID argument
func ChannelPointsTopic(channelID string) string {
	const f = `channel-points-channel-v1.%s`
	return fmt.Sprintf(f, channelID)
}
