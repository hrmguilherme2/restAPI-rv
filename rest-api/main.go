package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var tokIsValid = false

var signingKey = []byte("signing-key")

func signingKeyFn(*jwt.Token) (interface{}, error) {
	return signingKey, nil
}

func main() {
	//Conexao com mySQL
	db, err := sql.Open("mysql", "sa:12345@tcp(127.0.0.1:3306)/redventures")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// ter certeza que o servidor conecta
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}

	router := gin.Default()

	// GET login token
	router.GET("/login", func(c *gin.Context) {

		username := c.Query("user")
		password := c.Query("pass")

		rows, err := db.Query("select username, password,name from credentials where username = ?", username)

		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			var user string
			var pass string
			var name string
			if err := rows.Scan(&user, &pass, &name); err != nil {
				log.Fatal(err)
			}

			if username == user && password == pass {
				// c.JSON(http.StatusOK, gin.H{
				// 	"user parametro":  username,
				// 	"senha parametro": password,
				// 	"senha db":        pass,
				// 	"user db":         user,
				// })

				claims := UserClaims{
					UserProfile{Name: name, Permissions: []string{"admin"}},
					jwt.StandardClaims{
						Issuer: "test-rest-api-rv",
					},
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				ss, err := token.SignedString(signingKey)
				if err != nil {
					log.Printf("err: %+v\n", err)
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"token gerado": ss,
				})
			} else {

				c.JSON(http.StatusOK, gin.H{
					"log": "erro de autenticação",
				})
			}

		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

	})

	// GET valid token
	router.GET("/validar", func(c *gin.Context) {
		var claims UserClaims
		//token model "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJwcm9maWxlIjp7Im5hbWUiOiJKYW1lcyBTbWl0aCIsInBlcm1pc3Npb25zIjpbImRvU3R1ZmYiXX0sImlzcyI6InRlc3QtcHJvamVjdCJ9.ikG0n-dLC-AYZw6dNbULkdW4u8Ctq9k1VfH3YK8WX7A"
		tokenn := c.Query("token")
		token, err := jwt.ParseWithClaims(tokenn, &claims, signingKeyFn)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"log": "falha ao validar token",
			})
			log.Println("Token falhado")
			return
		}

		if !token.Valid {
			c.JSON(http.StatusOK, gin.H{
				"log": "false",
			})
			log.Println("Invalido token")
			return
		} else if token.Valid {

			tokIsValid = true
		}

		claimsString := fmt.Sprintf("claims: %v", claims)

		log.Println(claimsString)
		//Se o token for válido retorno é true
		c.JSON(http.StatusOK, gin.H{
			"log": "true",
		})

	})

	//get all users
	router.GET("/users", func(c *gin.Context) {
		if tokIsValid == true {
			var (
				usr   _User
				users []_User
			)
			rows, err := db.Query("select id, name, gravatar from users;")
			if err != nil {
				fmt.Print(err.Error())
			}
			for rows.Next() {
				err = rows.Scan(&usr.Id, &usr.Name, &usr.Gravatar)
				users = append(users, usr)
				if err != nil {
					fmt.Print(err.Error())
				}
			}
			defer rows.Close()
			c.JSON(http.StatusOK, gin.H{
				"Users": users,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"log": "sem autorização",
			})
		}

	})

	// GET a user detail
	router.GET("/users/:id", func(c *gin.Context) {
		if tokIsValid == true {
			var (
				usr    _User
				result gin.H
			)
			id := c.Param("id")
			row := db.QueryRow("select id, name, gravatar from users where id = ?;", id)
			err = row.Scan(&usr.Id, &usr.Name, &usr.Gravatar)
			if err != nil {
				// If no results send null
				result = gin.H{
					"result": nil,
					"count":  0,
				}
			} else {
				result = gin.H{
					"result": usr,
					"count":  1,
				}
			}
			c.JSON(http.StatusOK, result)

		} else {
			c.JSON(http.StatusOK, gin.H{
				"log": "sem autorização",
			})

		}

	})
	//get widgets all
	router.GET("/widgets", func(c *gin.Context) {
		if tokIsValid == true {
			var (
				wdt  _Widgets
				wdts []_Widgets
			)
			rows, err := db.Query("select * from widgets;")
			if err != nil {
				fmt.Print(err.Error())
			}
			for rows.Next() {
				err = rows.Scan(&wdt.ID, &wdt.Name, &wdt.Color, &wdt.Price, &wdt.Inventory, &wdt.Melts)
				wdts = append(wdts, wdt)
				if err != nil {
					fmt.Print(err.Error())
				}
			}
			defer rows.Close()
			c.JSON(http.StatusOK, gin.H{
				"Widgets": wdts,
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"log": "sem autorização",
			})
		}

	})
	//get widgets detail
	router.GET("/widgets/:id", func(c *gin.Context) {
		if tokIsValid == true {
			var (
				wdt    _Widgets
				result gin.H
			)
			id := c.Param("id")
			row := db.QueryRow("select * from widgets where id = ?;", id)
			err = row.Scan(&wdt.ID, &wdt.Name, &wdt.Color, &wdt.Price, &wdt.Inventory, &wdt.Melts)
			if err != nil {
				// If no results send null
				result = gin.H{
					"result": nil,
					"count":  0,
				}
			} else {
				result = gin.H{
					"result": wdt,
					"count":  1,
				}
			}
			c.JSON(http.StatusOK, result)

		} else {
			c.JSON(http.StatusOK, gin.H{
				"log": "sem autorização",
			})

		}

	})

	router.POST("/widgets", func(c *gin.Context) {
		if tokIsValid == true {
			id := c.DefaultQuery("id", "Guest")
			name := c.Query("name")
			color := c.Query("color")
			price := c.Query("price")
			inventory := c.Query("inventory")
			melts := c.Query("melts")

			stmt, err := db.Prepare("insert into widgets (id,name,color,price,inventory,melts) values(?,?,?,?,?,?);")
			if err != nil {
				fmt.Print(err.Error())
			}
			_, err = stmt.Exec(id, name, color, price, inventory, melts)

			if err != nil {
				fmt.Print(err.Error())
			}

			defer stmt.Close()

			c.JSON(http.StatusOK, gin.H{
				"log": fmt.Sprintf(" %s widget inserido com sucesso", name),
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"log": "sem autorização",
			})

		}

	})

	router.PUT("/widgets", func(c *gin.Context) {
		if tokIsValid == true {
			id := c.DefaultQuery("id", "Guest")
			name := c.Query("name")
			color := c.Query("color")
			price := c.Query("price")
			inventory := c.Query("inventory")
			melts := c.Query("melts")

			stmt, err := db.Prepare("update widgets set name =?,color=?,price=?,inventory=?,melts=? where id=?;")
			if err != nil {
				fmt.Print(err.Error())
			}
			_, err = stmt.Exec(name, color, price, inventory, melts, id)

			if err != nil {
				fmt.Print(err.Error())
			}

			defer stmt.Close()

			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("%s widgets atualizado", name),
			})

		} else {
			c.JSON(http.StatusOK, gin.H{
				"log": "sem autorização",
			})
		}

	})

	//Usando a porta 3000
	router.Run(":3000")

	log.Fatalln(http.ListenAndServe(":3000", nil))

}
