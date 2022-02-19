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
  * [Auth(OAuth2)]()
  <details> 
    <summary>JWT,OAuth,OAuth2</summary>
    Firstly, we have to differentiate JWT and OAuth. Basically, JWT is a token format. OAuth is an authorization protocol that can use JWT as a token. OAuth uses server-side and client-side storage. If you want to do real logout you must go with OAuth2. Authentication with JWT token can not logout actually. Because you don't have an Authentication Server that keeps track of tokens. If you want to provide an API to 3rd party clients, you must use OAuth2 also. OAuth2 is very flexible. JWT implementation is very easy and does not take long to implement. If your application needs this sort of flexibility, you should go with OAuth2. But if you don't need this use-case scenario, implementing OAuth2 is a waste of time.
  </details>
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

