// Package tagutil ...
package tagutil

import "testing"

func Test_versionExecutor_EG(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.EQ(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("EQ() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_GT(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: true, // 9-9-9-9 > 9
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.GT(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("GT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_GTE(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: true, // 9-9-9-9 > 9
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.GTE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_IN(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1", "2.2.2.2"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456788"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788;12aabb123.3454aabb12.123456712.123456789",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.IN(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_LCRC(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.10",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.10:9.9.9.11",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9:9.9.9.10",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0:10.1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1", "2.2.2.2"},
				configValue:  "1:zzz",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456788"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788:12aabb123.3454aabb12.123456712.123456789",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.LCRC(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LCRC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_LCRO(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.10",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.10:9.9.9.11",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9:9.9.9.10",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0:10.1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1", "2.2.2.2"},
				configValue:  "1:zzz",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456788"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788:12aabb123.3454aabb12.123456712.123456789",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.LCRO(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LCRO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_LORC(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.10",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.10:9.9.9.11",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9:9.9.9.10",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0:10.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "8.0:9.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1", "2.2.2.2"},
				configValue:  "1:zzz",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456788"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788:12aabb123.3454aabb12.123456712.123456789",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.LORC(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LORC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_LORO(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.10",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.10:9.9.9.11",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.8:9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9:9.9.9.10",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0:10.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "8.0:9.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1", "2.2.2.2"},
				configValue:  "1:zzz",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456788"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788:12aabb123.3454aabb12.123456712.123456789",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.LORO(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LORO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_LT(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false, // 9-9-9-9 > 9
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.1.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456787"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.LT(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_LTE(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: false, // 9-9-9-9 > 9
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.1.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456787"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456788"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.LTE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_NE(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9", "1.1.1.1"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.NE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("NE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_NotIN(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta", "1.1.1.1", "2.2.2.2"},
				configValue:  "9.0.0.9-beta;1.1.1.1",
			},
			want: false,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456789"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788",
			},
			want: true,
		},
		{
			name: "big version",
			args: args{
				unitTagValue: []string{"12aabb123.3454aabb12.123456712.123456788"},
				configValue:  "12aabb123.3454aabb12.123456712.123456788;12aabb123.3454aabb12.123456712.123456789",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.NotIN(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("NotIN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_REGEXP(t *testing.T) {
	type args struct {
		unitTagValue []string
		configValue  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.9.9.8"},
				configValue:  "9.9.9.9",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9-9-9-9"},
				configValue:  "9.9.9.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9*",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0*",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.0"},
				configValue:  "9.0.0.0",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "^9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.0.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.0.*.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.*.8",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9"},
				configValue:  "9.*.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"990.0-9"},
				configValue:  "9..*.9",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9.0.0.9-beta"},
				configValue:  "9.0.0.9",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.REGEXP(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("REGEXP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_versionExecutor_compareByDecimal(t *testing.T) {
	type args struct {
		unitTagValueVersion string
		configVersion       string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal",
			args: args{
				unitTagValueVersion: "1234567812345678987654321",
				configVersion:       "1234567812345678987654321",
			},
			want: 0,
		},
		{
			name: "normal",
			args: args{
				unitTagValueVersion: "",
				configVersion:       "1234567812345678987654321",
			},
			want: -1,
		},
		{
			name: "normal",
			args: args{
				unitTagValueVersion: "1234567812345678987654321",
				configVersion:       "",
			},
			want: 1,
		},
		{
			name: "normal",
			args: args{
				unitTagValueVersion: "12345678.12345678.987654321",
				configVersion:       "12345678.12345678.987654321",
			},
			want: 0,
		},
		{
			name: "normal",
			args: args{
				unitTagValueVersion: "",
				configVersion:       "123456781.23456789.87654321",
			},
			want: -1,
		},
		{
			name: "normal",
			args: args{
				unitTagValueVersion: "12345678.1234567898.7654321",
				configVersion:       "",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &versionExecutor{}
			if got := e.compareByDecimal(tt.args.unitTagValueVersion, tt.args.configVersion); got != tt.want {
				t.Errorf("compareByDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
