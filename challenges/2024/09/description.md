Challenge 9
===========

Nicht nur `printf` hat einen Format-String, sondern auch andere Funktionen wie:
* printf
* fprintf
* dprintf
* sprintf
* snprintf
* vprintf
* vfprintf
* vdprintf
* vsprintf
* vsnprintf


Wenn Format-String in kombination mit `sprintf` verwendet wird,
kann checks umgehen und wieder Buffer-Overruns erzeugen.

Ich habe dieses Mal gleich die Challenges kombiniert.
Also versuche erste die `change_me` irgendwie abzuändern und schaue,
ob du es gezielt auf `0xdeadbeef` setzen kannst.

In der Challenge ist es wieder sehr ähnlich wie in der ersten Challenge.
Es muss `change_me` mit einem beliebigen Wert überschrieben werden.

Nur dieses mal werden die Eingeben überprüft, damit sie nicht länger als der Buffer sind.
Nur leider wird dann der Wert an `sprintf` übergeben.

Die Challenge scheint jetzt vielleicht wieder bisschen zu einfach, aber das ist nur der Anfang.
Und ich persönlich finde format string vulnerabilities
generell etwas schwieriger zu verstehen und zu handhaben als einen einfachen Buffer-Overflow.
Daher spiele am besten ein bisschen mit der Challenge rum und überlege,
was man damit noch so anstellen könnte.

```c
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

  if (locals.change_me == 0xdeadbeef) {
    win();
  } else if (locals.change_me != 0) {
      printf("Getting closer! change_me is currently 0x%08x, but we want 0xdeadbeef\n",
              locals.change_me);
  } else {
    puts("Uh oh, 'change_me' has not yet been changed.");
  }

  exit(0);
}
```

Viel Spaß beim Experimentieren! 🧪
