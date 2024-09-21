package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type Tasks struct {
    ID                primitive.ObjectID `bson:"_id"`
    SSL               bool               `bson:"ssl"`
    Captcha           bool               `bson:"captcha"`
    Logo              bool               `bson:"logo"`
    GreenNumber       bool               `bson:"green_number"`
    CardsValidation   bool               `bson:"cards_validation"`
    ReqValidation     bool               `bson:"req_validation"`
    CreatedAt         primitive.DateTime  `bson:"created_at"`
}
