package trans

import "testing"

func TestGetTLD(t *testing.T) {

	tests := []struct {
		name   string
		getUrl string
		want   string
	}{
		{
			name:   "com",
			getUrl: "http://www.google.com/abc",
			want:   "com",
		},
		{
			name:   "com",
			getUrl: "http://google.com/abc",
			want:   "com",
		},
		{
			name:   "co.id",
			getUrl: "http://www.google.co.id/abc",
			want:   "co.id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTLD(tt.getUrl); got != tt.want {
				t.Errorf("GetTLD() = %v, want %v", got, tt.want)
			}
		})
	}
}
