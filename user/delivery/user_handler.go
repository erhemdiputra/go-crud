package delivery

import (
	"context"
	"net/http"
	"strconv"

	"github.com/erhemdiputra/go-crud/database"
	"github.com/erhemdiputra/go-crud/models"
	"github.com/erhemdiputra/go-crud/user"
	"github.com/erhemdiputra/go-crud/user/repository"
	"github.com/erhemdiputra/go-crud/user/usecase"
	"github.com/erhemdiputra/go-crud/views"
	"github.com/julienschmidt/httprouter"
)

type UserHandler struct {
	UserUsecase user.IUserUsecase
}

func NewUserHandler(router *httprouter.Router) {
	userRepository := repository.NewUserRepository(database.Get())
	userUsecase := usecase.NewUserUsecase(userRepository)
	handler := UserHandler{userUsecase}

	router.GET("/view/user/list", handler.HandleUserListPage)
	router.GET("/view/user/add", handler.HandleAddUserPage)
	router.POST("/view/user/add", handler.HandlePostAddUserPage)
	router.GET("/view/user/edit/:id", handler.HandleEditUserPage)
	router.POST("/view/user/edit", handler.HandlePostEditUserPage)
	router.GET("/view/user/delete/:id", handler.HandleDeleteUserPage)
}

func (u *UserHandler) HandleUserListPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userList, err := u.UserUsecase.GetList(context.Background())
	if err != nil {
		http.Error(w, "failed to get user list", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"userList": userList,
	}
	views.Get().ExecuteTemplate(w, "user-list.html", data)
}

func (u *UserHandler) HandleAddUserPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	views.Get().ExecuteTemplate(w, "add-user.html", nil)
}

func (u *UserHandler) HandlePostAddUserPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	name := r.FormValue("name")
	age := r.FormValue("age")

	parsedAge, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "invalid input", http.StatusInternalServerError)
		return
	}

	user := models.User{
		Name: name,
		Age:  parsedAge,
	}

	_, err = u.UserUsecase.Add(context.Background(), user)
	if err != nil {
		http.Error(w, "failed to add user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/user/list", http.StatusSeeOther)
}

func (u *UserHandler) HandleEditUserPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	idParsed, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid input", http.StatusInternalServerError)
		return
	}

	user, err := u.UserUsecase.GetByID(context.Background(), idParsed)
	if err != nil {
		http.Error(w, "failed to get user info", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"user": user,
	}

	views.Get().ExecuteTemplate(w, "edit-user.html", data)
}

func (u *UserHandler) HandlePostEditUserPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	age := r.FormValue("age")

	parsedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid input", http.StatusInternalServerError)
		return
	}

	parsedAge, err := strconv.Atoi(age)
	if err != nil {
		http.Error(w, "invalid input", http.StatusInternalServerError)
		return
	}

	user := models.User{
		ID:   parsedID,
		Name: name,
		Age:  parsedAge,
	}

	_, err = u.UserUsecase.Update(context.Background(), user)
	if err != nil {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/user/list", http.StatusSeeOther)
}

func (u *UserHandler) HandleDeleteUserPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	parsedID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "invalid input", http.StatusInternalServerError)
		return
	}

	_, err = u.UserUsecase.Delete(context.Background(), parsedID)
	if err != nil {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/user/list", http.StatusSeeOther)
}
