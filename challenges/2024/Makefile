CHALLENGES := $(wildcard */.)

all: $(CHALLENGES)

$(CHALLENGES):
	$(MAKE) -C $@

dockerfiles:
	for challenge in $(CHALLENGES); do \
		$(MAKE) -C $$challenge dockerfile; \
	done

challenge_files:
	for challenge in $(CHALLENGES); do \
		$(MAKE) -C $$challenge archive; \
	done

.PHONY: all $(CHALLENGES)
