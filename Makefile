SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --no-print-directory

ENV_PAHT=.env

ifneq ($(wildcard $(ENV_PAHT)),)
	include $(ENV_PAHT)
	export
endif

.PHONY: setup tf-plan tf-deploy

setup:
	go install github.com/google/ko@latest

tf-init:
	@./scripts/copy_tf.sh
	@cd tmp/enviroments/dev; \
	terraform init; \
	terraform validate;

tf-plan:
	@cd tmp/enviroments/dev; \
	terraform plan

tf-deploy:
	@cd tmp/enviroments/dev; \
	terraform apply -auto-approve
