// Package tagutil ...
package tagutil

import "testing"

func Test_stringExecutor_EG(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.EG(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("EG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_GT(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "b",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b"},
				configValue:  "ba",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"z"},
				configValue:  "11",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.GT(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("GT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_GTE(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "b",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b"},
				configValue:  "ba",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"z"},
				configValue:  "11",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.GTE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_IN(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "b",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b"},
				configValue:  "b;a;c;dddd",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd"},
				configValue:  "b;a;c;dddd",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "eeee"},
				configValue:  "b;a;c;dddd",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd", ""},
				configValue:  ";b;a;c;dddd",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"z"},
				configValue:  "zz",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.IN(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_LT(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "b",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b"},
				configValue:  "bbb",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd"},
				configValue:  "dddd11",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "eeee"},
				configValue:  "eeee",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd", ""},
				configValue:  "dddd1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"z"},
				configValue:  "zz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.LT(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_LTE(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "b",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b"},
				configValue:  "bbb",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd"},
				configValue:  "dddd",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "eeee"},
				configValue:  "eeee",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd", ""},
				configValue:  "dddd1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"z"},
				configValue:  "zz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.LTE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_NE(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "b",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b"},
				configValue:  "bbb",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd"},
				configValue:  "dddd",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "eeee"},
				configValue:  "eeee",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd", ""},
				configValue:  "dddd1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"z"},
				configValue:  "zz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.NE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("NE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_NotIN(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "b",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b"},
				configValue:  "bbb",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd"},
				configValue:  "dddd",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "eeee"},
				configValue:  "eeee",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd", ""},
				configValue:  "dddd1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd", ""},
				configValue:  "dddd;b;c;;",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4"},
				configValue:  "dddd;b;c;;",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1", "2", "3", "4", ""},
				configValue:  "dddd;b;c;;",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"z"},
				configValue:  "zz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.NotIN(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("NotIN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_REGEXP(t *testing.T) {
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
				unitTagValue: []string{"hello", "world"},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"hello", "hello"},
				configValue:  "hello",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{},
				configValue:  "hello",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"ba"},
				configValue:  "^b",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"bbb"},
				configValue:  "b*",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd"},
				configValue:  "dddd",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "eeee"},
				configValue:  "^[a-zA-Z]",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"112231", "1aabb", "4eeb", "5hello"},
				configValue:  "^[a-zA-Z]",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"b", "c", "dddd", ""},
				configValue:  "^[1-9]",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"112231", "1aabb", "4eeb", "5hello"},
				configValue:  "^[1-9]",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"112231", "1aabb", "4eeb", "5hello"},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"112231", "1aabb", "4eeb", "5hello"},
				configValue:  "*",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.REGEXP(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("REGEXP() = %v, want %v", got, tt.want)
			}
		})
	}
}
