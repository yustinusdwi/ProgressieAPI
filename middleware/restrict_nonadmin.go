package middleware

import (
	"strings"

	auth "github.com/SymbioSix/ProgressieAPI/models/auth"
	s "github.com/SymbioSix/ProgressieAPI/setup"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm/clause"
)

func RestrictNonAdmin() fiber.Handler {
	return func(c fiber.Ctx) error {
		user, err := s.Client.Auth.GetUser()
		if err != nil || user == nil {
			return c.Redirect().Status(fiber.StatusUnauthorized).To("/v1/unauthorized")
		}
		var userRoleResponse *auth.UserRoleResponse
		if getUserRole := s.DB.Table("usr_roleuser").Preload(clause.Associations).Find(&userRoleResponse, "user_id = ?", user.ID); getUserRole.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": getUserRole.Error.Error()})
		}
		var isAdmin bool
		for i := 0; i < len(userRoleResponse.RoleData); i++ {
			if strings.Contains(strings.ToLower(userRoleResponse.RoleData[i].RoleName), "admin") {
				isAdmin = true
			}
		}
		if !isAdmin {
			return c.Redirect().Status(fiber.StatusForbidden).To("/v1/forbidden")
		} else {
			return c.Next()
		}
	}
}