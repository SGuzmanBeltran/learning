package main

import (
	"testing"
)

func TestConvertNumberToFizzBuzzString(t *testing.T) {
	tests := []struct {
		name    string
		input   int
		want    string
		wantErr bool
	}{
		{
			name:    "convert 1 to 1",
			input:   1,
			want:    "1",
			wantErr: false,
		},
		{
			name:    "convert 2 to 2",
			input:   4,
			want:    "4",
			wantErr: false,
		}, {
			name:    "convert 4 to 4",
			input:   4,
			want:    "4",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fizzBuzz := &FizzBuzz{}
			got, err := fizzBuzz.Convert(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("Convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
