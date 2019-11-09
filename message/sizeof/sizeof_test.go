package sizeof

import (
	"math"
	"testing"
)

func TestBytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want int32
	}{
		{
			name: "nil",
			data: nil,
			want: ArrayLength,
		},
		{
			name: "some",
			data: []byte("hello world"),
			want: ArrayLength + int32(len("hello world")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Bytes(tt.data); got != tt.want {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringArray(t *testing.T) {
	tests := []struct {
		name string
		ss   []string
		want int32
	}{
		{
			name: "nil",
			ss:   nil,
			want: 4,
		},
		{
			name: "simple",
			ss:   []string{"hello", "world"},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringArray(tt.ss); got != tt.want {
				t.Errorf("StringArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Array(t *testing.T) {
	tests := []struct {
		name string
		ii   []int32
		want int32
	}{
		{
			name: "nil",
			ii:   nil,
			want: 4,
		},
		{
			name: "some",
			ii:   []int32{1, 2, 3},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Array(tt.ii); got != tt.want {
				t.Errorf("Int32Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Array(t *testing.T) {
	tests := []struct {
		name string
		ii   []int64
		want int32
	}{
		{
			name: "nil",
			ii:   nil,
			want: 4,
		},
		{
			name: "some",
			ii:   []int64{1, 2, 3},
			want: 28,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int64Array(tt.ii); got != tt.want {
				t.Errorf("Int64Array() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarInt(t *testing.T) {
	tests := []struct {
		name string
		i    int64
		want int32
	}{
		{
			name: "1 byte",
			i:    0,
			want: 1,
		},
		{
			name: "3 bytes",
			i:    math.MaxInt16,
			want: 3,
		},
		{
			name: "10 bytes",
			i:    math.MaxInt64,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VarInt(tt.i); got != tt.want {
				t.Errorf("VarInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkVarInt(t *testing.B) {
	for i := 0; i < t.N; i++ {
		got := VarInt(int64(i))
		if got == 0 {
			t.Fatalf("got 0; want not 0")
		}
	}
}

func TestVarBytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want int32
	}{
		{
			name: "nil",
			data: nil,
			want: 1,
		},
		{
			name: "0",
			data: []byte{},
			want: 1,
		},
		{
			name: "some",
			data: []byte("hello world"),
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VarBytes(tt.data); got != tt.want {
				t.Errorf("VarBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVarString(t *testing.T) {
	tests := []struct {
		name string
		data string
		want int32
	}{
		{
			name: "blank",
			data: "",
			want: 1,
		},
		{
			name: "some",
			data: "hello world",
			want: 12,
		},
		{
			name: "unicode",
			data: "你好",
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VarString(tt.data); got != tt.want {
				t.Errorf("VarBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
