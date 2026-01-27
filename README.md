# go-gRPC-MicroSvc

GraphQL gateway in front of three gRPC microservices (account, catalog, order). The GraphQL server listens on port **8000** in the container and is mapped to host **8080** via docker-compose.

## Running locally
```bash
# from repo root
docker compose up -d --build

# GraphQL endpoint & UI
open http://localhost:8080/playground
# or POST to http://localhost:8080/graphql
```

Environment (already provided in `.env`):
```
ACCOUNT_SERVICE_URL=account:8080
CATALOG_SERVICE_URL=catalog:8080
ORDER_SERVICE_URL=order:8080
```

## Core GraphQL operations
Use these in Playground or any client. All examples assume the compose stack is running.

### Query all accounts
```graphql
query {
  accounts {
    id
    name
  }
}
```

### Create an account
```graphql
mutation {
  createAccount(account: { name: "New Account" }) {
    id
    name
  }
}
```

### Create a product
```graphql
mutation {
  createProduct(product: { name: "New Product", description: "A new product", price: 300 }) {
    id
    name
    price
  }
}
```

### Create an order
Replace the IDs with real ones returned from the mutations above.
```graphql
mutation {
  createOrder(order: { accountId: "38qrf6QmzyFL9IoTPiz2HiQeAu4", products: [{ id: "38qztUuYG6alms3BD9hTVZ9hmlj", quantity: 2 }] }) {
    id
    totalPrice
    products {
      name
      quantity
    }
  }
}
```

### Get a specific account with its orders
```graphql
query {
  accounts(id: "38qrf6QmzyFL9IoTPiz2HiQeAu4") {
    name
    orders {
      id
      createdAt
      totalPrice
      products {
        name
        quantity
        price
      }
    }
  }
}
```

### Search products with pagination
```graphql
query {
  products(pagination: { skip: 0, take: 5 }, query: "New") {
    id
    name
    description
    price
  }
}
```

### Additional useful examples
- Single product by ID:
```graphql
query {
  products(id: "PUT_PRODUCT_ID_HERE") {
    id
    name
    price
  }
}
```
- Accounts with pagination:
```graphql
query {
  accounts(pagination: { skip: 0, take: 10 }) {
    id
    name
  }
}
```
- Orders total for an account:
```graphql
query {
  accounts(id: "38qrf6QmzyFL9IoTPiz2HiQeAu4") {
    name
    orders {
      totalPrice
    }
  }
}
```

## Service ports (inside compose)
- account: 8080
- catalog: 8080
- order: 8080
- graphql: 8000 (mapped to host 8080)

## Troubleshooting tips
- If GraphQL returns “internal system error”, check dependent services via `docker compose logs account|catalog|order`.
- Ensure IDs are strings in mutations/queries (wrap with quotes).
- If Elasticsearch (catalog_db) fails to start on ARM, `platform: linux/amd64` is already set in `docker-compose.yaml`.
