generate:
	@$(MAKE) -C e2e/httptest generate

.PHONY: e2e
e2e:
	@echo "IMPORTANT: make sure you have run \`make generate\` before running the e2e\n"
	@$(MAKE) -C e2e/httptest e2e