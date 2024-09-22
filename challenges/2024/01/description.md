Die erste Challenge ist hervorragend geeignet, um sich mit dem Stack und den Basics von Buffer Overflows vertraut zu machen.

Außerdem empfehle ich, sich mit [GDB](https://www.gnu.org/software/gdb/) (mit **gef** oder **pwndbg**) und [Pwntools](https://github.com/Gallopsled/pwntools) vertraut zu machen.
Auch wenn es jetzt noch nicht notwendig ist, wird es in den späteren Challenges sehr hilfreich sein.

Der Code ist sehr einfach gehalten und es gibt nur eine einzige Schwachstelle, die es zu finden gilt.

```c
char *gets(char *);

void printFlag() {
  char flag[64];
  FILE *f = fopen("flag.txt", "r");
  if (f == NULL) {
    puts(
        "Flag file not found. Make sure to create a dummy"
        "flag.txt file in the same directory as the challenge binary."
    );
    exit(1);
  }
  fgets(flag, sizeof(flag), f);
  printf("Flag: %s\n", flag);
}

int main(int argc, char **argv) {
  struct {
    char buffer[64];
    volatile int changeme;
  } locals;

  printf("%s\n", BANNER);

  locals.changeme = 0;
  gets(locals.buffer);

  if (locals.changeme != 0) {
    printFlag();
  } else {
    puts("Uh oh, 'changeme' has not yet been changed.");
  }

  exit(0);
}
```