package initCheck

import (
	"testing"
)

func TestEnvLoader(t *testing.T) {

	t.Run("panic", func(t *testing.T) {
		defer func() {
			if re := recover(); re == nil {
				t.Errorf("function should panic")
			}
		}()

		EnvLoader()

	})

	// if you want to test the function without panic, you can use this code
	//t.Run("pass", func(t *testing.T) {
	//	EnvLoader()
	//	t.Log("pass")
	//})
}
