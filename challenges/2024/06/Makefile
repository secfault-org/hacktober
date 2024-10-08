TARGET_NAME=chall
DOCKER_NAME=hacktober2024_challenge_06
SRCS=chall.c

CC=gcc
CFLAGS=-Wall -Wextra -Werror -Wno-deprecated-declarations -O0 -std=c99 -pedantic -fno-strict-aliasing -fno-omit-frame-pointer -fno-stack-protector -fno-stack-check -fno-pie -no-pie -z execstack -m32
LDFLAGS=-no-pie -z execstack

all: $(TARGET_NAME)

$(TARGET_NAME): $(SRCS)
	$(CC) $(CFLAGS) $(LDFLAGS) -o $@ $^

run: $(TARGET_NAME)
	FLAG=this_is_only_a_test_flag ./$(TARGET_NAME)

dockerfile: $(TARGET_NAME)
	docker build -t $(DOCKER_NAME) .

clean:
	rm -f $(TARGET_NAME)
