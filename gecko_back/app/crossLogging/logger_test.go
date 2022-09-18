package crossLogging

import "testing"

func TestLoggerSetup(t *testing.T) {

	type args struct {
		path string
	}

	testCase := map[string]struct {
		name    string
		args    args
		wantErr bool
	}{
		"pass1": {
			name: "pass1",
			args: args{
				path: "yamori",
			},
			wantErr: false,
		},
		"fail1": {
			name: "fail1",
			args: args{
				path: "errrrrrrrrrrrrrrrrrrrrrrrr",
			},
			wantErr: true,
		},
	}

	tests := []struct {
		name string
		args args
	}{
		{
			testCase["pass1"].name,
			testCase["pass1"].args,
		},
		{
			testCase["fail1"].name,
			testCase["fail1"].args,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			LoggerSetup(tt.args.path)
		})
	}
}
