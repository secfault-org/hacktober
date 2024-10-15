#!/bin/sh

DISABLE_ASLR=0
WRITE_FLAG=0
PORT=1337
BIN_PATH="/chall"
ARCH=x86_64

while getopts "Dwb:p:a:" opt; do
  case $opt in
    D)
      DISABLE_ASLR=1
      ;;
    w)
      WRITE_FLAG=1
      ;;
    b)
      BIN_PATH=$OPTARG
      ;;
    p)
      PORT=$OPTARG
      ;;
    a)
      ARCH=$OPTARG
      ;;
    \?)
      echo "Invalid option: -$OPTARG" >&2
      ;;
  esac
done

shift $((OPTIND-1))

if [ $WRITE_FLAG -eq 1 ]; then
  echo "$FLAG" > /flag
fi

CMD="socat TCP-LISTEN:$PORT,reuseaddr,fork EXEC:$BIN_PATH"

if [ $DISABLE_ASLR -eq 1 ]; then
  CMD="setarch $ARCH -R $CMD"
fi

$CMD
