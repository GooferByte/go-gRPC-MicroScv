package main

import (
	"context"
	"log"
	"time"
)

type queryResolver struct {
	server *Server
}

func (r *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput, id *string) ([]*Account, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if id != nil {
		r, err := r.server.accountClient.GetAccount(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*Account{{
			ID:   r.ID,
			Name: r.Name,
		}}, nil
	}
	skip, take := uint64(0), uint64(0)
	if pagination != nil {
		skip, take = pagination.bounds()
	}
	accountList, err := r.server.accountClient.GetAccounts(ctx, skip, take)
	if err != nil {
		log.Panicln(err)
		return nil, err
	}
	var accounts []*Account
	for _, a := range accountList {
		account := &Account{
			ID:   a.ID,
			Name: a.Name,
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (r *queryResolver) Products(ctx context.Context, pagination *PaginationInput, query *string, id *string) ([]*Product, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if id != nil {
		product, err := r.server.catalogClient.GetProduct(ctx, *id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		return []*Product{{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		}}, nil
	}
	skip, take := uint64(0), uint64(0)
	if pagination != nil {
		skip, take = pagination.bounds()
	}
	q := ""
	if query != nil {
		q = *query
	}
	productList, err := r.server.catalogClient.GetProducts(ctx, skip, take, nil, q)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var products []*Product
	for _, p := range productList {
		products = append(products,
			&Product{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
			},
		)
	}
	return products, nil
}

func (p PaginationInput) bounds() (uint64, uint64) {
	skip := uint64(0)
	take := uint64(0)
	if p.Skip != nil {
		skip = uint64(*p.Skip)
	}
	if p.Take != nil {
		take = uint64(*p.Take)
	}
	return skip, take
}
