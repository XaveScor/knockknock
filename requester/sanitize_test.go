package requester

import "testing"
import "github.com/stretchr/testify/assert"

func TestCommon(t *testing.T) {
	domain, _ := sanitize("http://a.b/asd")

	assert.Equal(t, "a.b", domain)
}

func TestAnyString(t *testing.T) {
	_, err := sanitize("asdasdadas")

	assert.Error(t, err)
}

func TestDomainOnly(t *testing.T) {
	domain, _ := sanitize("a.b")

	assert.Equal(t, "a.b", domain)
}

func TestIpOnly(t *testing.T) {
	domain, _ := sanitize("http://192.168.0.1/asdasd")

	assert.Equal(t, "192.168.0.1", domain)
}

func TestIpPort(t *testing.T) {
	domain, _ := sanitize("http://192.168.0.1:123/asdasd")

	assert.Equal(t, "192.168.0.1:123", domain)
}
