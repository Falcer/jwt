buildserver:
	cd server && go build -o build/server main.go
runserver:
	cd server && ./build/server

buildweb:
	cd web && yarn build && yarn export
runweb:
	cd web && yarn dev