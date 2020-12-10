package main

import (
	"testing"
)

func TestCountNumberOfQuestionsAnsweredByAnyone(t *testing.T) {

	tests := []struct {
		input string
		want  int
	}{

		{input: "abc\n", want: 3},
		{input: "a\nb\nc\n\nab\nac\n", want: 6},
	}
	for _, tt := range tests {
		t.Run("TestCountNumberOfAnsweredQuestions", func(t *testing.T) {
			if got := CountNumberOfQuestionsAnsweredByAnyone(tt.input); got != tt.want {
				t.Errorf("CountNumberOfAnsweredQuestions() = %v, want %v", got, tt.want)
			}
		})
	}
}

