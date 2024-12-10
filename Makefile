GO := go

DAY := day$(day)

.PHONY: sample
sample:
ifeq ($(day),)
	@echo "Please specify a day using 'make run day=XX'"
else
	@echo "Running day $(day) with sample input..."
	sh -c "time $(GO) run $(DAY)/main.go $(DAY)/sample"
	@echo ""
endif

.PHONY: input
input:
ifeq ($(day),)
	@echo "Please specify a day using 'make run-input day=XX'"
else
	@echo "Running day $(day) with full input..."
	sh -c "time $(GO) run $(DAY)/main.go $(DAY)/input"
	@echo ""
endif

# Catch-all help target
.PHONY: help
help:
	@echo "Usage:"
	@echo "  make (sample|input) day=XX      Run the main.go file for dayXX with either the example or full input"
