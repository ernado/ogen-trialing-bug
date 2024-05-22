package bug

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"bug/oas"
)

func TestResolve(t *testing.T) {
	var h oas.Server
	route, ok := h.FindRoute(http.MethodGet, "/clusters/")
	require.True(t, ok)
	require.Equal(t, "listClusters", route.OperationID())
}
