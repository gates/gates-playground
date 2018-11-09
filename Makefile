REPO = github.com/lujjjh/gates-playground

build:
	gopherjs build $(REPO)

.PHONY: build

run:
	gopherjs serve $(REPO)

.PHONY: run
