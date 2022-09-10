package initCheck

import (
	"github.com/labstack/echo/v4"
	"testing"
)

func TestCheckRun(t *testing.T) {
	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckRun(tt.args.e)
		})
	}
}
