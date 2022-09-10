package initCheck

import (
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"testing"
)

func TestCheckRun(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Error(err)
	}

	type args struct {
		e *echo.Echo
	}
	tests := []struct {
		name string
		args args
	}{
		{"pass", args{e: &echo.Echo{}}},
		{"fail", args{e: &echo.Echo{}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CheckRun(tt.args.e)
		})
	}

}
