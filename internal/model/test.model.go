package model

import "go.mongodb.org/mongo-driver/bson/primitive"


type Test struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Url      string             `bson:"url"`
    Status   string             `bson:"status,omitempty"`
    TasksID  primitive.ObjectID `bson:"tasks_id,omitempty"` 
    CreatedAt primitive.DateTime  `bson:"created_at,omitempty"`
}
