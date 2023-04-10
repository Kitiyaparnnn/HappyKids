package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID		`json:"_id" bson:"_id"`
	First_Name      *string					`json:"first_name" validate:"reqiured,min=2,max=30"`
	Last_Name       *string					`json:"last_name" validate:"reqiured,min=2,max=30"`
	Password        *string					`json:"password" validate:"reqiured,min=6"`
	Email           *string					`json:"email" validate:"email, reqiured"`
	Phone           *string					`json:"phone" validate:"reqiured"`
	Token           *string					`json:"token"`
	Refresh_Token   *string					`json:"refresh_token"`
	Created_At      time.Time				`json:"create_at"`
	Updated_At      time.Time				`json:"update_at"`
	User_ID         string					`json:"user_id"`
	User_Cart       []ProductUser			`json:"user_cart" bson:"usercart"`
	Address_Details []AddressDetails		`json:"address_detail" bson:"address"`
	Order_Status    []Order					`json:"order_satus" bson:"orders"`
}

type Product struct {
	Product_ID  primitive.ObjectID			`bson:"_id"`
	Prduct_Name *string						`json:"product_name"`
	Price       *uint64						`json:"price"`
	Rating      *uint8						`json:"rating"`
	Image       *string						`json:"image"`
}

type ProductUser struct {
	Product_ID  primitive.ObjectID			`bson:"_id"`
	Prduct_Name *string						`json:"product_name" bson:"product_name"`
	Price       int							`json:"price" bson:"price"`
	Rating      *uint						`json:"rating" bson:"rating"`
	Image       *string						`json:"image" bson:"image"`
}

type AddressDetails struct {
	Address_ID primitive.ObjectID			`bson:"_id"`
	House      *string						`json:"house_name" bson:"house_name"`
	Street     *string						`json:"street_name" bson:"street_name"`
	City       *string						`json:"city" bson:"city"`
	Country    *string						`json:"country" bson:"country"`
	Pincode    *string						`json:"pincode" bson:"pincode"`
}

type Order struct {
	Order_ID       primitive.ObjectID		`bson:"_id"`
	Order_Cart     []ProductUser			`json:"order_list" bson:"order_list"`
	Ordered_At     time.Time				`json:"ordered_at" bson:"ordered_at"`
	Price          int						`json:"total_price" bson:"total_price"`
	Discount       *int						`json:"discpunt" bson:"discount"`
	Payment_Method Payment					`json:"payment" bson:"payment"`
}

type Payment struct {
	Digital bool
	COD     bool
}
