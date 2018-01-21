package kernel

import (
	"testing"
)

func BenchmarkKernelfish(b *testing.B) {

	for i := 0; i < b.N; i++ {
		KernelTwofish(nil, nil, nil)
	}

}
