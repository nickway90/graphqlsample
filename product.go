package main

import (
	"github.com/graphql-go/graphql"
)

var ProductObject = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"product_id": &graphql.Field{
			Name: "ProductID",
			Type: graphql.Int,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(ProductData)
				return product.ProductID, nil
			},
		},
		"product_name": &graphql.Field{
			Name: "ProductName",
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(ProductData)
				return product.ProductName, nil
			},
		},
		"product_price": &graphql.Field{
			Name: "ProductPrice",
			Type: graphql.Float,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				product := p.Source.(ProductData)
				return product.ProductPrice, nil
			},
		},
		"shop": &graphql.Field{
			Name: "Shop",
			Type: ShopObject,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return ShopData{
					ShopID:   123,
					ShopName: "Tokopedia",
				}, nil
			},
		},
	},
})
