#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

char *gets(char *);

#define BANNER "Welcome to Challenge 1! Can you change 'change_me'?"

void win() {
  char* flag = getenv("FLAG");
  if (flag == NULL) {
    puts("Uh oh, the flag is missing. Please contact an admin if you are running ");
    exit(1);
  }
  printf("Flag: %s\n", flag);
}

int main() {
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
  struct {
    char buffer[64];
    volatile int change_me;
  } locals;

  printf("%s\n", BANNER);

  locals.change_me = 0;
  gets(locals.buffer);

  if (locals.change_me != 0) {
    win();
  } else {
    puts("Uh oh, 'changeme' has not yet been changed.");
  }

  exit(0);
}
