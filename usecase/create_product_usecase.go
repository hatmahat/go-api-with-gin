package usecase

import (
	"fmt"
	"mime/multipart"

	"github.com/jutionck/golang-api-with-gin/model"
	"github.com/jutionck/golang-api-with-gin/repository"
)

type CreateProductUseCase interface {
	CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error
}

type createProductUseCase struct {
	repo     repository.ProductRepository
	fileRepo repository.FileRepository
}

func (c *createProductUseCase) CreateProduct(newProduct *model.Product, file multipart.File, fileExt string) error {
	fileName := fmt.Sprintf("img-%s.%s", newProduct.ProductId, fileExt) //-> img-C001.jpg | img-C002.jpg dst...
	fileLocation, err := c.fileRepo.Save(file, fileName)
	if err != nil {
		return err
	}
	newProduct.ImgPath = fileLocation
	newProduct.UrlPath = fmt.Sprintf("/product/image/%s", newProduct.ProductId)
	err = c.repo.Add(newProduct)
	if err != nil {
		return err
	}

	return nil
}

func NewCreateProductUseCase(repo repository.ProductRepository, fileRepo repository.FileRepository) CreateProductUseCase {
	return &createProductUseCase{
		repo:     repo,
		fileRepo: fileRepo,
	}
}
