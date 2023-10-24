package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/google/uuid"
	"github.com/street-jackal/gardenwars/env"
	"github.com/street-jackal/gardenwars/repository/models"
)

const Plants = "Plants"

type PlantsRepo interface {
	Insert(ctx context.Context, p *models.Plant) error
	InsertMany(ctx context.Context, ps []*models.Plant) error

	Get(ctx context.Context, id string) (*models.Plant, error)
	GetByCommon(ctx context.Context, common string) ([]*models.Plant, error)
	GetByBotanical(ctx context.Context, botanical string) ([]*models.Plant, error)

	Delete(ctx context.Context, id string) error
}

type plantsRepo struct {
	col *mongo.Collection
}

// NewPlantsRepo returns a PlantsRepo instance
func NewPlantsRepo(ctx context.Context) (PlantsRepo, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(env.GetDbURL()))
	if err != nil {
		return nil, err
	}

	col := client.Database(env.GetDbName()).Collection(Plants)

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
					bson.E{Key: "Botanical", Value: 1},
					bson.E{Key: "Common", Value: 1},
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

	return &plantsRepo{col: col}, nil
}

func (r *plantsRepo) Insert(ctx context.Context, p *models.Plant) error {
	p.CreatedAt = time.Now()
	p.ID = uuid.NewString()
	_, err := r.col.InsertOne(ctx, p)
	return err
}

func (r *plantsRepo) InsertMany(ctx context.Context, ps []*models.Plant) error {
	// set the create-time and ID of all objects, and populate the documents to be inserted
	docs := make([]any, 0, len(ps))
	for _, p := range ps {
		p.ID = uuid.NewString()
		p.CreatedAt = time.Now()
		docs = append(docs, p)
	}

	// insert the final objects
	_, err := r.col.InsertMany(ctx, docs)
	if err != nil {
		return err
	}

	return nil
}

func (r *plantsRepo) Get(ctx context.Context, id string) (*models.Plant, error) {
	// define how much time this call is allowed to take(from the full context time).
	// this has to be less than the allowed time of the full context time.
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	filter := bson.M{"ID": id}

	plant := &models.Plant{}
	err := r.col.FindOne(ctx, filter).Decode(plant)
	if err != nil {
		return nil, err
	}

	return plant, nil
}

func (r *plantsRepo) GetByCommon(ctx context.Context, common string) ([]*models.Plant, error) {
	filter := bson.M{"Common": common}

	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	plants := make([]*models.Plant, 0)
	for cursor.Next(ctx) {
		plant := &models.Plant{}
		if err := cursor.Decode(plant); err != nil {
			return nil, err
		}
		plants = append(plants, plant)
	}

	return plants, nil
}

func (r *plantsRepo) GetByBotanical(ctx context.Context, botanical string) ([]*models.Plant, error) {
	filter := bson.M{"Botanical": botanical}

	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var plants []*models.Plant
	if err := cursor.All(ctx, &plants); err != nil {
		return nil, err
	}

	return plants, nil
}

func (r *plantsRepo) Delete(ctx context.Context, id string) error {
	filter := bson.M{"ID": id}

	if res, err := r.col.DeleteOne(ctx, filter); err != nil {
		return err
	} else if res.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}

	return nil
}
