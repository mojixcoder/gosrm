package gosrm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResponse_IsOk(t *testing.T) {
	res := Response{Code: CodeOK}
	assert.True(t, res.IsOk())

	res.Code = CodeInvalidQuery
	assert.False(t, res.IsOk())
}
