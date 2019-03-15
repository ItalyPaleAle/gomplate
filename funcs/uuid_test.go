package funcs

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	uuidPattern   = "^[[:xdigit:]]{8}-(?:[[:xdigit:]]{4}-){3}[[:xdigit:]]{12}$"
	uuidV1Pattern = "^[[:xdigit:]]{8}-[[:xdigit:]]{4}-1[[:xdigit:]]{3}-[89ab][[:xdigit:]]{3}-[[:xdigit:]]{12}$"
	uuidV4Pattern = "^[[:xdigit:]]{8}-[[:xdigit:]]{4}-4[[:xdigit:]]{3}-[89ab][[:xdigit:]]{3}-[[:xdigit:]]{12}$"
)

func TestV1(t *testing.T) {
	u := UUIDNS()
	i, err := u.V1()
	assert.NoError(t, err)
	assert.Regexp(t, uuidV1Pattern, i)
}

func TestV4(t *testing.T) {
	u := UUIDNS()
	i, err := u.V4()
	assert.NoError(t, err)
	assert.Regexp(t, uuidV4Pattern, i)
}

func TestNil(t *testing.T) {
	u := UUIDNS()
	i, err := u.Nil()
	assert.NoError(t, err)
	assert.Equal(t, "00000000-0000-0000-0000-000000000000", i)
}

func TestIsValid(t *testing.T) {
	u := UUIDNS()
	var in interface{}
	in = false
	i, err := u.IsValid(in)
	assert.NoError(t, err)
	assert.False(t, i)

	in = 12345
	i, err = u.IsValid(in)
	assert.NoError(t, err)
	assert.False(t, i)

	testdata := []interface{}{
		"123456781234123412341234567890ab",
		"12345678-1234-1234-1234-1234567890ab",
		"urn:uuid:12345678-1234-1234-1234-1234567890ab",
		"{12345678-1234-1234-1234-1234567890ab}",
	}

	for _, d := range testdata {
		i, err = u.IsValid(d)
		assert.NoError(t, err)
		assert.True(t, i)
	}
}

func TestParse(t *testing.T) {
	u := UUIDNS()
	var in interface{}
	in = false
	_, err := u.Parse(in)
	assert.Error(t, err)

	in = 12345
	_, err = u.Parse(in)
	assert.Error(t, err)

	in = "12345678-1234-1234-1234-1234567890ab"
	testdata := []interface{}{
		"123456781234123412341234567890ab",
		"12345678-1234-1234-1234-1234567890ab",
		"urn:uuid:12345678-1234-1234-1234-1234567890ab",
		must(url.Parse("urn:uuid:12345678-1234-1234-1234-1234567890ab")),
		"{12345678-1234-1234-1234-1234567890ab}",
	}

	for _, d := range testdata {
		uid, err := u.Parse(d)
		assert.NoError(t, err)
		assert.Equal(t, in, uid.String())
	}
}
