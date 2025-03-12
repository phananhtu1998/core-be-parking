package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Lưu ý đặt tên func có chữ Test đầu tiên để nhận diện được hàm test
func TestAdd(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )
	// actual := AddOne(1)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), input %d, actual %d", input, output, actual)
	// }
	assert.Equal(t, AddOne(2), 4, "AddOne(2) shoud be 3")
	assert.NotEqual(t, AddOne(2), 4, "AddOne(2) should not be 4")
	assert.Nil(t, nil, nil)
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3)
	fmt.Println("Not excuting")
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("excuting")
}
