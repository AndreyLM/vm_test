host=localhost
port=8000
json-file=/home/andrew/Projects/go-modules/vm_test/cmd/server/users.json
use-db=true

run:
	go run ./cmd/server/server.go -host=${host} -port=${port} -json-file=${json-file} -use-db=${use-db} 

build: 
	go build ./cmd/server/server.go