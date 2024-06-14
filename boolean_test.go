// Package tagutil ...
package tagutil

import "testing"

func Test_booleanExecutor_EG(t *testing.T) {
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
			name: "normal bool",
			args: args{
				unitTagValue: nil,
				configValue:  "1",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{},
				configValue:  "1",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"1"},
				configValue:  "1",
			},
			want: true,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"t"},
				configValue:  "1",
			},
			want: true,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"true"},
				configValue:  "1",
			},
			want: true,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"1"},
				configValue:  "false",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"t"},
				configValue:  "f",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"true"},
				configValue:  "0",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"0"},
				configValue:  "1",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"f"},
				configValue:  "1",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"false"},
				configValue:  "1",
			},
			want: false,
		},
		{
			name: "normal bool",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1",
			},
			want: false,
		},
		{
			name: "invalid bool",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "invalid bool",
			args: args{
				unitTagValue: []string{"1"},
				configValue:  "1;1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &booleanExecutor{}
			if got := e.EQ(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("EQ() = %v, want %v", got, tt.want)
			}
		})
	}
}
