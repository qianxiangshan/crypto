package ct

import "testing"

func TestCtwofish(t *testing.T) {

	var data = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
	var key = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
	var out = make([]byte, 16)
	CtwoFish(key, data, out)
	for i := 0; i < 16; i++ {
		t.Logf("%02x", out[i])
	}
	t.Logf("\n")
}

func BenchmarkCtwofish(b *testing.B) {

	var data = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
	var key = []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 4, 5, 6}
	var out = make([]byte, 16)
	for i := 0; i < b.N; i++ {
		CtwoFish(key, data, out)
	}
}
