module backend

go 1.14

replace backend => /home/suchil/workspace/famg/back

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/handlers v1.4.2
	github.com/gorilla/mux v1.7.3
	github.com/gorilla/securecookie v1.1.1
	github.com/jinzhu/gorm v1.9.12
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20200208060501-ecb85df21340
)
