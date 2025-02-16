
migrate-up:
	goose -dir migrations postgres "postgres://postgres:1@localhost:5432/sungurTravel" up


migrate-down:
	goose -dir migrations postgres "postgres://postgres:1@localhost:5432/sungurTravel" down