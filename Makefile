KUBE_NS := react-golang-user-login-oauth
ENV_FILE := server/.env.prod

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
	kubectl -n $(KUBE_NS) apply -f k8s/server.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/client.yaml
	kubectl -n $(KUBE_NS) apply -f k8s/ingress.yaml
	$(MAKE) k8s-wait

# First-time deploy alias (same as apply-all)
k8s-deploy: k8s-apply

# Subsequent release (image update): assumes manifests unchanged
k8s-release:
	kubectl -n $(KUBE_NS) rollout restart deploy/user-login-api
	kubectl -n $(KUBE_NS) rollout restart deploy/user-login-client
	$(MAKE) k8s-wait

k8s-wait:
	kubectl -n $(KUBE_NS) rollout status statefulset/postgres --timeout=180s
	kubectl -n $(KUBE_NS) rollout status deploy/user-login-api --timeout=180s
	kubectl -n $(KUBE_NS) rollout status deploy/user-login-client --timeout=180s

k8s-status:
	kubectl -n $(KUBE_NS) get all
	kubectl -n $(KUBE_NS) get ingress,certificate,pvc

k8s-describe:
	kubectl -n $(KUBE_NS) describe deploy/user-login-api
	kubectl -n $(KUBE_NS) describe deploy/user-login-client
	kubectl -n $(KUBE_NS) describe statefulset/postgres
	kubectl -n $(KUBE_NS) describe ingress/oauth

# DANGEROUS: wipes DB data (PVC)
k8s-db-reset:
	kubectl -n $(KUBE_NS) delete statefulset postgres --ignore-not-found
	kubectl -n $(KUBE_NS) delete pvc data-postgres-0 --ignore-not-found
