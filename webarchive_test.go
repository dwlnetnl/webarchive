package webarchive

import (
	"fmt"
	"hash/crc32"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew_content(t *testing.T) {
	assert := assert.New(t)

	r, err := os.Open("testdata/example.webarchive")
	require.NoError(t, err)

	a, err := New(r)
	assert.Equal(uint32(0x61f96615), crc32.ChecksumIEEE(a.Content.Data))
	assert.Equal("text/html", a.Content.MIMEType)
	assert.Equal("UTF-8", a.Content.Encoding)
	assert.Equal("http://example.org/", a.Content.URL)
	assert.Empty(a.Content.FrameName)
	assert.Empty(a.Resources)
	assert.NoError(err)
}

func TestNew_resources(t *testing.T) {
	assert := assert.New(t)

	r, err := os.Open("testdata/wikipedia.webarchive")
	require.NoError(t, err)

	a, err := New(r)
	assert.Equal(uint32(0x7a24addb), crc32.ChecksumIEEE(a.Content.Data))
	assert.Equal("text/html", a.Content.MIMEType)
	assert.Equal("UTF-8", a.Content.Encoding)
	assert.Equal("https://en.wikipedia.org/wiki/Webarchive", a.Content.URL)
	assert.Empty(a.Content.FrameName)
	assert.Len(a.Resources, 13)
	assert.NoError(err)

	resources := []struct {
		data uint32
		mime string
		enc  string
		url  uint32
		resp uint32
	}{
		{0xe0352bc6, "text/css", "utf-8", 0xa88c8657, 0x3e495c5f},
		{0xd5fee4ca, "text/css", "utf-8", 0x348faf99, 0xe1b3e4d0},
		{0xab5b0999, "text/javascript", "utf-8", 0xe9d253b4, 0x58d61084},
		{0xf3c27e95, "text/javascript", "utf-8", 0xfeaf026a, 0x96e53929},
		{0x59637a, "image/png", "", 0xbd73e04a, 0x6dd3aad5},
		{0xf79b5fc7, "image/png", "", 0x8d210858, 0xfd3b31c4},
		{0xfbea27a5, "image/jpeg", "", 0x482f361f, 0xdf0d81fd},
		{0xdae1d80a, "image/png", "", 0x2ab8618b, 0x570a3987},
		{0x285bcd96, "image/png", "", 0x5903bafa, 0xad214977},
		{0xc296a07a, "image/png", "", 0x23103ba7, 0xd845f71},
		{0x1af78e35, "image/png", "", 0x9af068df, 0x19fc7a6},
		{0x7c4edfa, "image/png", "", 0x7de2a6e8, 0x36084aec},
		{0x9885885, "text/css", "utf-8", 0x8c2feaff, 0xb01023b6},
	}

	for i := 0; i < 13; i++ {
		sr := a.Resources[i]
		r := resources[i]
		s := fmt.Sprintf("subresource %d", i)

		assert.Equal(r.data, crc32.ChecksumIEEE(sr.Data), s)
		assert.Equal(r.mime, sr.MIMEType, s)
		assert.Equal(r.enc, sr.Encoding, s)
		assert.Equal(r.url, crc32.ChecksumIEEE([]byte(sr.URL)), s)
		assert.Equal(r.resp, crc32.ChecksumIEEE(sr.Response), s)
	}
}
