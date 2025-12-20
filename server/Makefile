# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go
# Create DB container
docker-run:
	@if docker compose up --build 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose up --build; \
	fi

# Shutdown DB container
docker-down:
	@if docker compose down 2>/dev/null; then \
		: ; \
	else \
		echo "Falling back to Docker Compose V1"; \
		docker-compose down; \
	fi

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v
# Integrations Tests for the application
itest:
	@echo "Running integration tests..."
	@go test ./internal/database -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f main

# Live Reload
watch:
	@if command -v air > /dev/null; then \
            air; \
            echo "Watching...";\
        else \
            read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
            if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
                go install github.com/air-verse/air@latest; \
                air; \
                echo "Watching...";\
            else \
                echo "You chose not to install air. Exiting..."; \
                exit 1; \
            fi; \
        fi

.PHONY: all build run test clean watch docker-run docker-down itest

KUBE_NS := react-golang-user-login-oauth
ENV_FILE := .env.prod

k8s-ns:
	kubectl create namespace $(KUBE_NS) --dry-run=client -o yaml | kubectl apply -f -

k8s-secret: k8s-ns
	kubectl -n $(KUBE_NS) create secret generic user-login-api-env \
		--from-env-file=$(ENV_FILE) \
		--dry-run=client -o yaml | kubectl apply -f -

# First-time deploy (includes DB + ingress)
k8s-apply: k8s-secret
	kubectl -n $(KUBE_NS) apply -f k8s/postgres-init.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/postgres.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/deploy.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/ingress.yaml
	$(MAKE) k8s-wait

# First-time deploy alias (same as apply-all)
k8s-deploy: k8s-apply

# Subsequent release (image update): assumes manifests unchanged
k8s-release:
	kubectl -n $(KUBE_NS) rollout restart deploy/user-login-api
	$(MAKE) k8s-wait

k8s-wait:
	kubectl -n $(KUBE_NS) rollout status statefulset/postgres --timeout=180s
	kubectl -n $(KUBE_NS) rollout status deploy/user-login-api --timeout=180s

k8s-status:
	kubectl -n $(KUBE_NS) get all
	kubectl -n $(KUBE_NS) get ingress,certificate,pvc

k8s-logs:
	kubectl -n $(KUBE_NS) logs -f deploy/user-login-api

k8s-logs-db:
	kubectl -n $(KUBE_NS) logs -f statefulset/postgres

k8s-describe:
	kubectl -n $(KUBE_NS) describe deploy/user-login-api
	kubectl -n $(KUBE_NS) describe statefulset/postgres
	kubectl -n $(KUBE_NS) describe ingress/user-login-api

# DANGEROUS: wipes DB data (PVC)
k8s-db-reset:
	kubectl -n $(KUBE_NS) delete statefulset postgres --ignore-not-found
	kubectl -n $(KUBE_NS) delete pvc data-postgres-0 --ignore-not-found
