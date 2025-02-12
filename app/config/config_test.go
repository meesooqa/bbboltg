package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	r, err := Load("testdata/config.yml")
	require.NoError(t, err)

	assert.Equal(t, 2*time.Minute, r.UpdateInterval)

	assert.Equal(t, 2, len(r.Channels), "2 channels")
	assert.Equal(t, "from1", r.Channels["name1"].From)
	assert.Equal(t, "to1", r.Channels["name1"].To)

	//assert.Equal(t, "to_bolt1", r.Channels["name1"].ToBolt)
	assert.Equal(t, "⚡️", r.Channels["name1"].Filters.Bolt.Prefix)
	//assert.Equal(t, "", r.Channels["name1"].ToImportant)
	assert.Nil(t, r.Channels["name1"].Filters.Important)

	assert.Equal(t, "from2", r.Channels["name2"].From)
	assert.Equal(t, "to2", r.Channels["name2"].To)
}

func TestLoadConfigNotFoundFile(t *testing.T) {
	r, err := Load("/tmp/7ff6fe56-ad2c-455a-b0d2-ec11afee532c.txt")
	assert.Nil(t, r)
	assert.EqualError(t, err, "open /tmp/7ff6fe56-ad2c-455a-b0d2-ec11afee532c.txt: no such file or directory")
}

func TestLoadConfigInvalidYaml(t *testing.T) {
	r, err := Load("testdata/file.txt")

	assert.Nil(t, r)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `Not Yaml` into config.Conf")
}
