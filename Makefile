# Makefile

# Path to sqlc config
SQLC_CONFIG=sqlc.yaml

.PHONY: sqlc
sqlc:
	sqlc generate -f $(SQLC_CONFIG)
