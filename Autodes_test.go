package main

import "testing"

func Test_EstAutodesc2(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.segment(i,
		{
			name: "1",
			args: args{n: "1"},
			want: false,
		},
		{
			name: "2020",
			args: args{n: "2020"},
			want: true,
		},
		{
			name: "21200",
			args: args{n: "21200"},
			want: true,
		},
		{
			name: "21201",
			args: args{n: "21201"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EstAutodesc2(tt.args.n); got != tt.want {
				t.Errorf("est_autodesc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
