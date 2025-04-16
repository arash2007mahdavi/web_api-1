package migrations

import (
	"github.com/arash2007mahdavi/web-api-1/config"
	"github.com/arash2007mahdavi/web-api-1/data/database"
	"github.com/arash2007mahdavi/web-api-1/data/models"
	"github.com/arash2007mahdavi/web-api-1/pkg/logging"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var log = logging.NewLogger(config.GetConfig())

func Up_1() {
	database := database.GetDb()

	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	if !database.Migrator().HasTable(country) {
		tables = append(tables, country)
	}

	if !database.Migrator().HasTable(city) {
		tables = append(tables, city)
	}

	if !database.Migrator().HasTable(user) {
		tables = append(tables, user)
	}

	if !database.Migrator().HasTable(role) {
		tables = append(tables, role)
	}

	if !database.Migrator().HasTable(userRole) {
		tables = append(tables, userRole)
	}

	database.Migrator().CreateTable(tables...)
	log.Info(logging.Postgres, logging.Migration, "tables created", nil)
	CreateDefaultInformation(database)
}

func createRoleIfNotExists(database *gorm.DB, r *models.Role) {
	exists := 0

	database.
		Model(&models.Role{}).
		Select("1").
		Where("name = ?", r.Name).
		First(&exists)
	
	if exists == 0 {
		database.Create(r)
	}
}

func CreateDefaultInformation(database *gorm.DB) {
	adminRole := models.Role{Name: "admin"}
	createRoleIfNotExists(database, &adminRole)

	defaultRole := models.Role{Name: "default"}
	createRoleIfNotExists(database, &defaultRole)

	u := models.User{UserName: "default", FirstName: "test", LastName: "test", MobileNumber: "09876543211"}
	email := "test@gmail.com"
	pass := "12345678"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	u.Email = email

	createAdminUserIfNotExists(database, &u, adminRole.Id)
}

func createAdminUserIfNotExists(database *gorm.DB, u *models.User, roleId int) {
	exists := 0
	database.
		Model(&models.User{}).
		Select("1").
		Where("username = ?", u.UserName).
		First(&exists)

	if exists == 0 {
		database.Create(u)
		ur := models.UserRole{UserId: u.Id, RoleId: roleId}
		database.Create(&ur)
	}
}