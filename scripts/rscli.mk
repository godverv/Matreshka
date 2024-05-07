RSCLI_VERSION=v0.0.31
rscli-version:
	@echo $(RSCLI_VERSION)

buildc:
	docker build -t matreshka-be:local .