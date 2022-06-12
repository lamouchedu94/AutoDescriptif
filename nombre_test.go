package main

import (
	"reflect"
	"testing"
)

func TestItoNombre(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want nombre
	}{
		{"0",
			args{0},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}),
		},
		{"1",
			args{1},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
		{"2020",
			args{2020},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 2, 0, 2, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ItoNombre(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ItoNombre() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nombre_inc(t *testing.T) {
	type args struct {
		i int
	}
	tests := []struct {
		name string
		args args
		want nombre
	}{
		{"0",
			args{0},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 1}),
		},
		{"1",
			args{1},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 2}),
		},
		{"2020",
			args{2020},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 2, 0, 2, 1}),
		},
		{"2029",
			args{2029},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 2, 0, 3, 0}),
		},
		{"2999",
			args{2999},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 3, 0, 0, 0}),
		},
		{"2899",
			args{2899},
			nombre([10]uint8{0, 0, 0, 0, 0, 0, 2, 9, 0, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ItoNombre(tt.args.i)
			n.inc()
			if n != tt.want {
				t.Errorf("Inc() = %v, want %v", n, tt.want)

			}
		})
	}
}

func Test_nombre_estAutodescriptif(t *testing.T) {
	tests := []struct {
		name string
		n    nombre
		want bool
	}{
		{"0",
			nombre(ItoNombre(0)),
			false,
		},
		{"1",
			ItoNombre(1),
			false,
		},
		{"2020",
			ItoNombre(2020),
			true,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.estAutodescriptif(); got != tt.want {
				t.Errorf("nombre.estAutodescriptif() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nombre_int(t *testing.T) {
	tests := []struct {
		name string
		n    nombre
		want int
	}{
		{
			"0",
			ItoNombre(0),
			0,
		},
		{
			"9099",
			ItoNombre(9099),
			9099,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.int(); got != tt.want {
				t.Errorf("nombre.int() = %v, want %v", got, tt.want)
			}
		})
	}
}
