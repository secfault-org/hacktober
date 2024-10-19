#include <stdio.h>
#include <stdlib.h>

#define BANNER "Welcome to Challenge 9! Can you still change 'change_me'?"

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
  setvbuf(stderr, NULL, _IONBF, 0);
}

int main() {
  setup();

  struct {
    char dest[32];
    volatile int change_me;
  } locals;
  char buffer[16];

  printf("%s\n", BANNER);

  if (fgets(buffer, sizeof(buffer) - 1, stdin) == NULL) {
    puts("Unable to read input.");
    exit(1);
  }
  buffer[sizeof(buffer) - 1] = 0;
  locals.change_me = 0;

  sprintf(locals.dest, buffer);

  if (locals.change_me == (int) 0xdeadbeef) {
    win();
  } else if (locals.change_me != 0) {
      printf("Getting closer! change_me is currently 0x%08x, but we want 0xdeadbeef\n",
              locals.change_me);
  } else {
    puts("Uh oh, 'change_me' has not yet been changed.");
  }

  exit(0);
}
