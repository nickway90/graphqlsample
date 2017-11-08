package main

import "github.com/graphql-go/graphql"

type ShopData struct {
	ShopID   int    `json:"shop_id"`
	ShopName string `json:"shop_name"`
}

var ShopObject = createShopObject()

func createShopObject() *graphql.Object {
	fields := graphql.BindFields(ShopData{})
	return graphql.NewObject(graphql.ObjectConfig{
		Name:   "Shop",
		Fields: fields,
	})
}
