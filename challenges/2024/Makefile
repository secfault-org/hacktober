CHALLENGES := $(wildcard */.)

all: $(CHALLENGES)

$(CHALLENGES):
	$(MAKE) -C $@

dockerfiles: $(CHALLENGES)
	for challenge in $(CHALLENGES); do \
		$(MAKE) -C $$challenge dockerfile; \
	done

.PHONY: all $(CHALLENGES)
