# graduationWork
## TODO
### [Tags]
- цвет текста зависит от цвета фона(hex) 

### [Database]
1. селектор бд по конфигу, что заполнено то и использовать, если несколько то выбор режима
    - дублирование
    - резервная копия
    - резервный сервер
    - использовать что-то одно)
  
### [endpoints]
- [ORM](https://github.com/go-reform/reform)  
- [Router](https://github.com/gin-gonic/gin)
- [Middleware](https://github.com/gin-gonic/contrib)
  * *[Auth](https://github.com/appleboy/gin-jwt)*
- [Swagger](https://github.com/swaggo/gin-swagger)
   * [Example main](https://github.com/swaggo/swag/blob/master/example/celler/main.go)
   * [Example func](https://github.com/swaggo/swag/blob/master/example/celler/controller/examples.go)

**General**
- /healthcheck  
---
**User**
- /login
- /register
- /logout
---
**Roles**
- /roles [Get] Get list of all roles
- /role?name= [Post] Create role
- /role
- */role?id= [Get] Get role with this id*

