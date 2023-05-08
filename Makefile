dev:
	docker-compose up -d --build
dev-down:
	docker-compose down
install-modules:
	go get -u github.com/gin-gonic/gin
	go get go.mongodb.org/mongo-driver/mongo
	go get github.com/spf13/viper
	go get github.com/golang-jwt/jwt
	# Hot-reload server
	go install github.com/cosmtrek/air@latest