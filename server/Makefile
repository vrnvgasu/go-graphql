psql-run:
	@docker-compose up -d
	@sleep 5
	@~/go/bin/goose --dir ./migrations "postgres" "postgres://module12_task05:module12_task05@localhost:5432/module12_task05?sslmode=disable&connect_timeout=5" up
