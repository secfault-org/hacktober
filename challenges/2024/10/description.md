Challenge 10
===========

Das Schöne an einer Format-String vulnerability ist, dass man gezielt Werte überschreiben kann.
Damit muss man nicht alles bis zu diesem Wert überschreiben, oder kann auch in ganze andere
Speicherbereiche gehen.

Diese Challenge wagt mal einen extras größeren Schritt. Ursprünglich wollte ich diese Challenge
in drei kleine Challenges aufteilen. Also keine Sorte, wenn du dafür etwas länger brauchst.

Deswegen hier noch ein paar Worte zu der Challenge.

1. Die Address von `change_me` wird für dich schon ausgegeben.
Dadurch, dass hier ASLR und PIE aktiv ist, wird sich die Adresse jedes Mal, wenn die Program gestartet wird,
an eine andere Adresse sein. Das ist heutzutage der Standard und oft ist es die erste Herausforderung,
die überwunden werden muss. Dafür wird oft das Leaken von infos über eine Format-String Vulnerability
verwendet. Der Schritt wird dir hier geschenkt.
Dennoch finde ich es wichtig, das mal einzuüben, da es so häufig vorkommt.
Die ausgegebene Adresse wird also für den Payload dann benötigt.

2. Wenn du die Adresse kennst, musst du schauen, dass sie irgendwie im Stack steht, damit du es dann
als Ziel für deine Format-String Angriff benutzten kannst.

3. Der dritte Tipp ergibt vielleicht erst Sinn, wenn du vor dem Problem stehst,
aber versuche am besten nicht gleich alles auf einem zu schreiben. 


```c
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
```

Viel Erfolg und versuche nicht zu verzweifeln! 🍀