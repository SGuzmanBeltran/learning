package exercise2

import (
	"testing"
)

type Tests struct {
	name string
	text string
	want string
}

func TestStringManipulation_Reverse(t *testing.T) {
	tests := []Tests{
		{"Reverse with spaces", "C# is a interesting language", "egaugnal gnitseretni a si #C"},
		{"Reverse with even number of runes", "GO is a interesting language", "egaugnal gnitseretni a si OG"},
		{"Reverse with odd number of runes", "I have travel a few times", "semit wef a levart evah I"},
		{"Empty string", " ", " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &StringManipulation{}
			if got := sm.Reverse(tt.text); got != tt.want {
				t.Errorf("StringManipulation.Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringManipulation_ToUpperCase(t *testing.T) {
	tests := []Tests{
		{"Uppercase with spaces", "C# is a interesting language", "C# IS A INTERESTING LANGUAGE"},
		{"Uppercase with even number of runes", "GO is a interesting language", "GO IS A INTERESTING LANGUAGE"},
		{"Uppercase with odd number of runes", "I have travel a few times", "I HAVE TRAVEL A FEW TIMES"},
		{"Empty string", " ", " "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &StringManipulation{}
			if got := sm.ToUpperCase(tt.text); got != tt.want {
				t.Errorf("StringManipulation.ToUpperCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringManipulation_RemoveSpaces(t *testing.T) {
	tests := []Tests{
		{"Remove space with normal spaces", "C# is a interesting language", "C#isainterestinglanguage"},
		{"Remove space with new line", "GO is a interesting \n language", "GOisainterestinglanguage"},
		{"Remove space with tabulations", "I have travel a few\ttimes", "Ihavetravelafewtimes"},
		{"Remove space with multiple spaces", "I have travel a  few   times", "Ihavetravelafewtimes"},
		{"Empty string", " ", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &StringManipulation{}
			if got := sm.RemoveSpaces(tt.text); got != tt.want {
				t.Errorf("StringManipulation.RemoveSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}

//Benchmark example
func BenchmarkRemoveSpaces(b *testing.B) {
    sm := &StringManipulation{}
    for i := 0; i < b.N; i++ {
        sm.RemoveSpaces("GO is a interesting language")
    }
}
