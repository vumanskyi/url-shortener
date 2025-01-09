package service

import "testing"

func TestGenerateShortURL(t *testing.T) {
	type args struct {
		originalURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid URL",
			args: args{
				originalURL: "https://example.com",
			},
			want: "1niM5DUD",
		},
		{
			name: "Invalid URL",
			args: args{
				originalURL: "not a url",
			},
			want: "",
		},
		{
			name: "Empty URL",
			args: args{
				originalURL: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateShortURL(tt.args.originalURL); got != tt.want {
				t.Errorf("GenerateShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidURL(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid URL",
			args: args{
				rawURL: "http://example.com",
			},
			want: true,
		},
		{
			name: "Valid URL",
			args: args{
				rawURL: "https://example.com",
			},
			want: true,
		},
		{
			name: "Invalid URL",
			args: args{
				rawURL: "not a url",
			},
			want: false,
		},
		{
			name: "Empty URL",
			args: args{
				rawURL: "",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidURL(tt.args.rawURL); got != tt.want {
				t.Errorf("IsValidURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
