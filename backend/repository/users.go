package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/street-jackal/gardenwars/env"
	"github.com/street-jackal/gardenwars/repository/models"
)

const Users = "Users"

type UsersRepo interface {
	Insert(ctx context.Context, p *models.User) error

	Get(ctx context.Context, id string) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)

	AddFavorite(ctx context.Context, userID, plantID string) error
	RemoveFavorite(ctx context.Context, userID, plantID string) error

	Delete(ctx context.Context, id string) error
}

type usersRepo struct {
	col *mongo.Collection
}

// NewUsersRepo returns a UsersRepo instance
func NewUsersRepo(ctx context.Context) (UsersRepo, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.GetDbURL()))
	if err != nil {
		return nil, err
	}

	col := client.Database(env.GetDbName()).Collection(Users)

	_, err = col.Indexes().CreateMany(
		ctx,
		[]mongo.IndexModel{
			{
				Keys: bson.D{
					bson.E{Key: "ID", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{
					bson.E{Key: "Email", Value: 1},
				},
				Options: options.Index().SetUnique(true),
			},
			{
				Keys: bson.D{
					bson.E{Key: "createdAt", Value: 1},
				},
			},
			{
				Keys: bson.D{
					bson.E{Key: "updatedAt", Value: 1},
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &usersRepo{col: col}, nil
}

func (r *usersRepo) Insert(ctx context.Context, p *models.User) error {
	p.CreatedAt = time.Now()
	_, err := r.col.InsertOne(ctx, p)
	return err
}

func (r *usersRepo) Get(ctx context.Context, id string) (*models.User, error) {
	filter := bson.M{"ID": id}

	user := &models.User{}
	err := r.col.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *usersRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	filter := bson.M{"Email": email}

	user := &models.User{}
	err := r.col.FindOne(ctx, filter).Decode(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *usersRepo) GetAll(ctx context.Context) ([]*models.User, error) {
	cursor, err := r.col.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*models.User
	if err := cursor.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *usersRepo) AddFavorite(ctx context.Context, userID, plantID string) error {
	filter := bson.M{"ID": userID}
	update := bson.M{
		"$addToSet": bson.M{"Favorites": plantID},
		"$set":      bson.M{"updatedAt": time.Now()},
	}

	if err := r.col.FindOneAndUpdate(ctx, filter, update).Err(); err != nil {
		return err
	}

	return nil
}

func (r *usersRepo) RemoveFavorite(ctx context.Context, userID, plantID string) error {
	filter := bson.M{"ID": userID}
	update := bson.M{
		"$pull": bson.M{"Favorites": plantID},
		"$set":  bson.M{"updatedAt": time.Now()},
	}

	if err := r.col.FindOneAndUpdate(ctx, filter, update).Err(); err != nil {
		return err
	}

	return nil
}

func (r *usersRepo) Delete(ctx context.Context, id string) error {
	filter := bson.M{"ID": id}

	if res, err := r.col.DeleteOne(ctx, filter); err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
