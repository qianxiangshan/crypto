package kernel

/*#include <stdio.h>
#include <unistd.h>
#include <sys/socket.h>
#include <linux/if_alg.h>
#include <linux/socket.h>
#include <string.h>
#include <stdint.h>
#include <errno.h>

#ifndef SOL_ALG
#define SOL_ALG 279
#endif

int test(void)
{

  uint8_t data[16]={1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6};
  uint8_t key[16]={1,2,3,4,5,6,7,8,9,0,1,2,3,4,5,6};
  int opfd;
  int tfmfd;
  struct sockaddr_alg sa = {
    .salg_family = AF_ALG,
    .salg_type = "skcipher",
    .salg_name = "cbc(twofish)",
  };
  struct msghdr msg = {};
  struct cmsghdr *cmsg;
  char cbuf[CMSG_SPACE(4) + CMSG_SPACE(20)] = {0};
  char buf[16];
  struct af_alg_iv *iv;
  struct iovec iov;
  int i;
  int result;

  tfmfd = socket(AF_ALG, SOCK_SEQPACKET, 0);

  result=bind(tfmfd, (struct sockaddr *)&sa, sizeof(sa));
  //printf("bind result %d %d\n",result,errno);

  result=setsockopt(tfmfd, SOL_ALG, ALG_SET_KEY,key, 16);
 // printf("setsocket result %d %d\n",result,errno);

  opfd = accept(tfmfd, NULL, 0);

  msg.msg_control = cbuf;
  msg.msg_controllen = sizeof(cbuf);

  cmsg = CMSG_FIRSTHDR(&msg);
  cmsg->cmsg_level = SOL_ALG;
  cmsg->cmsg_type = ALG_SET_OP;
  cmsg->cmsg_len = CMSG_LEN(4);
  *(__u32 *)CMSG_DATA(cmsg) = ALG_OP_ENCRYPT;

  uint8_t tiv[16]={0};
  cmsg = CMSG_NXTHDR(&msg, cmsg);
  cmsg->cmsg_level = SOL_ALG;
  cmsg->cmsg_type = ALG_SET_IV;
  cmsg->cmsg_len = CMSG_LEN(20);
  iv = (void *)CMSG_DATA(cmsg);
  iv->ivlen = 16;
  memcpy(iv->iv, tiv, 16);

  iov.iov_base = data;
  iov.iov_len = 16;

  msg.msg_iov = &iov;
  msg.msg_iovlen = 1;

  result=sendmsg(opfd, &msg, 0);
 // printf("sendmsg result %d %d\n",result,errno);
  result=read(opfd, buf, 16);
 // printf("read result %d %d\n",result,errno);

//   for (i = 0; i < 16; i++) {
//     printf("%02x", (unsigned char)buf[i]);
//   }
//   printf("\n");

  close(opfd);
  close(tfmfd);

  return 0;
}
*/
import (
	"C"
)

func KernelTwofish(key, in, out []byte) {
	C.test()
}
