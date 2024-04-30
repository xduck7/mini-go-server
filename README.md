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
swag init -g cmd/app/main.go
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

---


## Endpoints
#### [![](https://img.shields.io/badge/-GET-mediumgreen?style=flat&logo=GET&logoColor=white)]()  /api/v1/invention - return all objects
#### [![](https://img.shields.io/badge/-GET-mediumgreen?style=flat&logo=GET&logoColor=white)]()  /api/v1/invention/:id - return object by ID
#### [![](https://img.shields.io/badge/-POST-orange?style=flat&logo=GET&logoColor=white)]()  /api/v1/invention/ - add object

---
## JSON Example with Postman
 [![](https://img.shields.io/badge/-POST-orange?style=flat&logo=GET&logoColor=white)]() ```/api/v1/invention/```
 
 Body
```
{
    "id": "{{randId}}",
    "title": "Test",
    "date": "{{currentDate}}",
    "description": "Test",
    "inventor": {
        "firstname": "Test",
        "lastname": "Test",
        "age": 43,
        "email": "test@mail.com"
    }
}
```
Pre-request Script
```
let currentDate = new Date().toISOString();
pm.environment.set("currentDate", currentDate);

function generateID() {
    return Math.floor(Math.random() * (1000)) + 1;
}

let randId = generateID().toString();
pm.environment.set("randId", randId);

```
## View

* ### `````/menu````` - menu for adding

![alt text](https://i.imgur.com/bcm17mK.png)


* ### `````/inventions````` - menu for displaying

![alt text](https://i.imgur.com/hp3oyVL.png)

* ### `````/swagger/index.html````` - Swagger UI



## License
[MIT license](https://choosealicense.com/licenses/mit/)
