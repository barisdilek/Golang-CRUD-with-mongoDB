#CRUD

![alt text](https://github.com/barisdilek/Golang-CRUD-with-mongoDB/blob/main/crud.jpg?raw=true)

1. go package install
2. Restart IDE
3. touch main.go

# for management of gin-gonic and other modules

# A mod file will be created. Here we can see the package information we have installed.

4. go mod init <folderName>

# Installing the required go packages for gin-gonic

5. go get -u github.com/gin-gonic/gin

# Installing required go packages for mongodb

6. go get -u go.mongodb.org/mongo-driver

# For the entity to match the MongoDB side

7. touch userDmo.go

# For CRUD Operations

8. touch userRepository.go

# Adding the necessary module to document annotation declarations in service methods with Swagger 2.0 support

9. go get -u github.com/swaggo/swag/cmd/swag

# In the meantime, after the annotation sections on the code side are completed, we must run the following command to generate the Swagger document.

10. swag init \_ "<folderName>/docs"

# build

11. go build

# run

12 ./<folderName>
Sample $ ./crud

# Sample

curl --location --request POST 'http://localhost:5003/api/v1/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
"Name": "Baris",
"Surname": "Dilek",
"Pass": "pass",
"Email":"barisdilek@hotmail.com"
}'

curl --location --request POST 'http://localhost:5003/api/v1/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
"Name": "Deneme",
"Surname": "Test",
"Pass": "pass",
"Email":"deneme@test.com"
}'

curl --location --request POST 'http://localhost:5003/api/v1/user/' \
--header 'Content-Type: application/json' \
--data-raw '{
"Name": "Deneme 2",
"Surname": "Test",
"Pass": "pass",
"Email":"deneme2@test.com"
}'

# List All

curl http://localhost:5003/api/v1/user/

#go run main.go
