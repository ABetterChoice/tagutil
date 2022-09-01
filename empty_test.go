// Package tagutil ...
package tagutil

import "testing"

func Test_emptyExecutor_IsEmpty(t *testing.T) {
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
			name: "nil",
			args: args{
				unitTagValue: nil,
			},
			want: true,
		},
		{
			name: "empty",
			args: args{
				unitTagValue: []string{},
			},
			want: true,
		},
		{
			name: "not empty",
			args: args{
				unitTagValue: []string{"1"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &emptyExecutor{}
			if got := e.IsEmpty(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_emptyExecutor_IsNotEmpty(t *testing.T) {
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
			name: "nil",
			args: args{
				unitTagValue: nil,
			},
			want: false,
		},
		{
			name: "empty",
			args: args{
				unitTagValue: []string{},
			},
			want: false,
		},
		{
			name: "not empty",
			args: args{
				unitTagValue: []string{"1"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &emptyExecutor{}
			if got := e.IsNotEmpty(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
