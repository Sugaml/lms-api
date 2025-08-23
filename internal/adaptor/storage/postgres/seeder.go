package postgres

import (
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sugaml/lms-api/internal/core/domain"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) {
	categories := []string{
		"Finance", "Marketing", "Management", "Economics",
		"Accounting", "Operations", "Human Resources", "Information Technology",
	}

	for _, name := range categories {
		category := &domain.Category{Name: name, Slug: GenerateSlug(name)}
		if err := db.FirstOrCreate(category, domain.Category{Name: name, Slug: GenerateSlug(name)}).Error; err != nil {
			logrus.Error("Failed to seed category:", name, err)
		}
	}
	logrus.Info("Categories seeded successfully")
}

func SeedPrograms(db *gorm.DB) {
	programs := []string{
		"MBA", "MBA IT", "MBA Finance", "MBA GLM",
	}
	for _, name := range programs {
		program := &domain.Program{Name: name, Slug: GenerateSlug(name)}
		if err := db.FirstOrCreate(program, domain.Program{Name: name, Slug: GenerateSlug(name)}).Error; err != nil {
			logrus.Error("Failed to seed program:", name, err)
		}
	}
	logrus.Info("Programs seeded successfully")
}

func GenerateSlug(name string) string {
	// Trim leading/trailing spaces
	slug := strings.TrimSpace(name)
	// Convert to lowercase
	slug = strings.ToLower(slug)
	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}
