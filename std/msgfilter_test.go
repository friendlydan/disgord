package std

import (
	"testing"

	"github.com/andersfylling/disgord"
)

type clientMock struct {
	id disgord.Snowflake
}

var _ msgFilterdg = (*clientMock)(nil)

func (c *clientMock) Myself() (*disgord.User, error) {
	return &disgord.User{ID: c.id}, nil
}

func TestNewMsgFilter(t *testing.T) {
	var botID disgord.Snowflake = 123
	filter, err := NewMsgFilter(&clientMock{botID})
	if err != nil {
		t.Fatal(err)
	}

	if filter.botID != botID {
		t.Errorf("expected filter to have the same id as client. Got %d, wants %d", filter.botID, botID)
	}
}

func TestMsgFilter_ContainsBotMention(t *testing.T) {
	var botID disgord.Snowflake = 123
	filter, _ := NewMsgFilter(&clientMock{botID})
	var evt interface{}
	e := &disgord.MessageCreate{
		Message: &disgord.Message{Content: "<@" + botID.String() + "> hello"},
	}
	evt = e

	result := filter.ContainsBotMention(evt)
	if result == nil {
		t.Error("expected to find a match")
	}

	e.Message.Content = "diff prefix " + e.Message.Content
	result = filter.ContainsBotMention(evt)
	if result == nil {
		t.Error("expected to find a match")
	}

	filter.botID = botID + 3
	result = filter.ContainsBotMention(evt)
	if result != nil {
		t.Error("did not expect a match")
	}
}

func TestMsgFilter_HasBotMentionPrefix(t *testing.T) {
	var botID disgord.Snowflake = 123
	filter, _ := NewMsgFilter(&clientMock{botID})
	var evt interface{}
	e := &disgord.MessageCreate{
		Message: &disgord.Message{Content: "<@" + botID.String() + "> hello"},
	}
	evt = e

	result := filter.HasBotMentionPrefix(evt)
	if result == nil {
		t.Error("expected to find a match")
	}

	e.Message.Content = "diff prefix " + e.Message.Content
	result = filter.HasBotMentionPrefix(evt)
	if result != nil {
		t.Error("did not expect a match")
	}
}

func TestMsgFilter_SetPrefix(t *testing.T) {
	filter, _ := NewMsgFilter(&clientMock{})
	if filter.prefix != "" {
		t.Fatal("expected prefix to be empty")
	}

	filter.SetPrefix("!")
	if filter.prefix != "!" {
		t.Errorf("wrong prefix. Got %s, wants %s", filter.prefix, "!")
	}
}

func TestMsgFilter_HasPrefix(t *testing.T) {
	prefix := "!!"
	filter, _ := NewMsgFilter(&clientMock{})
	filter.SetPrefix(prefix)

	var evt interface{}
	e := &disgord.MessageCreate{
		Message: &disgord.Message{Content: prefix + "hello"},
	}
	evt = e

	result := filter.HasPrefix(evt)
	if result == nil {
		t.Error("expected to find a match")
	}

	e.Message.Content = "diff prefix " + e.Message.Content
	result = filter.HasBotMentionPrefix(evt)
	if result != nil {
		t.Error("did not expect a match")
	}
}

func TestMsgFilter_StripPrefix(t *testing.T) {
	prefix := "!!"
	filter, _ := NewMsgFilter(&clientMock{})
	filter.SetPrefix(prefix)

	var evt interface{}
	e := &disgord.MessageCreate{
		Message: &disgord.Message{Content: prefix + "hello"},
	}
	evt = e

	result := filter.StripPrefix(evt)
	if result == nil {
		t.Error("expected prefix stripping to work")
	}
	if filter.HasPrefix(evt) != nil {
		t.Error("did not strip prefix off message")
	}
}
