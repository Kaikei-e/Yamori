package router

import (
	"github.com/labstack/echo/v4"
	"reflect"
	"testing"
)

func TestRouter(t *testing.T) {
	tests := []struct {
		name    string
		want    *echo.Echo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Router()
			if (err != nil) != tt.wantErr {
				t.Errorf("Router() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router() got = %v, want %v", got, tt.want)
			}
		})
	}
}
