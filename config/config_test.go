package config

import (
	"reflect"
	"testing"
)

var cfg Config = Config{
	Username:    "royge",
	AccessToken: "your-github-access-token",
	Exclude:     []string{"repo1", "repo2"},
	Timeout:     2,
}

const filename = "../testdata/config.test.json"

func TestGetConfig(t *testing.T) {
	tt := []struct {
		filename string
		c        *Config
	}{
		{
			filename: filename,
			c:        &cfg,
		},
		{
			filename: "../testdata/notfound.json",
			c:        nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.filename, func(t *testing.T) {
			c, err := GetConfig(tc.filename)
			if tc.c != nil {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
			}

			if tc.c != nil && !reflect.DeepEqual(*tc.c, *c) {
				t.Fatalf("%v and %v are not equal", tc.c, *c)
			}
		})
	}
}
