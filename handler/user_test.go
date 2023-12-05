package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Abdul-Rahim-K-T/pipeline/database"
	"github.com/Abdul-Rahim-K-T/pipeline/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

//Test User Get Api

func TestGetUser(t*testing.T){
	a:=assert.New(t)
	db:=database.InitDB()
	user,err:=InsertUser(db)
	if err!=nil{
		a.Error(err)
	}
	req,w:=SetUserGetRouter("/1")
	a.Equal(http.MethodGet,req.Method,"HTTP request method error")
	a.Equal(http.StatusOK,w.Code,"HTTP request status code error")
	
	body,err:=io.ReadAll(w.Body)
	if err!=nil{
		a.Error(err)
	}
	actual:=models.User{}
	if err:=json.Unmarshal(body,&actual);err!=nil{
		a.Error(err)
	}
	actual.Model=gorm.Model{}
	expected:=user
	expected.Model=gorm.Model{}
	a.Equal(expected,actual)
	cleanupDatabase(t)
}




func SetUserGetRouter(url string)(*http.Request,*httptest.ResponseRecorder){
	r:=gin.New()
	r.GET("/:id",GetUserById)

	// req,err:=http.NewRequest(http.MethodGet,url,nil)
	// if err!=nil{
	// 	panic(err)
	// }
	
	
	req,err:=http.NewRequest(http.MethodGet,url,nil)
	if err!=nil{
		panic(err)
	}

	req.Header.Set("Content-Type","application/json")
	w:=httptest.NewRecorder()
	r.ServeHTTP(w,req)
	return req,w
}

func InsertUser(db *gorm.DB)(models.User,error){
user:=models.User{
	First_Name: "Vaishakh",
	Last_Name: "Sukumar",
	Email: "Vaishakh@gmail.com",
	Password: "12345",
}
if err:=db.Create(&user).Error;err!=nil{
return user,err
}
return user,nil
}

//  test user create api
func TestCreateUser(t *testing.T){
	a:=assert.New(t)

	user:=models.User{
		First_Name: "Vaishakh",
		Last_Name: "Sukumar",
		Email: "Vaishakh@gmail.com",
		Password: "12345",
	}
	reqBody,err:=json.Marshal(user)
	if err!=nil{
		a.Error(err)
	}
	req,w,err:=SetCreateUserRout(bytes.NewBuffer(reqBody))
	if err!=nil{
		a.Error(err)
	}

	a.Equal(http.MethodPost,req.Method,"HTTP request methode error")
	a.Equal(http.StatusOK,w.Code,"HTTP request status code error")
	body,err:=io.ReadAll(w.Body)
	if err!=nil{
		a.Error(err)
	}
	actual:=models.User{}
	if err:=json.Unmarshal(body,&actual);err!=nil{
		a.Error(err)
	}
	actual=models.User{}
	if err:=json.Unmarshal(body,&actual);err!=nil{
		a.Error(err)
	}
	actual.Model=gorm.Model{}
	expected:=user
	expected.Model=gorm.Model{}
	a.Equal(expected,actual)
	cleanupDatabase(t)
}

func SetCreateUserRout(body *bytes.Buffer)(*http.Request,*httptest.ResponseRecorder,error){
	router:=gin.New()

	router.POST("/",CreateUser)
	req,err:=http.NewRequest(http.MethodPost,"/",body)
	if err!=nil{
		return req,httptest.NewRecorder(),err
	}
	req.Header.Set("Content-Type","application/json")
	w:=httptest.NewRecorder()
	router.ServeHTTP(w,req)
	return req,w,nil
}


// test user update api

func TestUpdateUser(t *testing.T){
	a:=assert.New(t)

	userBody:=models.User{
		First_Name: "rahim",
		Last_Name: "kt",
		Email: "rahimkt12@gmail.com",
		Password: "12345",
	}
	reqbody,err:=json.Marshal(userBody)
	if err!=nil{
		a.Error(err)
	}

	db:=database.InitDB()
	user,err:=InsertUser(db)
	if err!=nil{
		a.Error(err)
	}
	fmt.Println("Create Body:",user)

	if err!=nil{
		a.Error(err)
	}

	req,w,err:=SetUpdateUserRouter("/1",
bytes.NewBuffer(reqbody))
if err!=nil{
	a.Error(err)
}

a.Equal(http.MethodPut,req.Method,"HTTP request methode error")
a.Equal(http.StatusOK,w.Code,"HTTP request status code error")
body,err:=io.ReadAll(w.Body)
if err!=nil{
	a.Error(err)
}
actual:=models.User{}
if err:=json.Unmarshal(body,&actual);err!=nil{
	a.Error(err)
}
actual.Model=gorm.Model{}
expected:=userBody
a.Equal(expected,actual)
cleanupDatabase(t)
}


func SetUpdateUserRouter(url string,body *bytes.Buffer)(*http.Request,*httptest.ResponseRecorder,error){
	router:=gin.New()
	router.PUT("/:id",UpdateUserById)
	req,err:=http.NewRequest(http.MethodPut,url,body)
	if err!=nil{
		panic(err)
	}
	req.Header.Set("Content-Type","application/json")
	w:=httptest.NewRecorder()
	router.ServeHTTP(w,req)
	return req,w,err
}


func cleanupDatabase(t *testing.T){
	//  open a database  connection (use the same connection settings as in your tests)
	db:=database.InitDB()

	if err:=db.Exec("TRUNCATE TABLE  USERS  CASCADE").Error;err!=nil{
		t.Errorf("Error truncating the 'USERS' table:%v",err)
	}
	if err:=db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1").Error;err!=nil{
		t.Errorf("Error resetting the ID sequence:%v",err)
	}
}