ifeq ($(firstword $(MAKECMDGOALS)),run)
RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(if $(RUN_ARGS),$(eval $(RUN_ARGS):;@:))
endif

.PHONY: run
run:
	go run ./cmd/atproxy $(RUN_ARGS)
