package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		configFile string
	}
	tests := []struct {
		name       string
		args       args
		wantConfig *Config
		wantErr    bool
	}{
		{
			name: "Nominal use case",
			args: args{configFile: "../../testdata/golden/config/config_nominal.golden"},
			wantConfig: &Config{
				Api: ApiConfig{
					Url:                  "https://api.ethplorer.io/",
					ApiKey:               "XXXX",
					GetLastBlockFunction: "getLastBlock",
				},
				Server: HttServerConfig{
					Port:      8080,
					SecretKey: "R0hklKt1OKiOwCqL3llE",
				},
			},
			wantErr: false,
		},
		{
			name:       "Invalid file",
			args:       args{configFile: "../../testdata/golden/config/invalid_json.golden"},
			wantConfig: nil,
			wantErr:    true,
		},
		{
			name:       "No config file",
			args:       args{configFile: "none"},
			wantConfig: nil,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			service := ConfigService{
				configFile: tt.args.configFile,
			}
			gotConfig, err := service.loadConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("loadConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("loadConfiguration() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
