package models

import (
	"context"
	"strings"
	"time"

	"github.com/itanhaemprev/api/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//postCollection is a instance of database with collection seted
var postCollection *mongo.Collection

func init() {
	postCollection = database.Db.Collection("posts")
}

//Post is a model for crud
type Post struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string               `bson:"title,omitempty" json:"title,omitempty" binding:"required"`
	Description string               `bson:"description,omitempty" json:"description,omitempty" binding:"required"`
	Thumbnail   string               `bson:"thumbnail,omitempty" json:"thumbnail,omitempty"`
	Tags        []string             `bson:"tags,omitempty" json:"tags,omitempty" binding:"required"`
	Categories  []primitive.ObjectID `bson:"categories,omitempty" json:"categories,omitempty"`
	Slug        string               `bson:"slug,omitempty" json:"slug,omitempty"`
	CreatedAt   primitive.DateTime   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	ModifiedAt  primitive.DateTime   `bson:"modified_at,omitempty" json:"modified_at,omitempty"`
}

func (p *Post) init() {
	now := primitive.NewDateTimeFromTime(time.Now())
	p.CreatedAt = now
	p.ModifiedAt = now
	p.Slug = slug(p.Title)

}
func slug(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}

//GetPosts return []Post, err, paginator int64, limit int64
func (p *Post) GetPosts(page int64, limit int64) ([]Post, error, int64, int64) {
	var posts []Post
	//pular na paginacao
	skip := int64(0)
	if page > 0 {
		skip = (page * limit) - limit
	}
	total, err := postCollection.CountDocuments(context.TODO(), bson.D{})
	totalPages := total / limit
	findWithPaginate := options.FindOptions{Limit: &limit, Skip: &skip}
	cur, err := postCollection.Find(context.TODO(), bson.D{}, &findWithPaginate)

	if err != nil {
		return nil, err, 0, 0
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var post Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, err, 0, 0
		}
		posts = append(posts, post)
	}
	if err := cur.Err(); err != nil {
		return nil, err, 0, 0
	}
	return posts, nil, total, totalPages
}
func (p *Post) GetPost() error {
	result := postCollection.FindOne(context.TODO(), bson.M{"_id": p.ID})
	if err := result.Decode(&p); err != nil {
		return err
	}
	return nil
}
func (p *Post) CreatePost() error {
	p.init()
	result, err := postCollection.InsertOne(context.TODO(), p)
	if err != nil {
		return err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		p.ID = oid
	}
	return nil
}
func (p *Post) UpdatePost() error {
	result := postCollection.FindOneAndUpdate(context.TODO(), bson.M{"_id": p.ID}, bson.M{"$set": p, "$currentDate": bson.M{"modified_at": true}})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
func (p *Post) DeletePost() error {
	result := postCollection.FindOneAndDelete(context.TODO(), bson.M{"_id": p.ID})
	if result.Err() != nil {
		return result.Err()
	}
	return nil
}
