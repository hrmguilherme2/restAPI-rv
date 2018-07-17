<snippet>
  <content><![CDATA[
# ${1:Project Name}
REST API with GO
## Installation

Step 1:
- Install Golang https://golang.org/dl/
- Dependencies
```Golang
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
```
Go to terminal and get those frameworks
1. `go get github.com/dgrijalva/jwt-go`
2. `go get github.com/gin-gonic/gin`
3. `go get github.com/go-sql-driver/mysql`


Step 2:
- Install PostForm to access API
  https://www.getpostman.com/apps

Step 3:
-  Install MySQL WorkBench https://www.mysql.com/products/workbench/<br />
-  Install MySQL Server https://dev.mysql.com/downloads/mysql/<br />

- Setup to configure the database

Execute the following Query:<br />
- Creating the Schema
```SQL
CREATE SCHEMA `redventures` ;
  ```

- Creating the tables <br />

Users table
```SQL
CREATE TABLE `redventures`.`users` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `gravatar` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));
  ```
Widgets table
  ```SQL
  CREATE TABLE `redventures`.`widgets` (
  `id` INT NOT NULL,
  `name` VARCHAR(45) NULL,
  `color` VARCHAR(45) NULL,
  `price` DOUBLE NULL,
  `inventory` INT NULL,
  `melts` VARCHAR(10) NULL,
  PRIMARY KEY (`id`));
  ```

Credentials Table
  ```SQL
CREATE TABLE `redventures`.`credentials` (
  `id` INT NOT NULL,
  `username` VARCHAR(45) NULL,
  `password` VARCHAR(45) NULL,
  PRIMARY KEY (`id`));
  ```

Step 4:
- Filling the tables<br/>

USERS<br/>
```SQL
INSERT INTO `redventures`.`users` (`id`, `name`, `gravatar`) VALUES ('1', 'Colin', 'http://www.gravatar.com/avatar/a51972ea936bc3b841350caef34ea47e?s=64&d=monsterid');
INSERT INTO `redventures`.`users` (`id`, `name`, `gravatar`) VALUES ('2', 'Guilherme', 'http://www.gravatar.com/avatar/432f3e353c689fc37af86ae861d934f9?s=64&d=monsterid');
INSERT INTO `redventures`.`users` (`id`, `name`, `gravatar`) VALUES ('3', 'Thomas', 'http://www.gravatar.com/avatar/48009c6a27d25ef0ea03f985d1f186b0?s=64&d=monsterid');
INSERT INTO `redventures`.`users` (`id`, `name`, `gravatar`) VALUES ('4', 'James', 'http://www.gravatar.com/avatar/9372f138140c8578c82bbc77b2eca602?s=64&d=monsterid');
  ```
  
WIDGETS<br>
```SQL
INSERT INTO `redventures`.`widgets` (`id`, `name`, `color`, `price`, `inventory`, `melts`) VALUES (1, 'Losenoidenbdsfkv', 'Blue', 10, 1003, true); 
INSERT INTO `redventures`.`widgets` (`id`, `name`, `color`, `price`, `inventory`, `melts`) VALUES (2, 'Guilherme', 'Blue', 10, 1003, true); 
```


CREDENTIALS<br>
```SQL
INSERT INTO `redventures`.`credentials` (`id`, `username`, `password`, `name`) VALUES (1, 'rv', 'rv', 'redventures');
```




## Usage
Go to golang terminal and run the project (go run *.go).<br>
Start PostMan Software and use it to get the responses.<br>
- Getting the Token
`http://localhost:3000/login?user={your username}&pass={your password}`
![Alt Text](https://media.giphy.com/media/YllxobwKm3GBXXOEip/giphy.gif)

- Validating the token
`http://localhost:3000/validar?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwcm9maWxlIjp7Im5hbWUiOiJndWlsaGVybWUiLCJwZXJtaXNzaW9ucyI6WyJhZG1pbiJdfSwiaXNzIjoidGVzdC1yZXN0LWFwaS1ydiJ9.G1EJ9byLgAg74lRgjYP3gbvfKuIlzDnF1xOPzBojJPQ `
![Alt Text](https://media.giphy.com/media/kv5YRt0L0livTaXkiF/giphy.gif)


After the token validated you are enabled to access the full restAPI.



GET /users http://localhost:3000/users<br>
![Alt Text](https://media.giphy.com/media/nElFFnF7qoW8Bn6W4m/giphy.gif)

GET /users/:id http://localhost:3000/:id<br>
![Alt Text](https://media.giphy.com/media/pjm7gXQMw2R2TUMbJm/giphy.gif)

GET /widgets http://localhost:3000/widgets<br>
![Alt Text](https://media.giphy.com/media/XZKLQp0x1xoK8UrKVt/giphy.gif)

GET http://localhost:3000/widgets/:id<br>
![Alt Text](https://media.giphy.com/media/42vdMODrsbMOWjhTsj/giphy.gif)

POST /widgets for creating new widgets http://localhost:3000/widgets<br>
Syntax: `http://localhost:3000/widget?id={id}&name=Guilherme&color=Azul&price=1&inventory=1002&melts=true`
![Alt Text](https://media.giphy.com/media/2Yc392e5BI5hV5ZYoE/giphy.gif)

PUT /widgets/:id for updating existing http://localhost:3000/widgets/:id<br>
Syntax : ` http://localhost:3000/widget?name=GUILHERME&color=BLACK&price=10&inventory=1001&melts=true&id={id to updated}`
![Alt Text](https://media.giphy.com/media/bqC3eDPwvpz0OI3nGL/giphy.gif)




</content>
  <tabTrigger>Thanks</tabTrigger>
</snippet>
