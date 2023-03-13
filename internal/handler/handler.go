package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/AlexandrYar/simple_app/internal/db"
	"github.com/AlexandrYar/simple_app/pkg"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user db.User
	user.Login = c.PostForm("login")
	user.Password = c.PostForm("password")
	user.First_name = c.PostForm("first_name")
	user.Last_name = c.PostForm("last_name")
	user.Email = c.PostForm("email")
	user.Date_of_birth = c.PostForm("date_of_birth")
	user.Register(db.ConnDb.Connection(db.NewDb))

	c.HTML(http.StatusTemporaryRedirect, "register.html", gin.H{})

}

func Login(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")
	var user db.User
	message, isRedirectUserPage := user.LoginUser(db.ConnDb.Connection(db.NewDb), login, password)
	if isRedirectUserPage {
		c.Redirect(http.StatusTemporaryRedirect, "/userpage/"+login)
	}
	c.HTML(http.StatusTemporaryRedirect, "login.html", gin.H{
		"Message": message,
	})
}

func UserPage(c *gin.Context) {
	login_given := c.Params.ByName("login")
	var user db.User
	user.Find_info(db.ConnDb.Connection(db.NewDb), login_given)
	items := user.GetItems(db.NewDb.Connection())
	var lists []map[string]string
	for i := 0; i < len(items); i++ {
		mapOfItems := make(map[string]string)
		mapOfItems["Id"] = items[i].Id
		mapOfItems["Title"] = items[i].Title
		mapOfItems["Price"] = items[i].Price
		mapOfItems["Amount"] = items[i].Amount
		mapOfItems["PhotoUrl"] = items[i].PhotoUrl
		mapOfItems["SellerName"] = items[i].SellerName
		lists = append(lists, mapOfItems)
		log.Println(i, " - ", lists)
	}
	c.HTML(http.StatusOK, "userPage.html", gin.H{
		"Login":         user.Login,
		"First_name":    user.First_name,
		"Last_name":     user.Last_name,
		"Email":         user.Email,
		"Date_of_birth": user.Date_of_birth,
		"Items":         lists,
	})
}

func MainPage(c *gin.Context) {
	var items_list = []string{"а", "б", "в"}
	c.HTML(200, "main_page.html", gin.H{
		"Items": items_list,
	})
}

func AddItem(c *gin.Context) {
	c.Request.ParseMultipartForm(10 << 20)
	if c.Request.Method == "POST" {
		login := c.Params.ByName("login")
		files, err := ioutil.ReadDir("./usersFile")
		if err != nil {
			log.Fatal(err)
		}
		if pkg.IsFileExist(files, login) == false {
			err = os.Mkdir(fmt.Sprintf("./usersFile/"+login), 0777)
			if err != nil {
				panic(err)
			}
		}

		file, _, err := c.Request.FormFile("photo")
		if err != nil {
			log.Println(err)
		}
		out, err := os.Create(fmt.Sprintf("./usersFile/" + login + "/" + c.PostForm("id") + "_" + c.PostForm("title") + ".jpeg"))
		if err != nil {
			log.Println(err)
		}

		io.Copy(out, file)
		out.Close()
		var item = db.Item{
			Id:         c.PostForm("id"),
			Title:      c.PostForm("title"),
			Price:      c.PostForm("price"),
			Amount:     c.PostForm("amount"),
			PhotoUrl:   fmt.Sprintf(`usersFile/` + login + `/` + c.PostForm("id") + "_" + c.PostForm("title") + ".jpeg"),
			SellerName: login,
		}
		var user db.User
		user.AddNewItem(db.NewDb.Connection(), item)
	}

	c.HTML(200, "addItem.html", nil)
}
