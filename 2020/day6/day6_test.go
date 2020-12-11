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
		t.Run("TestCountNumberOfQuestionsAnsweredByAnyone", func(t *testing.T) {
			if got := CountNumberOfQuestionsAnsweredByAnyone(tt.input); got != tt.want {
				t.Errorf("CountNumberOfQuestionsAnsweredByAnyone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountNumberOfQuestionsAnsweredByEveryone(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		input string
		want  int
	}{

		{input: "abc\n", want: 3},
		{input: "a\nb\nc\n\nab\nac\n", want: 1},
	}

	for _, tt := range tests {
		t.Run("TestCountNumberOfQuestionsAnsweredByEveryone", func(t *testing.T) {
			if got := CountNumberOfQuestionsAnsweredByEveryone(tt.input); got != tt.want {
				t.Errorf("CountNumberOfQuestionsAnsweredByEveryone() = %v, want %v", got, tt.want)
			}
		})
	}
}
