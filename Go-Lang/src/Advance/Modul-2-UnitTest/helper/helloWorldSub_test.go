package helper

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubtest(t *testing.T){
	t.Run("Haris", func(t *testing.T) { // t.Run ("nama subtest", anonymous func)-> running subtest
		result := HelloWorld("Haris")
		require.Equal(t, "Hello Haris",result, "Result must be 'Hello Haris'") 
	})
	t.Run("Dimas", func(t *testing.T) {
		result := HelloWorld("Dimas")
		require.Equal(t, "Hello Dimas",result, "Result must be 'Hello Dimas'") 
	})
	// go test -v -run TestUnitTest/TestSubtest -> running only 1 specific subtest
	// go test -v -run /TestSubtest -> eksekusi semua subtest dengan nama func "TestSubtest"
}