package storage

import (
	"errors"
	"fmt"
	"catalog-service-management-api/internal/domain/models"
	"catalog-service-management-api/internal/infrastructure/entity"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const VersionTableName = "Versions"

type PostgresStorage struct {
	db *gorm.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	// 获取数据库主机名和其他连接参数
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&entity.User{}, &entity.Service{}, &entity.Version{})
	if err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	log.Println("Database migration completed successfully")

	return &PostgresStorage{db: db}, nil
}

func (p *PostgresStorage) ListServices(query string, sortBy string, sortDir string, page int, pageSize int) ([]models.Service, int, error) {
	var services []entity.Service
	var total int64

	db := p.db.Model(&entity.Service{}).Preload(VersionTableName)

	if query != "" {
		// GORM 自动为我们处理了 SQL 参数化
		db = db.Where("name ILIKE ? OR description ILIKE ?", "%"+query+"%", "%"+query+"%")
	}

	// Count total before pagination
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("error counting services: %w", err)
	}

	// Apply sorting
	switch sortBy {
	case "name", "created_at":
		if sortDir != "asc" && sortDir != "desc" {
			sortDir = "desc"
		}
		db = db.Order(fmt.Sprintf("%s %s", sortBy, sortDir))
	default:
		db = db.Order("created_at desc")
	}

	// Apply pagination
	offset := (page - 1) * pageSize
	db = db.Offset(offset).Limit(pageSize)

	if err := db.Find(&services).Error; err != nil {
		return nil, 0, fmt.Errorf("error fetching services: %w", err)
	}

	var result []models.Service
	for _, s := range services {
		result = append(result, toDomainService(s))
	}

	return result, int(total), nil
}

var ErrServiceNotFound = errors.New("service not found")

func (p *PostgresStorage) GetService(id string) (models.Service, error) {
	var service entity.Service
	err := p.db.Preload(VersionTableName).First(&service, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Service{}, ErrServiceNotFound
		}
		return models.Service{}, err
	}
	return toDomainService(service), nil
}

func toDomainService(e entity.Service) models.Service {
	var versions []models.Version
	for _, v := range e.Versions {
		versions = append(versions, models.Version{
			Number:      v.Number,
			Description: v.Description,
			CreatedAt:   v.CreatedAt,
		})
	}

	return models.Service{
		ID:          fmt.Sprint(e.ID),
		Name:        e.Name,
		Description: e.Description,
		OwnerID:     fmt.Sprint(e.OwnerID),
		Versions:    versions,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}
