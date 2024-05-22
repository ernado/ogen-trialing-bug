//go:build tools

package bug

import (
	_ "github.com/ogen-go/ogen"
	_ "github.com/ogen-go/ogen/conv"
	_ "github.com/ogen-go/ogen/gen"
	_ "github.com/ogen-go/ogen/gen/genfs"
	_ "github.com/ogen-go/ogen/middleware"
	_ "github.com/ogen-go/ogen/ogenerrors"
	_ "github.com/ogen-go/ogen/otelogen"
)
