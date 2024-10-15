#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BANNER \
  "Welcome to Challenge 7! Make some space for the fabulous shellcode!\n"

char *gets(char *);

void setup() {
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stdin, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
}

void vulnerable() {
  char buffer[512];
  printf("%p\n", buffer);
  gets(buffer);
}

int main() {
  setup();
  printf("%s\n", BANNER);

  vulnerable();
  exit(0);
}
