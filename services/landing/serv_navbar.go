package landing

import (
	"errors"
	"time"

	landing "github.com/SymbioSix/ProgressieAPI/models/landing"
	status "github.com/SymbioSix/ProgressieAPI/models/status"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type LandNavbarService struct {
	DB *gorm.DB
}

func NewLandNavbarService(db *gorm.DB) LandNavbarService {
	return LandNavbarService{DB: db}
}

// GetAllNavbar godoc
//
//	@Summary		Get all navbar components
//	@Description	Get all navbar components
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]landing.Land_Navbar
//	@Failure		500	{object}	status.StatusModel
//	@Router			/navbar [get]
func (service *LandNavbarService) GetAllNavbar(c fiber.Ctx) error {
	var navbars []landing.Land_Navbar
	if err := service.DB.Find(&navbars).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}
	return c.Status(fiber.StatusOK).JSON(navbars)
}

// CreateNavbarRequest godoc
//
//	@Summary		Create a new navbar component
//	@Description	Create a new navbar component
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Param			request	body		landing.LandNavbarRequest	true	"Navbar component data"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/navbar [post]
func (service *LandNavbarService) CreateNavbarRequest(c fiber.Ctx) error {
	var request landing.LandNavbarRequest
	if err := c.Bind().JSON(&request); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	navbar := landing.Land_Navbar{
		NavComponentName:  request.NavName,
		NavComponentGroup: request.NavGroup,
		NavComponentIcon:  request.NavIcon,
		Tooltip:           request.Tooltip,
		Endpoint:          request.Endpoint,
		CreatedBy:         "SYSTEM",
		CreatedAt:         time.Now(),
	}

	if err := service.DB.Create(&navbar).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Created successfully"})
}

// GetNavbarRequestByID godoc
//
//	@Summary		Get a navbar component by ID
//	@Description	Get a navbar component by ID
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Navbar component ID"
//	@Success		200	{object}	landing.Land_Navbar
//	@Failure		400	{object}	status.StatusModel
//	@Failure		404	{object}	status.StatusModel
//	@Failure		500	{object}	status.StatusModel
//	@Router			/navbar/{id} [get]
func (service *LandNavbarService) GetNavbarRequestByID(c fiber.Ctx) error {
	navComponentID := c.Params("id")
	var request landing.Land_Navbar
	if err := service.DB.Where("nav_component_id = ?", navComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Navbar component not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	response := request

	return c.Status(fiber.StatusOK).JSON(response)
}

// UpdateNavbarRequest godoc
//
//	@Summary		Update a navbar component
//	@Description	Update a navbar component
//	@Tags			Navbar Service
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int							true	"Navbar component ID"
//	@Param			request	body		landing.LandNavbarRequest	true	"Updated navbar component data"
//	@Success		200		{object}	status.StatusModel
//	@Failure		400		{object}	status.StatusModel
//	@Failure		404		{object}	status.StatusModel
//	@Failure		500		{object}	status.StatusModel
//	@Router			/navbar/{id} [put]
func (service *LandNavbarService) UpdateNavbarRequest(c fiber.Ctx) error {
	navComponentID := c.Params("id")

	var updatedRequest landing.LandNavbarRequest
	if err := c.Bind().JSON(&updatedRequest); err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusBadRequest).JSON(stat)
	}

	var request landing.Land_Navbar
	if err := service.DB.Where("nav_component_id = ?", navComponentID).First(&request).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			stat := status.StatusModel{Status: "fail", Message: "Navbar component not found"}
			return c.Status(fiber.StatusNotFound).JSON(stat)
		}
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	request.NavComponentName = updatedRequest.NavName
	request.NavComponentGroup = updatedRequest.NavGroup
	request.NavComponentIcon = updatedRequest.NavIcon
	request.Tooltip = updatedRequest.Tooltip
	request.Endpoint = updatedRequest.Endpoint
	request.UpdatedBy = "SYSTEM"
	request.UpdatedAt = time.Now()

	if err := service.DB.Save(&request).Error; err != nil {
		stat := status.StatusModel{Status: "fail", Message: err.Error()}
		return c.Status(fiber.StatusInternalServerError).JSON(stat)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Updated successfully"})
}
