package increment

import (
	"fmt"
	"testing"
)

func TestIncrement(t *testing.T) {

	t.Run("patch Increment", func(t *testing.T) {
		var a = Increment("aaa")
		fmt.Println(a)

		if ret := Increment("1.0.0"); ret != "1.0.1" {
			t.Fatalf("1.0.0 must be incremented to 1.0.1 but %s", ret)
		}
	})
}
