package cache

import (
	"github.com/sirupsen/logrus"
	"psw/pb"
)

var Cache = new(cache)

type cache struct {
	Config *config

	MapApplicationCategory map[string]int32
	Applications           []*pb.Application

	DefaultCategoryFields []*CategoryField

	Categories    []*category
	MapCategories map[int32]*category

	MapProducts map[string]int64
	Products    []string
}

func Init() error {
	// config
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	Cache.Config = cfg
	logrus.Debugf("cache.Init() config is %#v", Cache.Config)

	// Applications
	mapApplicationCategory, applications, err := loadApplications()
	if err != nil {
		return err
	}

	Cache.MapApplicationCategory = mapApplicationCategory
	Cache.Applications = applications
	logrus.Debugf("cache.Init() Applications Size is %d", len(Cache.Applications))

	// Categories
	defaultCategoryFields, categories, err := loadCategories()
	if err != nil {
		return err
	}

	Cache.DefaultCategoryFields = defaultCategoryFields
	Cache.Categories = categories
	logrus.Debugf("cache.Init() Categories Size is %d", len(Cache.Categories))

	// Map Categories
	mapCategories := make(map[int32]*category)
	for _, category := range categories {
		mapCategories[category.ProtoId] = category
	}
	Cache.MapCategories = mapCategories

	// MapProducts
	mapProducts, err := loadProducts()
	if err != nil {
		return err
	}
	Cache.MapProducts = mapProducts
	logrus.Debugf("cache.Init() Purchases Size is %d", len(Cache.MapProducts))

	// Products
	products := make([]string, 0)
	for product, _ := range mapProducts {
		products = append(products, product)
	}
	Cache.Products = products

	return nil
}
