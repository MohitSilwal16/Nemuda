run_server_chat:
	cd server && rm log.txt && ./server > log.txt 2>&1 & cd ..
	cd chat && rm log.txt && ./chat > log.txt 2>&1 & cd ..

run:
	cd server && rm log.txt && ./server > log.txt 2>&1 & cd ..
	cd client && rm log.txt && sudo ./client > log.txt 2>&1 & cd ..
	cd chat && rm log.txt && ./chat > log.txt 2>&1 & cd ..

kill:
	pkill server && pkill chat && sudo pkill client

log:
	tail -n 25 server/log.txt && tail -n 25 client/log.txt && tail -n 25 chat/log.txt

run_server:
	cd server && rm log.txt && ./server > log.txt 2>&1 & cd ..
kill_server:
	pkill server
log_server:
	tail -n 25 server/log.txt

run_client:
	cd client && rm log.txt && sudo ./client > log.txt 2>&1 & cd ..
kill_client:
	sudo pkill client
log_client:
	tail -n 25 client/log.txt

run_chat:
	cd chat && rm log.txt && ./chat > log.txt 2>&1 & cd ..
kill_chat:
	pkill chat
log_chat:
	tail -n 25 chat/log.txt
