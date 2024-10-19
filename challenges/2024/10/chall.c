#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define BANNER "Welcome to Challenge 10! Can you change 'change_me' to 0xdeadbeef?"

unsigned int change_me;

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

void vulnerable(char *str) {
    printf(str);
}

int main() {
  setup();

  char buffer[100];
  memset(buffer, 0, sizeof(buffer));
  printf("%s\n", BANNER);
  printf("%p\n", &change_me);

  if (fgets(buffer, sizeof(buffer) - 1, stdin) == NULL) {
    puts("Unable to read input.");
    exit(1);
  }

  vulnerable(buffer);

  if (change_me == 0xdeadbeef) {
    win();
  } else if (change_me == 0) {
    printf("Change_me is still 0. Try again!\n");
  } else {
     printf("Getting closer! change_me is currently 0x%08x, but we want 0xdeadbeef\n",
            change_me);
  }

  exit(0);
}
