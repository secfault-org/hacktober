Challenge 3
===========

In dieser Challenge wird das erste mal der Kontrollfluss des Programms beeinflusst.
Damit tauchen wir in die Welt der Exploits und Wired-Machines das erste Mal richtig ein.

Die Aufgabe ist es, die Funktion `win` aufzurufen, um die Flagge zu erhalten.
Hierfür ist auch das erstmal das Binary relevant, welches wir analysieren müssen.

Du findest die Binary unter folgendem Link:
https://share.riseup.net/#QV457bkZ-H9yTD-ic59aEg

Hinweis: Die Binary ohne PIE compiliert, damit die Adressen im Binary immer gleich sind.

```c
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
```
