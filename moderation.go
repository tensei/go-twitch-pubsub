package twitchpubsub

import (
	"encoding/json"
	"fmt"
)

// ModerationAction describes an incoming "Moderation" action coming from Twitch's PubSub servers
type ModerationAction struct {
	Type             string   `json:"type"`
	ModerationAction string   `json:"moderation_action"`
	Arguments        []string `json:"args"`
	CreatedBy        string   `json:"created_by"`
	CreatedByUserID  string   `json:"created_by_user_id"`
	MsgID            string   `json:"msg_id"`
	TargetUserID     string   `json:"target_user_id"`
}

// GetModerationAction attempts to parse a chunk of bytes into a ModerationAction structure
func GetModerationAction(bytes []byte) (*ModerationAction, error) {
	innerData, err := getInnerData(bytes)
	if err != nil {
		return nil, err
	}

	var e ModerationAction
	err = json.Unmarshal(innerData, &e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// ModerationActionTopic returns a properly formatted moderation action topic string with the given user and channel ID arguments
func ModerationActionTopic(userID, channelID string) string {
	const f = `chat_moderator_actions.%s.%s`
	return fmt.Sprintf(f, userID, channelID)
}
