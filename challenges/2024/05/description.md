Challenge 5
===========

Das ist die erste Challenge, in dem wir wirklich ret2win ausnutzen können.
Der Code ist sehr ähnlich zu Challenge 4, aber dieses Mal werden die Singals nicht gehandelt.
Schaffst du es trotzdem `win` aufzurufen um die Flag zu bekommen?

```c
#include <err.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define BANNER \
  "Welcome to Challenge 5! Signals will not help you here. Good luck!\n"

char *gets(char *);

void win() {
  write(STDOUT_FILENO, getenv("FLAG"), 32);
  exit(0);
}

void setup() {
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stdin, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
}

void vulnerable() {
  char buffer[64];
  void *ret;

  gets(buffer);

  ret = __builtin_return_address(0);
  printf("and will be returning to %p\n", ret);
}

int main() {
  setup();
  printf("%s\n", BANNER);

  vulnerable();
  exit(0);
}
```
