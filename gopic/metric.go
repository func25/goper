package gopic

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/func25/gop/gopmat"
)

func TestMe(t *testing.T) {
	x := runtime.MemStats{}
	runtime.ReadMemStats(&x)
	fmt.Println(gopmat.JSON.StrPretty(x))
}
