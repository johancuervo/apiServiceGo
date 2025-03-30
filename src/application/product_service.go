package application

import (
	"log"

	"github.com/johancuervo/apiServiceGo/src/domain"
)

// ProductUseCase maneja la lógica de negocio
type ProductUseCase struct {
	repo domain.ProductRepository
}

// NewProductUseCase crea un nuevo caso de uso de productos
func NewProductUseCase(repo domain.ProductRepository) *ProductUseCase {
	return &ProductUseCase{repo: repo}
}

// SaveProducts guarda los productos en la base de datos
func (uc *ProductUseCase) SaveProducts(products []domain.Product) error {
	log.Printf("📝 Intentando guardar %d productos...", len(products))

	for _, product := range products {
		log.Printf("➡️ Guardando producto: %v", product)

		err := uc.repo.SaveProducts([]domain.Product{product})
		if err != nil {
			log.Printf("❌ Error al guardar producto %s: %v", product.SKU, err)
			return err
		}
	}

	log.Println("✅ Todos los productos fueron guardados exitosamente")
	return nil
}
func (uc *ProductUseCase) GetProducts() ([]domain.Product, error) {
	log.Println("🔍 Recuperando productos de MongoDB...")
	return uc.repo.GetProducts()
}
