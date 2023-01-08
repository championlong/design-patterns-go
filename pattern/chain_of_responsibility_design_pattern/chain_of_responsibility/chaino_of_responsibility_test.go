package chain_of_responsibility

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChainOf(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	assert.Equal(t, false, chain.doFilter("test"))

	chain.AddFilter(&PoliticalWordFilter{})
	assert.Equal(t, true, chain.doFilter("test"))
}
