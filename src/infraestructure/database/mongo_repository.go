package infrastructure

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/johancuervo/apiServiceGo/src/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoRepository implementa ProductRepository
type MongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// NewMongoRepository crea un nuevo repositorio MongoDB
func NewMongoRepository() (*MongoRepository, error) {
	uri := os.Getenv("DB_CONEXION")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECCION")
	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB!")

	collection := client.Database(dbName).Collection(collectionName)
	return &MongoRepository{client: client, collection: collection}, nil
}

// SaveProducts almacena los productos en MongoDB con Upsert
func (r *MongoRepository) SaveProducts(products []domain.Product) error {
	ctx := context.Background()
	now := time.Now()

	for _, product := range products {
		product.Fecha_entrada = now

		filter := bson.M{"sku": product.SKU}
		update := bson.M{"$set": product}
		opts := options.Update().SetUpsert(true)

		_, err := r.collection.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			log.Printf("Error saving product %s to MongoDB: %v", product.SKU, err)
			return err
		}
	}
	return nil
}
func (r *MongoRepository) GetProducts() ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Creamos un cursor para iterar sobre los documentos
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("❌ Error al obtener productos de MongoDB: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []domain.Product
	if err = cursor.All(ctx, &products); err != nil {
		log.Printf("❌ Error al decodificar productos: %v", err)
		return nil, err
	}

	log.Printf("✅ %d productos recuperados de MongoDB", len(products))
	return products, nil
}
