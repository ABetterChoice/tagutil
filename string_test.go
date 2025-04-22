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
			if got := e.EQ(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("EQ() = %v, want %v", got, tt.want)
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
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"1-ios-mac"},
				configValue:  "ios",
			},
			want: true,
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

func Test_stringExecutor_IsEmpty(t *testing.T) {
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
			name: "empty slice",
			args: args{
				unitTagValue: []string{},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "nil slice",
			args: args{
				unitTagValue: nil,
				configValue:  "",
			},
			want: true,
		},
		{
			name: "slice with one empty string",
			args: args{
				unitTagValue: []string{""},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "slice with multiple empty strings",
			args: args{
				unitTagValue: []string{"", ""},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "slice with non-empty string",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "",
			},
			want: false,
		},
		{
			name: "slice with a mix of empty and non-empty strings",
			args: args{
				unitTagValue: []string{"", "hello"},
				configValue:  "",
			},
			want: false,
		},
		{
			name: "slice with only whitespace strings",
			args: args{
				unitTagValue: []string{" "},
				configValue:  "",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.IsEmpty(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_IsNotEmpty(t *testing.T) {
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
			name: "empty slice",
			args: args{
				unitTagValue: []string{},
				configValue:  "",
			},
			want: false,
		},
		{
			name: "nil slice",
			args: args{
				unitTagValue: nil,
				configValue:  "",
			},
			want: false,
		},
		{
			name: "slice with one empty string",
			args: args{
				unitTagValue: []string{""},
				configValue:  "",
			},
			want: false,
		},
		{
			name: "slice with multiple empty strings",
			args: args{
				unitTagValue: []string{"", ""},
				configValue:  "",
			},
			want: false,
		},
		{
			name: "slice with non-empty string",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "slice with a mix of empty and non-empty strings",
			args: args{
				unitTagValue: []string{"", "hello"},
				configValue:  "",
			},
			want: true,
		},
		{
			name: "slice with only whitespace strings",
			args: args{
				unitTagValue: []string{" "},
				configValue:  "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.IsNotEmpty(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_Like(t *testing.T) {
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
			name: "empty slice",
			args: args{
				unitTagValue: []string{},
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "nil slice",
			args: args{
				unitTagValue: nil,
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "matching value in single-element slice",
			args: args{
				unitTagValue: []string{"test"},
				configValue:  "test",
			},
			want: true,
		},
		{
			name: "non-matching value in single-element slice",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "matching one value in multi-element slice",
			args: args{
				unitTagValue: []string{"hello", "test", "world"},
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "non-matching value in multi-element slice",
			args: args{
				unitTagValue: []string{"hello", "world"},
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "configValue partially matches an element",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hel",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.Like(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("Like() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringExecutor_NotLike(t *testing.T) {
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
			name: "empty slice",
			args: args{
				unitTagValue: []string{},
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "nil slice",
			args: args{
				unitTagValue: nil,
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "matching value in single-element slice",
			args: args{
				unitTagValue: []string{"test"},
				configValue:  "test",
			},
			want: false,
		},
		{
			name: "non-matching value in single-element slice",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "test",
			},
			want: true,
		},
		{
			name: "matching one value in multi-element slice",
			args: args{
				unitTagValue: []string{"hello", "test", "world"},
				configValue:  "test",
			},
			want: true,
		},
		{
			name: "non-matching value in multi-element slice",
			args: args{
				unitTagValue: []string{"hello", "world"},
				configValue:  "test",
			},
			want: true,
		},
		{
			name: "configValue partially matches an element",
			args: args{
				unitTagValue: []string{"hello"},
				configValue:  "hel",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &stringExecutor{}
			if got := e.NotLike(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("NotLike() = %v, want %v", got, tt.want)
			}
		})
	}
}
