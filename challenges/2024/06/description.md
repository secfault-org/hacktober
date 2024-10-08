Challenge 6
===========

In dieser Challenge müssen wir win mit Parametern aufrufen. Dazu müssen wir die Call-Convention verstehen.
Hinweis: Unix artige Systeme verwenden die System V ABI.

Das Program wurde in 32-Bit kompiliert und kann hier heruntergeladen werden:
https://share.riseup.net/#9k9dTIKzBxMn2hXMQfZLAQ

```c
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BANNER \
  "Welcome to Challenge 6! Do you have the right argument to win?\n"

char *gets(char *);

void win(unsigned int arg1, unsigned int arg2) {
  char* flag = getenv("FLAG");
  if (flag == NULL) {
    puts("Uh oh, the flag is missing. Please contact an admin if you are running ");
    exit(1);
  }
  if (arg1 == 0xdeadbeef && arg2 == 0xcafebabe) {
    printf("Okay, I believe you! Here is your flag: %s\n", flag);
  } else {
    printf("Try again! arg1 = %x, arg2 = %x\n", arg1, arg2);
  }
  exit(0);
}

void setup() {
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stdin, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
}

void vulnerable() {
  char buffer[64];
  gets(buffer);
}

int main() {
  setup();
  printf("%s\n", BANNER);

  vulnerable();
  exit(0);
}
```
