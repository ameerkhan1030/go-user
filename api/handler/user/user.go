package user

import (
	"fmt"
	"test/api/server/response"
	"test/internal/app"
	"test/internal/database"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type User struct {

	Email string `json:"email"`
	Password string `json:"password"`
}

func RegisterRoutes(e *echo.Group,app app.App){

	e.POST("",userCreateHandler(app))
	e.GET("",userListHandler(app))
}

func userCreateHandler(app app.App) func(ctx echo.Context) error{

	return func(ctx echo.Context) error{

		var user User

		err:= ctx.Bind(&user)

		if err != nil{
			fmt.Println(err)
			code := 400
			message := err.Error()
			return response.BadRequest(ctx, nil, message, code)
		}
		if user.Email == "" {
			code := 400
			message := "Email shouldnt be null or empty"
			return response.BadRequest(ctx, nil, message, code)
		}

		if user.Password == "" {
			code := 400
			message := err.Error()
			return response.BadRequest(ctx, nil, message, code)
		}
		usr := &database.User{
			Email:user.Email,
			Password:user.Password,
		}
		err = database.AddUser(usr,app.DataStore)
		if err != nil{
			code := 500
			message := err.Error()
			return response.InternalServerError(ctx,nil,message,code)
		}
		app.Log.WithFields(logrus.Fields{
			"email": user.Email,
		}).Info("User Created Sucessfully")

		return response.OK(ctx,nil,"User Created Sucessfully")
	}
}

func userListHandler(app app.App) func(ctx echo.Context) error{

	return func(ctx echo.Context) error{
		users,err := database.UserList(app.DataStore)
		if err != nil{
			code := 500
			message := err.Error()
			return response.InternalServerError(ctx,nil,message,code)
		}
		var userList []User
		for _,v := range users {

			userList = append(userList,User{
				Email: v.Email,
				Password: v.Password,
			})
		}
		return response.OK(ctx,userList,"Got Users List Successfully")
	}
}