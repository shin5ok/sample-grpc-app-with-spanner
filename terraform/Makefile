.PHONY: init
init:
	@echo "Preparing environment of terraform to init"
	terraform init
	terraform providers lock \
		-platform=linux_amd64  \
		-platform=linux_arm64 \
		-platform=darwin_amd64 \
		-platform=windows_amd64 \
		-platform=darwin_arm64

.PHONY: graph
graph:
	terraform graph | dot -Tpng > graph.png
