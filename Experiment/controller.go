package experiment

import (
	"github.com/gin-gonic/gin"
	"gorm/internal"
	"gorm/models"
	"sync"
)

type Controller struct {
	svc internal.UserService
}

var (
	cont     *Controller
	contOnce sync.Once
)

func NewController(svc internal.UserService) *Controller {
	contOnce.Do(func() {
		cont = &Controller{svc: svc}
	})
	return cont
}

func (c *Controller) CreateUser(ctx *gin.Context) {
	//name := ctx.Query("name")
	var req models.UserData
	_ = ctx.BindJSON(&req)
	res := c.svc.CreateUser(req.Name, req.DOB, req.Contact)

	ctx.JSON(201, gin.H{
		"message": res,
	})
}

func (c *Controller) CreateUserAll(ctx *gin.Context) {
	var req models.UserBulk
	_ = ctx.BindJSON(&req)
	res := c.svc.CreateUserAll(req.Name, req.DOB, req.Contact)
	ctx.JSON(201, gin.H{
		"message": res,
	})
}

func (c *Controller) GetUser(ctx *gin.Context) {
	name := ctx.Query("name")
	dob := ctx.Query("dob")
	contact := ctx.Query("contact")
	res := c.svc.GetUser(name, dob, contact)
	ctx.JSON(202, gin.H{
		"List": res,
	})
}

//func (c *Controller) GetUser(name, age, contact string) {
//	c.svc.GetUser(name, age, contact)
//}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	var userUpd models.UserUpd
	_ = ctx.BindJSON(&userUpd)

	res := c.svc.UpdateUser(userUpd.NameOrg, userUpd.DOBOrg, userUpd.ContactOrg, userUpd.Name, userUpd.DOB, userUpd.Contact)

	ctx.JSON(202, gin.H{
		"Message": res,
	})
}

func (c *Controller) DeleteUser(ctx *gin.Context) {
	var req models.UserData
	_ = ctx.BindJSON(&req)
	res := c.svc.DeleteUser(req.Name, req.DOB, req.Contact)

	ctx.JSON(202, gin.H{
		"message": res,
	})
}

func (c *Controller) UpsertUser(ctx *gin.Context) {
	var req models.UserData
	_ = ctx.BindJSON(&req)
	res := c.svc.UpsertUser(req.Name, req.DOB, req.Contact)

	ctx.JSON(202, gin.H{
		"message": res,
	})
}
