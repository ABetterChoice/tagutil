// Package tagutil ...
package tagutil

import "testing"

func Test_numberExecutor_EG(t *testing.T) {
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
			name: "long number",
			args: args{
				unitTagValue: []string{"12.34234234343"},
				configValue:  "12.34234234343",
			},
			want: true,
		},
		{
			name: "empty number",
			args: args{
				unitTagValue: []string{},
				configValue:  "12.34234234343",
			},
			want: false,
		},
		{
			name: "empty number",
			args: args{
				unitTagValue: []string{""},
				configValue:  "12.34234234343",
			},
			want: false,
		},
		{
			name: "not number",
			args: args{
				unitTagValue: []string{"12.34234234343A"},
				configValue:  "12.34234234343A",
			},
			want: false,
		},
		{
			name: "not equal",
			args: args{
				unitTagValue: []string{"-12.34234234343A"},
				configValue:  "12.34234234343A",
			},
			want: false,
		},
		{
			name: "with zero",
			args: args{
				unitTagValue: []string{"0012.34234234343"},
				configValue:  "12.34234234343",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"8080"},
				configValue:  "8080",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"8080612312.4353242342344324123416457567657452345252634575684662352635756725623"},
				configValue:  "8080612312.4353242342344324123416457567657452345252634575684662352635756725623",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"8080612312.4353242342344324123416457567657452345252634575684662352635756725623", "1"},
				configValue:  "8080612312.4353242342344324123416457567657452345252634575684662352635756725623",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.EG(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("EG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_GT(t *testing.T) {
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
				unitTagValue: []string{"9999"},
				configValue:  "9998",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.998",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"09.999"},
				configValue:  "09.998",
			},
			want: true,
		},
		{
			name: "leader zero",
			args: args{
				unitTagValue: []string{"0009.999"},
				configValue:  "09.998",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: false,
		},
		{
			name: "char",
			args: args{
				unitTagValue: []string{"0a09.998"},
				configValue:  "09.997",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.997",
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.999",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.999",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.GT(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("GT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_GTE(t *testing.T) {
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
				unitTagValue: []string{"9999"},
				configValue:  "9998",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.998",
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.999",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"09.999"},
				configValue:  "09.998",
			},
			want: true,
		},
		{
			name: "leader zero",
			args: args{
				unitTagValue: []string{"0009.999"},
				configValue:  "09.998",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: false,
		},
		{
			name: "char",
			args: args{
				unitTagValue: []string{"0a09.998"},
				configValue:  "09.997",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.997",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.999",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.GTE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("GTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_LT(t *testing.T) {
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
				unitTagValue: []string{"9999"},
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.998",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: nil,
				configValue:  "9.998",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: nil,
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"1"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.999",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"09.999"},
				configValue:  "09.998",
			},
			want: false,
		},
		{
			name: "leader zero",
			args: args{
				unitTagValue: []string{"0009.999"},
				configValue:  "09.998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: true,
		},
		{
			name: "char",
			args: args{
				unitTagValue: []string{"0a09.998"},
				configValue:  "09.997",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.997",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.999",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.LT(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_LTE(t *testing.T) {
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
				unitTagValue: []string{"9999"},
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.998",
			},
			want: false,
		},
		{
			name: "equal",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.999",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"09.999"},
				configValue:  "09.998",
			},
			want: false,
		},
		{
			name: "leader zero",
			args: args{
				unitTagValue: []string{"0009.999"},
				configValue:  "09.998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: true,
		},
		{
			name: "char",
			args: args{
				unitTagValue: []string{"0a09.998"},
				configValue:  "09.997",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.997",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.999",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.LTE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_NE(t *testing.T) {
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
				unitTagValue: []string{"9999"},
				configValue:  "9998",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.998",
			},
			want: true,
		},
		{
			name: "equal",
			args: args{
				unitTagValue: []string{"9.999"},
				configValue:  "9.999",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"09.999"},
				configValue:  "09.998",
			},
			want: true,
		},
		{
			name: "leader zero",
			args: args{
				unitTagValue: []string{"0009.999"},
				configValue:  "09.998",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"0009.998"},
				configValue:  "09.999",
			},
			want: true,
		},
		{
			name: "char",
			args: args{
				unitTagValue: []string{"0a09.998"},
				configValue:  "09.997",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.997",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.998"},
				configValue:  "999999999999999999999999999999999.999",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.999"},
				configValue:  "999999999999999999999999999999999.999",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.NE(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("NE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_IN(t *testing.T) {
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
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.999"},
				configValue:  "999999999999999999999999999999999.999;1",
			},
			want: true,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1.999",
			},
			want: true,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1;2;3;4;5;1.999",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1;2;3",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.IN(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("IN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_NotIN(t *testing.T) {
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
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.999"},
				configValue:  "999999999999999999999999999999999.999;1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1;3;4;5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "zz",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1.999",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1;2;3;4;5;1.999",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"8", "9", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1;2;3",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.NotIN(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("NotIN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_LCRC(t *testing.T) {
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
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz:1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.9999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"9999999999999999999999999999999999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.9999999999999999999999999999999999999999999999"},
				configValue:  "1:2",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1:1.5555555555555555555555",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999999999999999999999999999999999999999999999999999999999999999999999999"},
				configValue:  "2:3",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1:2",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1:6",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"8", "9", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1:1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.LCRC(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LCRC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_LCRO(t *testing.T) {
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
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz:1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.9999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"9999999999999999999999999999999999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.9999999999999999999999999999999999999999999999"},
				configValue:  "1:2",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1:1.5555555555555555555555",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999999999999999999999999999999999999999999999999999999999999999999999999"},
				configValue:  "2:3",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1:2",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1:6",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1:7",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"8", "9", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1:1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.LCRO(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LCRO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_LORC(t *testing.T) {
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
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz:1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.9999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"9999999999999999999999999999999999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.9999999999999999999999999999999999999999999999"},
				configValue:  "1:2",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1:1.5555555555555555555555",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999999999999999999999999999999999999999999999999999999999999999999999999"},
				configValue:  "2:3",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1:2",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1:6",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1:7",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "-1:7",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"8", "9", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3", "-1"},
				configValue:  "1:1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.LORC(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LORC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_numberExecutor_LORO(t *testing.T) {
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
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: false,
		},
		{
			name: "float64",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:5",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: nil,
				configValue:  "9998",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "zz:1",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"9999"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "normal",
			args: args{
				unitTagValue: []string{"zz"},
				configValue:  "1:zz",
			},
			want: false,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"999999999999999999999999999999999.9999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: true,
		},
		{
			name: "big number",
			args: args{
				unitTagValue: []string{"9999999999999999999999999999999999"},
				configValue:  "999999999999999999999999999999999.999:9999999999999999999999999999999999",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.9999999999999999999999999999999999999999999999"},
				configValue:  "1:2",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999"},
				configValue:  "1:1.5555555555555555555555",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1.999999999999999999999999999999999999999999999999999999999999999999999999"},
				configValue:  "2:3",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3"},
				configValue:  "1:2",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1:6",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "1:7",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "6"},
				configValue:  "-1:7",
			},
			want: true,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"8", "9", "6"},
				configValue:  "1;2;3;4;5",
			},
			want: false,
		},
		{
			name: "normal number",
			args: args{
				unitTagValue: []string{"1", "2", "3", "-1"},
				configValue:  "1:1",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &numberExecutor{}
			if got := e.LORO(tt.args.unitTagValue, tt.args.configValue); got != tt.want {
				t.Errorf("LORO() = %v, want %v", got, tt.want)
			}
		})
	}
}
