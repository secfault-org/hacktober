#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BANNER \
  "Welcome to Challenge 8! Let's print the flag\n"

void setup() {
    setvbuf(stdout, NULL, _IONBF, 0);
    setvbuf(stderr, NULL, _IONBF, 0);
}

void vulnerable() {
  char* flag = getenv("FLAG");
  char buffer[32];

  printf("Enter your name: ");
  fgets(buffer, sizeof(buffer), stdin);
  printf("Hello ");
  printf(buffer);
  printf("!\nYou will never get the flag!\n");
}

int main() {
  setup();
  printf("%s\n", BANNER);

  vulnerable();
  exit(0);
}
