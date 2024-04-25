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
go mod tidy
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

---


## Endpoints  for JSON
#### [![](https://img.shields.io/badge/-GET-mediumgreen?style=flat&logo=GET&logoColor=white)]()  /api/v1/invention - return all objects
#### [![](https://img.shields.io/badge/-GET-mediumgreen?style=flat&logo=GET&logoColor=white)]()  /api/v1/invention/:id - return object by ID
#### [![](https://img.shields.io/badge/-POST-orange?style=flat&logo=GET&logoColor=white)]()  /api/v1/invention/ - add object

---
## JSON Example
Method [![](https://img.shields.io/badge/-POST-orange?style=flat&logo=GET&logoColor=white)]()
```
    {
        "id": "19",
        "title": "Exclusive context-sensitive parallelism",
        "date": "{{currentDate}}",
        "description": "generate virtual architectures",
        "inventor": {
            "firstname":"Rachel",
            "lastname":"Smith",
            "age": 23,
            "email":"rachel.smith@mail.com"
        }
    }
```
## Endpoints  for view

#### [![](https://img.shields.io/badge/-GET-mediumgreen?style=flat&logo=GET&logoColor=white)]()  /view/v1/invention - return all objects in HTML

![alt text](https://i.imgur.com/kZoOCty.png)

## License
[MIT license](https://choosealicense.com/licenses/mit/)
