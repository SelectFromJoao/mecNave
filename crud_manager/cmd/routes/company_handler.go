package routes

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"mecnave.com/mod/crud_manager/models"
	"mecnave.com/mod/crud_manager/repository"
)

func GetCompanyByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "0")

	Company, err := repository.GetOne(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "Company not found",
			Error:      true,
			Data:       nil,
		})
	}

	if Company.Id == 0 {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "Company not found",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(Company)
}

func GetAllCompanies(ctx *fiber.Ctx) error {

	Companys, err := repository.GetAllCompanies()

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "Company not found",
			Error:      true,
			Data:       err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(Companys)
}

func CreateCompany(ctx *fiber.Ctx) error {

	Company := new(models.Company)

	if err := ctx.BodyParser(Company); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Request body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(Company.Password), 12)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't create",
			Error:      true,
			Data:       nil,
		})
	}

	u := models.Company{
		Email:          Company.Email,
		BrandSpecialty: Company.BrandSpecialty,
		Localization:   Company.Localization,
		Description:    Company.Description,
		CompanyTitle:   Company.CompanyTitle,
		Password:       string(password),
		CreatedAt:      time.Now(),
	}

	_, err = repository.CreateCompany(&u)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't create",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "Company created successfully",
		Error:      false,
		Data:       nil,
	})
}

func UpdateComapany(ctx *fiber.Ctx) error {

	CompanyModel := new(models.Company)

	if err := ctx.BodyParser(CompanyModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Request body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	Company, err := repository.GetOneCompany(CompanyModel.Id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't find",
			Error:      true,
			Data:       nil,
		})
	}

	Company.Description = CompanyModel.Description
	Company.CompanyTitle = CompanyModel.CompanyTitle
	Company.Reviews = CompanyModel.Reviews
	Company.Localization = Company.Localization
	Company.Email = CompanyModel.Email

	if CompanyModel.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(CompanyModel.Password), 12)

		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
				StatusCode: 500,
				Message:    "Company couldn't update",
				Error:      true,
				Data:       nil,
			})
		}

		Company.Password = string(password)
	}

	_, err = repository.UpdateCompany(&Company)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't update",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "Company updated successfully",
		Error:      false,
		Data:       nil,
	})

}

func DeleteCompany(ctx *fiber.Ctx) error {

	CompanyModel := new(models.Company)

	if err := ctx.BodyParser(CompanyModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Response body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	Company, err := repository.GetOne(CompanyModel.Id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	_, err = repository.Delete(&Company)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "Company deleted successfully",
		Error:      false,
		Data:       nil,
	})

}

func LoginCompany(ctx *fiber.Ctx) error {
	CompanyModel := new(models.Company)

	if err := ctx.BodyParser(CompanyModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Response body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	Company, err := repository.GetByEmail(CompanyModel.Email)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	//isPasswordValid, err := checkPassword(Company.Password, CompanyModel.Password)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Company couldn't login",
			Error:      true,
			Data:       nil,
		})
	}

	Company.Password = ""

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "Company logged in successfully",
		Error:      false,
		Data:       Company,
	})
}
