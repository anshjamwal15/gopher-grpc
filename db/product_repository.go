package db

import (
	"context"

	pb "github.com/aimbot1526/test-go/generated"
	"github.com/aimbot1526/test-go/models"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductRepository struct {
	pb.UnimplementedProductServiceServer
}

func (p *ProductRepository) ProductList(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {

	products := make([]*pb.Product, 0)

	cur, err := db.Collection("product").Find(mongoCtx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(mongoCtx)

	for cur.Next(mongoCtx) {
		var product models.ProductModel
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		data := pb.Product{
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}
		products = append(products, &data)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &pb.ProductListResponse{Product: products}, nil
}

func (p *ProductRepository) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {

	products := make([]*pb.Product, 0)

	cur, err := db.Collection("product").Find(mongoCtx, bson.M{"name": req.Name})

	if err != nil {
		return nil, err
	}

	defer cur.Close(mongoCtx)

	for cur.Next(mongoCtx) {
		var product models.ProductModel
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		data := pb.Product{
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}
		products = append(products, &data)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{Product: products}, nil
}

func (p *ProductRepository) ServiceableProduct(ctx context.Context, req *pb.ServiceableProductRequest) (*pb.GetProductResponse, error) {

	cur, err := db.Collection("product").Find(mongoCtx, bson.M{"stock": bson.M{"$gt": 0}})

	if err != nil {
		return nil, err
	}

	defer cur.Close(mongoCtx)

	products := make([]*pb.Product, 0)

	for cur.Next(mongoCtx) {
		var product models.ProductModel
		if err := cur.Decode(&product); err != nil {
			return nil, err
		}
		data := pb.Product{
			Name:  product.Name,
			Price: product.Price,
			Stock: product.Stock,
		}
		products = append(products, &data)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &pb.GetProductResponse{Product: products}, nil
}
