#include <stdio.h>
#include <stdlib.h>

#define BANNER "Welcome to Challenge 2! Can you change 'change_me' to 0x496c5962?"

char* gets(char* buffer);

void win() {
  char* flag = getenv("FLAG");
  if (flag == NULL) {
    puts("Uh oh, the flag is missing. Please contact an admin if you are running ");
    exit(1);
  }
  printf("Flag: %s\n", flag);
}

void setup() {
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stdin, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
}

int main() {
  setup();

  struct {
    char buffer[64];
    volatile int change_me;
  } locals;

  printf("%s\n", BANNER);

  locals.change_me = 0;
  gets(locals.buffer);

  if (locals.change_me == 0x496c5962) {
    win();
  } else {
    printf("Getting closer! changeme is currently 0x%08x, we want 0x496c5962\n",
        locals.change_me);
  }

  exit(0);
}
