// Package tagutil ...
package tagutil

import "testing"

func Test_setExecutor_IsSubSet(t *testing.T) {
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
				configValue:  "",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5", "6"},
				configValue:  "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &setExecutor{}
			if got := e.IsSubSet(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IsSubSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setExecutor_IsSuperSet(t *testing.T) {
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
				configValue:  "",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5", "6"},
				configValue:  "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &setExecutor{}
			if got := e.IsSuperSet(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IsSuperSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setExecutor_EG(t *testing.T) {
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
				configValue:  "",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", "5", "6"},
				configValue:  "",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"12345678"},
				configValue:  "12345678",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"12345678"},
				configValue:  "12345679",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &setExecutor{}
			if got := e.EG(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("EG() = %v, want %v", got, tt.want)
			}
		})
	}
}
