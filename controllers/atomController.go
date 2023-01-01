package controllers

import (
	"strconv"

	"github.com/deuterium-cloud/go-web-crud/initializers"
	"github.com/deuterium-cloud/go-web-crud/models"
	"github.com/gin-gonic/gin"
)

var DefaultPage int64 = 0
var DefaultSize int64 = 10

func GetAtoms(context *gin.Context) {
	var body []models.Atom
	result := initializers.DB.Find(&body)
	if result.Error != nil {
		context.JSON(500, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, gin.H{"atoms": body, "total": result.RowsAffected})
}

func GetAtomsWithPagination(context *gin.Context) {

	pageString := context.Query("page")
	sizeString := context.Query("size")

	page := DefaultPage
	size := DefaultSize

	if pageString != "" {
		page, _ = strconv.ParseInt(pageString, 10, 32)
	}

	if pageString != "" {
		size, _ = strconv.ParseInt(sizeString, 10, 32)
	}

	var body []models.Atom
	result := initializers.DB.Limit(int(size)).Offset(int(page * size)).Find(&body)
	if result.Error != nil {
		context.JSON(500, gin.H{"error": result.Error})
		return
	}

	res := models.ResponseDto{
		Atoms: &body,
		Total: result.RowsAffected,
		Page:  page,
		Size:  size,
	}

	context.JSON(200, gin.H{"response": res})
}

func GetAtomById(context *gin.Context) {
	id := context.Param("id")
	var body models.Atom
	result := initializers.DB.First(&body, id)
	if result.Error != nil {
		context.JSON(500, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, gin.H{"atom": body})
}

func SaveNewAtom(context *gin.Context) {

	var body models.AtomRequest

	if err := context.Bind(&body); err != nil {
		return
	}

	atom := models.Map(body)

	result := initializers.DB.Create(&atom)

	if result.Error != nil {
		context.Status(500)
		return
	}

	context.JSON(200, gin.H{"createdAtom": atom})
}

func UpdateAtom(context *gin.Context) {
	id := context.Param("id")
	var body models.AtomRequest

	if err := context.Bind(&body); err != nil {
		return
	}

	var atom models.Atom
	result := initializers.DB.First(&atom, id)
	if result.Error != nil {
		context.JSON(500, gin.H{"error": result.Error})
		return
	}

	newResult := initializers.DB.Model(&atom).Updates(models.Atom{
		AtomNumber: body.AtomNumber,
		Mass:       body.Mass,
		Name:       body.Name,
		Symbol:     body.Symbol,
	})

	if newResult.Error != nil {
		context.JSON(500, gin.H{"error": result.Error})
		return
	}

	context.JSON(200, gin.H{"updatedAtom": atom})
}

func DeleteAtom(context *gin.Context) {
	id := context.Param("id")
	result := initializers.DB.Delete(&models.Atom{}, id)
	if result.Error != nil {
		context.JSON(500, gin.H{"error": result.Error})
		return
	}
	context.JSON(200, gin.H{"message": "Successfully deleted Atom with id=" + id})
}
