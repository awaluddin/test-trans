package trans

import "testing"

func TestSearchNearest(t *testing.T) {

	var nmap []float32 = []float32{-1.5, 0, 4.4, 5, 6, 7}

	tests := []struct {
		name string
		want float32
		inpt float32
	}{
		{
			name: "should get 4.4",
			want: 4.4,
			inpt: 4,
		},
		{
			name: "should get 4.4",
			want: 4.4,
			inpt: 4.5,
		},
		{
			name: "should get 6",
			want: 6,
			inpt: 5.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchNearest(tt.inpt, nmap); got != tt.want {
				t.Errorf("SearchNearest() = %v, want %v", got, tt.want)
			}
		})
	}
}
