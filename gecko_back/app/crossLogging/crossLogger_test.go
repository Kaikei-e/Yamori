package crossLogging

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"io"
	"reflect"
	"testing"
)

func TestWatcher(t *testing.T) {
	type args struct {
		e      *echo.Echo
		config Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Watcher(tt.args.e, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Watcher() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_configureLogger(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want *zerolog.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := configureLogger(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configureLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_logRotation(t *testing.T) {
	type args struct {
		config Config
	}
	tests := []struct {
		name string
		args args
		want io.Writer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := logRotation(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
