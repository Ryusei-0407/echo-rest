test:
	go test ./*/**_test.go

restart:
	docker stop echo
	docker compose up -d

run:
	docker compose up -d --build

down:
	docker compose down

pause:
	docker compose pause

unpause:
	docker compose unpause
