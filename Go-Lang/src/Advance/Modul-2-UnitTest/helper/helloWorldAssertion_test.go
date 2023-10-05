package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Haris")
	assert.Equal(t, "Hello Haris",result, "Result must be 'Hello Haris'") 

	// assert jika gagal -> automatic call Fail()

	// assert.Equal(unitTest, expected, resultTest, message)
	// message akan di print jika error
	fmt.Println("Dieksekusi")
}

func TestHelloWorldDimasRequire(t *testing.T) {
	result := HelloWorld("Dimas")
	require.Equal(t, "Hello Dimas",result, "Result must be 'Hello Dimas'") 

	// require jika gagal -> memanggil fungsi FailNow()
	fmt.Println("Dieksekusi")
}

func TestSkip(t *testing.T){
	sistemOperasi := "MAC"
	if sistemOperasi == "MAC" {
		t.Skip("Sistem operasi berupa MAC (tidak support)") 
		//melakukan skip pada unit test, line dibawahnya tidak di run
		//statusnya -> PASS (success)
	}
	result := HelloWorld("Dimas")
	require.Equal(t, "Hello Dimas",result, "Result must be 'Hello Dimas'") 
}