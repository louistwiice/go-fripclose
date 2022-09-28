package mocks

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/fripclose/entity"
)

type fixtureMap struct {
	UserCreate1   	*entity.UserCreateUpdate
	UserDisplay1  	*entity.UserDisplay
	UserAdmin		*entity.UserDisplay
	UserStaff  		*entity.UserDisplay
	UserDeactivated *entity.UserDisplay
	User1Password	string
	UserCreate2  	*entity.UserCreateUpdate
	UserDisplay2 	*entity.UserDisplay
	User2Password 	string

	CategoryRoot	*entity.Category
	CategoryCloses	*entity.Category
	CategoryShoes	*entity.Category
	CategoryTshirt	*entity.Category

	UserList      	[]*entity.UserDisplay
	CategoryList	[]*entity.Category

	Server        	*gin.Engine
}

func GenerateFixture() (f fixtureMap) {

	f.UserDisplay1 = &entity.UserDisplay{
		ID:          "783ed845-387f-4c50-9a64-fef701b4dbb8",
		Email:       "mike@mail.com",
		Username:    "mike",
		FirstName:   "Mike",
		LastName:    "Spensor",
		IsActive:    false,
		IsStaff:     false,
		IsSuperuser: false,
	}
	f.User1Password = "mike_password"

	f.UserCreate1 = &entity.UserCreateUpdate{
		UserDisplay: *f.UserDisplay1,
		Password:    f.User1Password,
	}

	f.UserDisplay2 = &entity.UserDisplay{
		ID:          "883ed845-387f-4c50-9a64-fef701b4dbb8",
		Email:       "john@mail.com",
		Username:    "John",
		FirstName:   "Alfred",
		LastName:    "Goldman",
		IsActive:    false,
		IsStaff:     false,
		IsSuperuser: false,
	}
	f.User2Password = "john_password"

	f.UserCreate2 = &entity.UserCreateUpdate{
		UserDisplay: *f.UserDisplay1,
		Password:    f.User2Password,
	}

	f.UserAdmin = &entity.UserDisplay{
		ID:          "993ed845-387f-4c50-9a64-fef701b4dbb8",
		Email:       "admin@gmail.com",
		Username:    "admin",
		FirstName:   "the",
		LastName:    "Admin",
		IsActive:    true,
		IsStaff:     true,
		IsSuperuser: true,
	}

	f.UserStaff = &entity.UserDisplay{
		ID:          "443ed845-387f-4c50-9a64-fef701b4dbb8",
		Email:       "staff@gmail.com",
		Username:    "staff",
		FirstName:   "the",
		LastName:    "staff",
		IsActive:    true,
		IsStaff:     true,
		IsSuperuser: false,
	}

	f.UserDeactivated = &entity.UserDisplay{
		ID:          "463ed845-387f-4c50-9a64-fef701b4dbb8",
		Email:       "deactivated@gmail.com",
		Username:    "deactivated",
		FirstName:   "the",
		LastName:    "deactivated",
		IsActive:    false,
		IsStaff:     true,
		IsSuperuser: false,
	}

	f.UserList = append(f.UserList, f.UserDisplay1)
	f.UserList = append(f.UserList, f.UserDisplay2)

	f.CategoryRoot = &entity.Category{
		ID: 1,
		Title: "Men",
		ParentID: 0,
	}

	f.CategoryCloses = &entity.Category{
		ID: 2,
		Title: "Closes",
		ParentID: 1,
	}

	f.CategoryTshirt = &entity.Category{
		ID: 3,
		Title: "T shirt",
		ParentID: 2,
	}

	f.CategoryShoes = &entity.Category{
		ID: 4,
		Title: "Shoes",
		ParentID: 1,
	}

	f.CategoryList = append(f.CategoryList, f.CategoryRoot)
	f.CategoryList = append(f.CategoryList, f.CategoryCloses)
	f.CategoryList = append(f.CategoryList, f.CategoryTshirt)
	f.CategoryList = append(f.CategoryList, f.CategoryShoes)
	

	gin.SetMode(gin.TestMode)
	f.Server = gin.Default()

	return
}
