package adapters

import (
	"log"
	"strings"

	"gorm.io/gorm"

	"github.com/ceblay/payments/pkg/app/query"
	pg "github.com/ceblay/payments/pkg/domain/paymentgateway"
)

type providerDBModel struct {
	gorm.Model
	ID       string `gorm:"primaryKey"`
	Platform string
	Country  string
}

type SqliteRepository struct {
	db *gorm.DB
}

func applySchema(db *gorm.DB) error {
	return db.AutoMigrate(&providerDBModel{})
}

func NewSqliteRepository(db *gorm.DB) *SqliteRepository {
	err := applySchema(db)
	if err != nil {
		log.Print("Could not apply DB schema")
	}
	return &SqliteRepository{
		db: db,
	}
}

func marshalProvider(p pg.Provider) providerDBModel {
	return providerDBModel{
		ID:       p.UUID(),
		Platform: p.Platform().String(),
		Country:  strings.Join(p.Country(), ","),
	}
}

func unmarshalProvider(p providerDBModel) (*pg.Provider, error) {
	data, err := pg.UnmarshalProviderFromDatabase(
		p.ID,
		p.Platform,
		strings.Split(p.Country, ","),
	)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func modelToQuery(p providerDBModel) query.Gateway {
	return query.Gateway{
		ID:       p.ID,
		Platform: p.Platform,
		Country:  p.Country,
	}
}

func (repo *SqliteRepository) GetGatewayByID(id string) (*pg.Provider, error) {
	model := &providerDBModel{}
	if err := repo.db.Where("id = ?", id).First(model).Error; err != nil {
		return nil, err
	}
	return unmarshalProvider(*model)
}

func (repo *SqliteRepository) FindAllGateways() ([]query.Gateway, error) {
	var models []*providerDBModel
	results := make([]query.Gateway, 0)

	if err := repo.db.Find(&models).Error; err != nil {
		return nil, err
	}
	for _, obj := range models {
		temp := modelToQuery(*obj)
		results = append(results, temp)
	}
	return results, nil
}

func (repo *SqliteRepository) AddNewGateway(data pg.Provider) (*pg.Provider, error) {
	model := marshalProvider(data)

	if err := repo.db.Create(&model).Error; err != nil {
		return nil, err
	}
	return unmarshalProvider(model)
}
