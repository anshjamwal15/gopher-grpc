package db

import (
	"context"
	"fmt"

	pb "github.com/aimbot1526/test-go/generated"
	"github.com/aimbot1526/test-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ShopRepository struct {
	pb.UnimplementedShopServiceServer
}

func (s *ShopRepository) CreateShop(ctx context.Context, req *pb.CreateShopReq) (*pb.CreateShopRes, error) {

	shop := req.Shop

	var users []models.UserModel

	for _, user := range req.Shop.Users {

		u := models.UserModel{
			UserId: primitive.NewObjectID(),
			Name:   user.Name,
		}

		users = append(users, u)
	}

	location := &models.Location{
		Type:        "Point",
		Coordinates: []float64{shop.Location.Long, shop.Location.Lat},
	}

	data := &models.ShopModel{
		ShopId:   primitive.NewObjectID(),
		Name:     shop.Name,
		Users:    users,
		Location: *location,
	}

	_, err := db.Collection("shop").InsertOne(mongoCtx, data)

	if err != nil {

		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}

	return &pb.CreateShopRes{Shop: &pb.Shop{Name: shop.Name, Location: &pb.Location{Lat: location.Coordinates[0], Long: location.Coordinates[1]}}}, nil
}

func (p *ShopRepository) NearByShop(ctx context.Context, req *pb.NearByShopReq) (*pb.NearByShopRes, error) {

	loc := req.Location

	collection := db.Collection("shop")

	indexView := collection.Indexes()

	_, er := indexView.CreateOne(mongoCtx, mongo.IndexModel{
		Keys: bson.M{
			"location": "2dsphere",
		},
	})

	if er != nil {
		return nil, er
	}

	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{loc.Long, loc.Lat},
				},
				"$maxDistance": 50000, // 50 km in meters
			},
		},
	}

	cur, err := collection.Find(mongoCtx, filter)

	if err != nil {
		return nil, err
	}

	defer cur.Close(mongoCtx)

	shops := make([]*pb.Shop, 0)

	for cur.Next(mongoCtx) {
		var shop models.ShopModel
		if err := cur.Decode(&shop); err != nil {
			return nil, err
		}
		data := pb.Shop{
			Name: shop.Name,
			Location: &pb.Location{
				Lat:  shop.Location.Coordinates[1],
				Long: shop.Location.Coordinates[0],
			},
		}
		shops = append(shops, &data)
	}

	return &pb.NearByShopRes{Shop: shops}, nil
}

func (p *ShopRepository) NearByUsers(ctx context.Context, req *pb.NearByNeighbourReq) (*pb.NearByNeighbourRes, error) {

	loc := req.Location

	distance := 50000

	if req.Range > 0 {
		distance = int(req.Range)
	}

	collection := db.Collection("shop")

	indexView := collection.Indexes()

	_, er := indexView.CreateOne(mongoCtx, mongo.IndexModel{
		Keys: bson.M{
			"location": "2dsphere",
		},
	})

	if er != nil {
		return nil, er
	}

	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{loc.Long, loc.Lat},
				},
				"$maxDistance": distance,
			},
		},
	}

	cur, err := collection.Find(mongoCtx, filter)

	if err != nil {
		return nil, err
	}

	defer cur.Close(mongoCtx)

	users := make([]*pb.Users, 0)

	for cur.Next(mongoCtx) {
		var shop models.ShopModel
		if err := cur.Decode(&shop); err != nil {
			return nil, err
		}
		for _, user := range shop.Users {
			data := pb.Users{
				Name: user.Name,
			}
			users = append(users, &data)
		}
	}

	return &pb.NearByNeighbourRes{Users: users}, nil
}
