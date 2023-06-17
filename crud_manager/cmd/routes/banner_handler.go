package routes

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"

	"mecnave.com/mod/crud_manager/models"
	"mecnave.com/mod/crud_manager/repository"
	"mecnave.com/mod/crud_manager/util"
)

func GetBannerByID(ctx *fiber.Ctx) error {

	id := ctx.Params("id", "0")

	banner, err := repository.GetOne(id)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "Banner not found",
			Error:      true,
			Data:       nil,
		})
	}

	if banner.Id == 0 {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "Banner not found",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(banner)
}

func GetAllBanners(ctx *fiber.Ctx) error {

	users, err := repository.GetAllBanners()

	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(models.ResponseModel{
			StatusCode: 404,
			Message:    "Banner not found",
			Error:      true,
			Data:       err,
		})
	}

	return ctx.Status(http.StatusOK).JSON(users)
}

func CreateBanner(ctx *fiber.Ctx) error {

	banner := new(models.Banner)
	file, err := ctx.FormFile("Banner")
	fileContent, _ := file.Open()
	byteContainer, err := ioutil.ReadAll(fileContent)

	var base64Encoding string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(byteContainer)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	base64Encoding += util.ToBase64(byteContainer)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "File not found",
			Error:      true,
			Data:       nil,
		})
	}

	if err := ctx.BodyParser(banner); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Request body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	u := models.Banner{
		BannerTitle: banner.BannerTitle,
		Banner:      base64Encoding,
		CreatedAt:   time.Now(),
	}

	_, err = repository.CreateBanner(&u)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Banner couldn't create",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "Banner created successfully",
		Error:      false,
		Data:       nil,
	})
}

func DeleteBanner(ctx *fiber.Ctx) error {

	userModel := new(models.Banner)

	if err := ctx.BodyParser(userModel); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Response body couldn't parse",
			Error:      true,
			Data:       nil,
		})
	}

	banner, err := repository.GetOne(userModel.Id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Banner couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	_, err = repository.Delete(&banner)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(models.ResponseModel{
			StatusCode: 500,
			Message:    "Banner couldn't delete",
			Error:      true,
			Data:       nil,
		})
	}

	return ctx.Status(http.StatusOK).JSON(models.ResponseModel{
		StatusCode: 200,
		Message:    "Banner deleted successfully",
		Error:      false,
		Data:       nil,
	})

}
