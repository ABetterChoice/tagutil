// Package tagutil ...
package tagutil

import (
	"git.code.oa.com/tencent_abtest/protocol/protoc_cache_server"
	"testing"
)

//go:generate go test ./... -coverprofile=size_coverage.out -gcflags "all=-N -l"
//go:generate go tool cover -html=size_coverage.out
func TestIsHit(t *testing.T) {
	type args struct {
		tagType      protoc_cache_server.TagType
		operator     protoc_cache_server.Operator
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "big number",
			args: args{
				tagType:      protoc_cache_server.TagType_TAG_TYPE_NUMBER,
				operator:     protoc_cache_server.Operator_OPERATOR_GTE,
				unitTagValue: []string{"98765432123456789867654433221"},
				configValue:  "98765432123456789867654433220",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				tagType:      protoc_cache_server.TagType_TAG_TYPE_VERSION,
				operator:     protoc_cache_server.Operator_OPERATOR_GTE,
				unitTagValue: []string{"987654321.234567898.676544.33221"},
				configValue:  "987654321.234567898.676544.33220",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				tagType:      protoc_cache_server.TagType_TAG_TYPE_VERSION,
				operator:     protoc_cache_server.Operator_OPERATOR_LTE,
				unitTagValue: []string{"987654321.234567898.676544.33221"},
				configValue:  "987654321.234567898.676544.33220",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				tagType:      protoc_cache_server.TagType_TAG_TYPE_BOOLEAN,
				operator:     protoc_cache_server.Operator_OPERATOR_LTE,
				unitTagValue: []string{"987654321.234567898.676544.33221"},
				configValue:  "987654321.234567898.676544.33220",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsHit(tt.args.tagType, tt.args.operator, tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IsHit() = %v, want %v", got, tt.want)
			}
		})
	}
}
