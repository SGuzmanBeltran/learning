package main

import (
	"math"
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
			name:    "discard 0",
			input:   0,
			want:    "0",
			wantErr: true,
		},
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
		{
			name:    "convert 3 to Fizz",
			input:   3,
			want:    "FizzFizz",
			wantErr: false,
		},
		{
			name:    "convert 6 to Fizz",
			input:   6,
			want:    "Fizz",
			wantErr: false,
		}, {
			name:    "convert 9 to Fizz",
			input:   9,
			want:    "Fizz",
			wantErr: false,
		},
		{
			name:    "convert 5 to BuzzBuzz",
			input:   5,
			want:    "BuzzBuzz",
			wantErr: false,
		},
		{
			name:    "convert 10 to Buzz",
			input:   10,
			want:    "Buzz",
			wantErr: false,
		}, {
			name:    "convert 15 to FizzBuzzBuzz",
			input:   15,
			want:    "FizzBuzzBuzz",
			wantErr: false,
		},
		{
			name:    "convert 20 to Buzz",
			input:   20,
			want:    "Buzz",
			wantErr: false,
		},
		{
			name:    "convert 33 to FizzFizz",
			input:   33,
			want:    "FizzFizz",
			wantErr: false,
		},
		{
			name:    "convert 51 to FizzBuzz",
			input:   51,
			want:    "FizzBuzz",
			wantErr: false,
		},
		{
			name:    "convert 35 to FizzBuzzBuzz",
			input:   35,
			want:    "FizzBuzzBuzz",
			wantErr: false,
		},
		{
			name:    "accept max int64",
			input:   math.MaxInt64,
			want:    "FizzBuzz",
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
