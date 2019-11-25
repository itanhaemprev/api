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

//PostCollection is a instance of database with collection seted
var PostCollection *mongo.Collection

func init() {
	PostCollection = database.Db.Collection("posts")
}

//Post is a model for crud
type Post struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string               `bson:"title" json:"title" binding:"required"`
	Description string               `bson:"description" json:"description" binding:"required"`
	Thumbnail   string               `bson:"thumbnail" json:"thumbnail" binding:"required"`
	Tags        []string             `bson:"tags" json:"tags" binding:"required"`
	Categories  []primitive.ObjectID `bson:"categories" json:"categories"`
	Slug        string               `bson:"slug" json:"slug"`
	CreatedAt   primitive.DateTime   `bson:"created_at,omitempty" json:"created_at,omitempty"`
	ModifiedAt  primitive.DateTime   `bson:"modified_at,omitempty" json:"modified_at,omitempty"`
}

func (p *Post) init() {
	now := primitive.NewDateTimeFromTime(time.Now())
	p.CreatedAt = now
	p.ModifiedAt = now
	p.Slug = slug(p.Title)
	id, _ := primitive.ObjectIDFromHex("5ddb69ce792f3f34b490fc4e")
	p.Categories = append(p.Categories, id)

}
func slug(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), " ", "-")
}
func (p *Post) GetPosts(page int64, limit int64) ([]Post, error) {
	var posts []Post
	//pular na paginacao
	skip := int64(0)
	if page > 0 {
		skip = (page * limit) - limit
	}
	findWithPaginate := options.FindOptions{Limit: &limit, Skip: &skip}
	cur, err := PostCollection.Find(context.TODO(), bson.D{}, &findWithPaginate)

	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var post Post
		err := cur.Decode(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
func (p *Post) CreatePost(post Post) (Post, error) {
	post.init()
	result, err := PostCollection.InsertOne(context.TODO(), post)
	if err != nil {
		return post, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		post.ID = oid
	}
	return post, nil
}
