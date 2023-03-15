package retag

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetJSONTag(t *testing.T) {
	type data struct {
		Username string `json:""`
	}
	v := data{
		Username: `ok`,
	}
	r := SetJSONTag(&v, `Username`, `UserName`)
	b, err := json.Marshal(r)
	assert.NoError(t, err)
	assert.Equal(t, `{"UserName":"ok"}`, string(b))

	values := []data{
		{Username: `ok`},
	}
	r = SetJSONTag(&values, `Username`, `UserName`)
	b, err = json.Marshal(r)
	assert.NoError(t, err)
	assert.Equal(t, `[{"UserName":"ok"}]`, string(b))
}
