package models

import (
	"context"
	"github.com/itanhaemprev/api2/database"
	"github.com/leandrofreires/afiliate/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)
//UserCollection is a collections instanced
var UserCollection *mongo.Collection
func init() {
	UserCollection = database.Db.Collection("users")
}
type User struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" `
	FirstName  string             `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName   string             `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Email      string             `bson:"email,omitempty" json:"email,omitempty"`
	Password   string             `bson:"password,omitempty" json:"password,omitempty"`
	Active     bool               `bson:"active,omitempty" json:"active,omitempty"`
	CreatedAt  primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	ModifiedAt primitive.DateTime `bson:"modified_at,omitempty" json:"modified_at,omitempty"`
}
//GetUsers return of database a list of users
func (u *User) GetUsers(page int64, limit int64) ([]User, error) {
	//ira receber os usuarios
	var users []User
	//pular na paginacao
	skip := int64(0)

	if page > 0 {
		skip = (page * limit) - limit
	}

	findWithPaginate := options.FindOptions{Limit: &limit, Skip: &skip}
	cur, err := UserCollection.Find(context.TODO(), bson.D{}, &findWithPaginate)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var result User
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		users = append(users, result)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

//CreateUser save on database user
func (u *User) CreateUser(user User) (User, error) {
	user.Password = util.HashAndSalt([]byte(user.Password))

	now := primitive.NewDateTimeFromTime(time.Now())
	user.CreatedAt = now
	user.ModifiedAt = now

	result, err := UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return user, err
	}
	err = UserCollection.FindOne(context.TODO(), bson.M{"_id": result.InsertedID}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//UpdateUser on database
func (u *User) UpdateUser(id string, user User) (User, error) {
	idBson, err := primitive.ObjectIDFromHex(id)
	if user.Password != "" {
		user.Password = util.HashAndSalt([]byte(user.Password))
	}
	if err != nil {
		return user, err
	}
	_, err = UserCollection.UpdateOne(context.TODO(), bson.M{"_id": idBson}, bson.M{"$set": user, "$currentDate": bson.M{"modified_at": true}})

	if err != nil {
		return user, err
	}
	err = UserCollection.FindOne(context.TODO(), bson.M{"_id": idBson}).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}