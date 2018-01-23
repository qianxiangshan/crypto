package twofish

/*
#cgo LDFLAGS:  -lgcrypt -lgpg-error
#include <gcrypt.h>
#include <gpg-error.h>
#include <stdint.h>
*/
import (
	"C"
)
import (
	"errors"
	"fmt"
	"strconv"
	"unsafe"
)

const (
	TwofishKeySize = 32
	BLOCKSIZE      = 16
)

type TwofishCipher struct {
	//twofish 加密句柄
	fd        C.gcry_cipher_hd_t
	blockSize int
}

type KeySizeError int

func (k KeySizeError) Error() string {
	return "dana-crypto/twofish: invalid key size " + strconv.Itoa(int(k))
}

func char2string(src []C.char) string {
	data := ""
	for i := 0; i < len(src); i++ {
		data = fmt.Sprintf("%c", src[i])
	}
	return data
}

func init() {
	ver := C.CString(C.GCRYPT_VERSION)
	C.gcry_check_version(ver)
	C.free(unsafe.Pointer(ver))
}

// 创建一个twofish的加密key。 如果key相同则使用相同的即可。 无需使用。 是否是并发安全的呢
func NewCipher(key []byte) (*TwofishCipher, error) {

	var err_string [100]C.char
	kl := len(key)
	if kl != TwofishKeySize {
		return nil, KeySizeError(kl)
	}
	var cipher TwofishCipher
	var err C.gcry_error_t
	err = C.gcry_cipher_open(&cipher.fd, C.GCRY_CIPHER_TWOFISH, C.GCRY_CIPHER_MODE_ECB, 0)
	if err != 0 {
		C.gpg_strerror_r(C.gpg_error_t(err), (*C.char)(&err_string[0]), 100)
		return nil, errors.New("twofish new cipher error " + char2string(err_string[:]))
	}
	cipher.blockSize = BLOCKSIZE
	err = C.gcry_cipher_setkey(cipher.fd, (unsafe.Pointer)(&key[0]), C.size_t(TwofishKeySize))
	if err != 0 {
		C.gpg_strerror_r(C.gpg_error_t(err), (*C.char)(&err_string[0]), 100)
		return nil, errors.New("twofish set key error " + char2string(err_string[:]))
	}
	return &cipher, nil
}

// BlockSize returns the cipher's block size.
func (tf *TwofishCipher) BlockSize() int {
	return tf.blockSize
}

// Encrypt encrypts the first block in src into dst. dst src 长度相同， 都是16B的整数倍
func (tf *TwofishCipher) Encrypt(dst, src []byte) {

	if len(dst) != len(src) || len(src)%tf.blockSize != 0 {
		panic("twofish encrypt dst or src slice len not 16")
	}
	C.gcry_cipher_encrypt(tf.fd, (unsafe.Pointer)(&dst[0]), C.size_t(len(dst)), (unsafe.Pointer)(&src[0]), C.size_t(len(src)))
	//printf("gcry_cipher_encrypt %s\n",gpg_strerror(err));
}

// Decrypt decrypts the first block in src into dst.dst src 长度相同， 都是16B的整数倍
func (tf *TwofishCipher) Decrypt(dst, src []byte) {
	if len(dst) != len(src) || len(src)%tf.blockSize != 0 {
		panic("twofish encrypt dst or src slice len not 16")
	}
	C.gcry_cipher_decrypt(tf.fd, (unsafe.Pointer)(&dst[0]), C.size_t(len(dst)), (unsafe.Pointer)(&src[0]), C.size_t(len(src)))

}

func (tf *TwofishCipher) Destroy() {
	C.gcry_cipher_close(tf.fd)
}
