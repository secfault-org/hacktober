#include <stdio.h>
#include <stdlib.h>

#define BANNER "Welcome to Challenge 3! It's time to test your skills. Can you call win to get the Flag?"

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
    int (*fp)();
  } locals;

  printf("%s\n", BANNER);

  locals.fp = NULL;
  gets(locals.buffer);

  if (locals.fp) {
    printf("calling function pointer @ %p\n", (void *)(size_t) locals.fp);
    locals.fp();
  } else {
    printf("function pointer remains unmodified :~( better luck next time!\n");
  }

  exit(0);
}
