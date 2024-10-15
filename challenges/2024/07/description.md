Challenge 7
===========

Diese Challenge ist etwas anders als die Vorherigen.
Es fehlt die Funktion `win()`, die die Flag ausgibt.
Stattdessen liegt die Flag in einer Datei auf dem System und `/flag`.

Es wird Zeit für ein bisschen Shellcode!
Der Buffer ist dieses mal größer das heißt es ist viel Platz für spannende Sachen auf dem Stack.

Für diese Challenge gehen wir ca. 20 Jahre in die Vergangenheit.
Es wurden noch keine Sicherheitsmechanismen wie ASLR oder NX/DEP implementiert.
Das heißt das Segment, in dem sich der Stack befindet, ist ausführbar und immer an der gleichen Stelle.

Konkret wurde das Binary mit folgenden Flags kompiliert:
```bash
gcc -Wall -Wextra -Werror -Wno-deprecated-declarations -O0 -std=c99 -pedantic -fno-strict-aliasing -fno-omit-frame-pointer -fno-stack-protector -fno-stack-check -fno-pie -m32 -no-pie -z execstack -o challenge7 challenge7.c
```

Und ASLR wurde deaktiviert:
```bash
  echo 0 | sudo tee /proc/sys/kernel/randomize_va_space
```

Hinweis:
* Es gibt eine vielzahl an bereits geschriebener Shellcodes
  * z.B. unter http://shell-storm.org/shellcode/index.html
  * Ausserdem gibt es auch Tools, die Shellcode generieren können (z.B. msfvenom)

```c
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
  gets(buffer);
}

int main() {
  setup();
  printf("%s\n", BANNER);

  vulnerable();
  exit(0);
}
```
