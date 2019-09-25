package scm

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerialization(t *testing.T) {
	pushHookUnMarshaler := &WebhookSerializer{
		Webhook: &PushHook{
			Ref:     "789def",
			BaseRef: "123abc",
			Repo: Repository{
				Name: "test-repo",
			},
			Created: true,
			Deleted: false,
			Forced:  false,
			GUID:    "test-repo",
		},
	}

	marshaledPushHook, err := json.Marshal(pushHookUnMarshaler)
	assert.NoError(t, err)
	assert.NotNil(t, marshaledPushHook)

	var unmarshaledHook WebhookSerializer
	err = json.Unmarshal(marshaledPushHook, &unmarshaledHook)
	assert.NoError(t, err)
	assert.NotNil(t, unmarshaledHook)
	t.Logf("%v", unmarshaledHook)

	unmarshaledPushHook, ok := unmarshaledHook.Webhook.(*PushHook)
	assert.True(t, ok)
	assert.NotNil(t, unmarshaledPushHook)

}
