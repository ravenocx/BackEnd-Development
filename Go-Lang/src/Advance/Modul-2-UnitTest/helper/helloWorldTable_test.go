package helper

import (
	"testing"

	_"github.com/stretchr/testify/assert"
)

func TestHelloWorldTable(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string 
	}{
		{
			name : "HelloWorld(Haris)",
			request : "Haris",
			expected : "Hello Haris",
		},
		{
			name : "HelloWorld(Dimas)",
			request : "Dimas",
			expected : "Hello Dimas",
		},
	}

	for _, test := range tests{
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			if result != test.expected{
				t.Errorf("\nExpected : %s\nGot : %s", test.expected, result)
			}
			// assert.Equal(t, test.expected, result )
		})
	}
}