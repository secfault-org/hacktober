#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <signal.h>

#define BANNER "Welcome to Challenge 4! Sinals will help you this time. Can you get the flag?"

char* gets(char* buffer);

void win() {
  char* flag = getenv("FLAG");
  if (flag == NULL) {
    puts("Uh oh, the flag is missing. Please contact an admin if you are running ");
    exit(1);
  }
  printf("Flag: %s\n", flag);
  exit(0);
}

void setup() {
  setvbuf(stdout, NULL, _IONBF, 0);
  setvbuf(stdin, NULL, _IONBF, 0);
  setvbuf(stderr, NULL, _IONBF, 0);
  signal(SIGSEGV, win);
}

void vulnerable() {
  char buf[32];
  gets(buf);
}

int main() {
  setup();
  printf("%s\n", BANNER);

  vulnerable();
  exit(0);
}

