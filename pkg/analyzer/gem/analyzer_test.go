package gem

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatching(t *testing.T) {
	a := Analyzer()
	assert.True(t, a.Match("/usr/local/bundle/specifications/rails-4.2.5.1.gemspec", nil))
}