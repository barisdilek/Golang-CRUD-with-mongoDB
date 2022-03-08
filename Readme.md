#CRUD

1. go package install
2. Restart IDE
3. touch main.go

# gin-gonic ve diğer modüllerin yönetimi için

# mod uzantılı bir dosya oluşacaktır. Burada yüklediğimiz paket bilgilerini görebiliriz.

4. go mod init <folderName>

# gin-gonic için gerekli go paketlerinin yüklenmesi

5. go get -u github.com/gin-gonic/gin

# mongodb için gerekli go paketlerinin yüklenmesi

6. go get -u go.mongodb.org/mongo-driver

# MongoDB tarafıyla eşlecek entity için

7. touch userDmo.go

# CRUD Operasyonları için

8. touch userRepository.go

# Servis metotlarındaki annotation bildirimlerinin Swagger 2.0 destekli olarak dokümante edilmesi için gerekli modülün eklenmesi

9. go get -u github.com/swaggo/swag/cmd/swag

# Bu arada kod tarafındaki annotation bölümleri tamamlandıktan sonra Swagger dokümanının üretilmesi için aşağıdaki komutu çalıştırmalıyız

10. swag init \_ "<folderName>/docs"

# önce bir build etmek lazım

11. go build

# sonrasında çalıştırabiliriz

12 ./<folderName>
Sample $ ./crud

# Örnek birkaç alıntı girelim

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

# Şimdi de listeleme yapalım

curl http://localhost:5003/api/v1/user/

#go run main.go
