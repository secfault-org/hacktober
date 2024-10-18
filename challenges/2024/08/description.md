Challenge 8
===========

So erstmal genug von Buffer Overflows.
Jetzt wird es Zeit für eine andere Art von Schwachstelle: Format String Vulnerabilities.

In dieser Challenge gibt es eine Funktion,
die eine Eingabe des Benutzers direkt in einen `printf`-Befehl einsetzt.

Oft sieht man in C sowas wie:

```c
printf("%s", name);
```

Warum schreibt man das so und nicht einfach:

```c
printf(name);
```

Und was bedeutet dieses `%s` überhaupt?

Das `%s` ist ein sogenannter Specifier in einem Format String.
Es gibt viele verschiedene Specifier wie z.B.:

| Specifier | Beschreibung |
|-----------|--------------|
| `%s`      | String       |
| `%d`      | Integer      |
| `%x`      | Hexadezimal  |
| `%p`      | Pointer      |

Alle möglichen Format Specifier kann unter `man 3 printf` nachgelesen werden.
Sie können als Platzhalter für die Argumente dienen, die nach dem Format String kommen.
z.B.:

```c
printf("Hello %s!\n", name);
```

statt

```c
printf("Hello ");
printf(name);
printf("!\n");
```

Aber was ist jetzt das Problem, wenn man `printf(name)` schreibt?
Das Problem entsteht, wenn der String `name` vom Benutzer eingegeben wird.
Dann hat der Benutzer den Format String in der Hand,
währende der Entwickler nur den String ausgeben wollten, den der User eingegeben hat.

Okay, das ist vielleicht nicht das, was der Entwickler wollte, aber was kann der Benutzer damit anstellen?

In dieser Challenge geht es darum Daten zu Leaken, die der Benutzer eigentlich nicht sehen sollte.

```c
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
  char buffer[32];
  char* flag = getenv("FLAG");

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
```

Hinweis: Diese Challenge ist wieder in 32Bit kompiliert,
da ich es so einfacher finde die Format String Vulnerability nachzuvollziehen.
(Wirst du beim Debuggen auch merken)

Das hängt wieder mit der Calling Convention zusammen,
da in 32-Bit Mode alle Argumente auf dem Stack übergeben werden.

Wenn du die Flag gefunden hast, nimm dir am besten hier nochmal die Zeit,
um die anderen zurückgegebenen Werte von `printf` zu verstehen.

Vielleicht versuchst du es auch mal mit 64-Bit und schaust dir an, wie es sich verändert.
Warum ist es so?

Viel Erfolg! 🍀