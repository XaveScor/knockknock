package requester

import "testing"
import "github.com/stretchr/testify/assert"

func TestGoogle(t *testing.T) {
	available := TouchWebsite("http://google.com")

	assert.False(t, available)
}

// pornhub is unavailable in KZ
func TestPornhub(t *testing.T) {
	available := TouchWebsite("https://pornhub.com")

	assert.False(t, available)
}
