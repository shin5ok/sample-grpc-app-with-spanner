init:
	@echo "Prepareing environment of terraform to init"
	terraform init
	terraform providers lock -platform=darwin_amd64 -platform=linux_amd64

graph:
	terraform graph | dot -Tpng > graph.png
