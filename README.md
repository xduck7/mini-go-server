# Mini Go Server
Go web server on GIN

---
## Instalation

Instructions on how to get a copy of the project and running on your local machine.

```bash
git clone https://github.com/xduck7/mini-go-server.git
```

```bash
cd ./mini-go-server
```

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

```bash
swag init
```

```bash
go mod tidy
```

```bash
go run main.go -port=8080
```



---
## Usage

For running server on Windows use

```bash
mingw32-make run
```
For running server on Linux use

```bash
make run
``` 

For building server on Windows use
```bash
mingw32-make build-exe
``` 
For building server on Linux use
```bash
mingw32-make build-bin
``` 
---
## Technologies


* [Gin](https://github.com/gin-gonic/gin/) - Web Framework
* [Swagger](https://github.com/swaggo/swag) - Swagger UI
* [Gorm](https://gorm.io/) - Gorm.io
* [JWT](https://github.com/dgrijalva/jwt-go) - JWT-go
---


## View

* ### `````/menu````` - menu for adding

![alt text](https://i.imgur.com/bcm17mK.png)


* ### `````/inventions````` - menu for displaying

![alt text](https://i.imgur.com/hp3oyVL.png)

* ### `````/swagger/index.html````` - Swagger UI
![alt text](https://i.imgur.com/cK8hsr8.png)


## License
[MIT license](https://choosealicense.com/licenses/mit/)
