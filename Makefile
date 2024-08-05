run_all:
	rm server/log.txt ; server/server > server/log.txt 2>&1 &
	rm client/log.txt ; client/client > client/log.txt 2>&1 &
	rm chat/log.txt ; chat/chat > chat/log.txt 2>&1 &

run_server:
	rm server/log.txt ; server/server > server/log.txt 2>&1 &
log_server:
	less server/log.txt

run_client:
	rm client/log.txt ;sudo client/client > client/log.txt 2>&1 &
log_client:
	less client/log.txt

run_chat:
	rm chat/log.txt ; chat/chat > chat/log.txt 2>&1 &
log_client:
	less chat/log.txt

go_build:
	env GOOS=linux go build -o server