Challenge 1
===========

Die erste Challenge ist hervorragend geeignet, um sich mit dem Stack und den Basics von Buffer Overflows vertraut zu machen.

Außerdem empfehle ich, sich mit [GDB](https://www.gnu.org/software/gdb/) (mit **gef** oder **pwndbg**) und [Pwntools](https://github.com/Gallopsled/pwntools) vertraut zu machen.
Auch wenn es jetzt noch nicht notwendig ist, wird es in den späteren Challenges sehr hilfreich sein.

Der Code ist sehr einfach gehalten und es gibt nur eine einzige Schwachstelle, die es zu finden gilt.

```c
#include <stdio.h>
#include <stdlib.h>
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

  if (locals.change_me != 0) {
    win();
  } else {
    puts("Uh oh, 'changeme' has not yet been changed.");
  }

  exit(0);
}
```

Um die Challenge zu starten einfach innnerhalb der Challenge (also hier) ctrl+s drücken.
Daraufhin wird die Challenge gestartet und ist über das Netzwerk erreichbar. Der Port wird
unten in der Statusleiste angezeigt.

```bash
nc hacktober2024.secfault.org <port>
```

Wenn die Challenge gelöst wurde, kann die Flag mit `ctrl+f` eingereicht werden.

Viel Erfolg!

