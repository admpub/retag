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

func TestSetJSONTag2(t *testing.T) {
	type row struct {
		Name string
	}
	type data struct {
		Username string `json:""`
		List     []row
	}
	v := data{
		Username: `ok`,
		List: []row{
			{Name: `test`},
		},
	}
	r := SetJSONTag(&v, `Username`, `UserName`, `List.Name`, `name`)
	b, err := json.Marshal(r)
	assert.NoError(t, err)
	assert.Equal(t, `{"UserName":"ok","List":[{"name":"test"}]}`, string(b))

	values := []data{
		{
			Username: `ok`,
			List: []row{
				{Name: `test`},
			},
		},
	}
	r = SetJSONTag(&values, `Username`, `UserName`, `List.Name`, `title`)
	b, err = json.Marshal(r)
	assert.NoError(t, err)
	assert.Equal(t, `[{"UserName":"ok","List":[{"title":"test"}]}]`, string(b))
}
