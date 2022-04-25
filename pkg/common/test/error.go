package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// CheckErrorsType Compare errors type as string
func CheckErrorsType(t *testing.T, expErr, gotErr error) {
	gotType := fmt.Sprintf("%T", gotErr)
	expType := fmt.Sprintf("%T", expErr)
	require.Equal(t, expType, gotType, "unexpected error type")
}
