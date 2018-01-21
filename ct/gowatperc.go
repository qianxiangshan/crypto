package ct

/*
#cgo LDFLAGS:  -lgcrypt -lgpg-error
#include <gcrypt.h>
#include <gpg-error.h>
#include <stdio.h>
#include <stdint.h>

int ctwofish(uint8_t *key,uint8_t *data,uint8_t *out){
    gcry_error_t err;
    gcry_cipher_hd_t twofishfd;
    err=gcry_cipher_open(&twofishfd,GCRY_CIPHER_TWOFISH128,GCRY_CIPHER_MODE_CBC,0);
    //printf("open %s\n",gpg_strerror(err));

    err=gcry_cipher_setkey(twofishfd, key, 16);
    //printf("gcry_cipher_setkey %s\n",gpg_strerror(err));

    err=gcry_cipher_encrypt(twofishfd,out,16,data,16);
    //printf("gcry_cipher_encrypt %s\n",gpg_strerror(err));
    gcry_cipher_close(twofishfd);
    return 0;
}
*/
import (
	"C"
)

func CtwoFish(key, in, out []byte) error {

	C.ctwofish(((*C.uint8_t)(&key[0])), ((*C.uint8_t)(&in[0])), ((*C.uint8_t)(&out[0])))
	return nil
}
