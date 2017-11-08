package main

import (
	"log"

	"github.com/graphql-go/graphql"
)

func CreateSchema() graphql.Schema {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "RootQuery",
			Fields: graphql.Fields{
				"get_product": &graphql.Field{
					Name: "GetProduct",
					Type: ProductObject,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// query db product
						product := getProduct()
						return product, nil
					},
				},
				"get_product_by_id": &graphql.Field{
					Name: "GetProductByID",
					Type: ProductObject,
					Args: graphql.FieldConfigArgument{
						"product_id": &graphql.ArgumentConfig{
							Type:         graphql.Int,
							DefaultValue: 0,
						},
					},
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						// query db product
						productID := p.Args["product_id"].(int)
						product := getProductByID(productID)
						return product, nil
					},
				},
				"get_product_list": &graphql.Field{
					Name: "GetProductList",
					Type: graphql.NewList(ProductObject),
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						products := getProductList()
						return products, nil
					},
				},
			},
		}),
	})
	if err != nil {
		log.Fatal(err)
	}
	return schema
}

type ProductData struct {
	ProductID    int
	ProductName  string
	ProductPrice float64
}

func getProduct() ProductData {
	return ProductData{
		ProductID:    123,
		ProductName:  "sepatu",
		ProductPrice: 1000,
	}
}

func getProductByID(productID int) ProductData {
	return ProductData{
		ProductID:    productID,
		ProductName:  "sepatu",
		ProductPrice: 1000,
	}
}

func getProductList() []ProductData {
	return []ProductData{
		{
			ProductID:    1,
			ProductName:  "Satu",
			ProductPrice: 10000,
		},
		{
			ProductID:    2,
			ProductName:  "Dua",
			ProductPrice: 20000,
		},
	}
}
