package main

import (
	"testing"

)

func TestCleanInput(t *testing.T) {
	cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "  charmander  bulbaSAUR  PIKACHU  ",
		expected: []string{"charmander", "bulbasaur", "pikachu"},
	},
	{
		input:    "  BOOTS got in the server room again  ",
		expected: []string{"boots", "got", "in", "the", "server", "room", "again"},
	},
}
for _, c := range cases {
	actual := cleanInput(c.input)
	if len(actual) != len(c.expected){
		t.Errorf("len does not match")
		continue
	}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("does not match")
			t.Fail()
		}else {
			i++
		}
	}
}
}