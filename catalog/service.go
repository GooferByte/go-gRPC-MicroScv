 

 import ()

 type Service interface{
	PostProduct(ctx context.Context, name, description string, price float64) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error)
	GetProducstByID(ctx context.Context, ids []string)([]Product, error)
	SearchProducts(ctx context.Context, query string, skip uint64, take uint64)	([]Product, error)
 }

 typr Product struct{
	ID			string		`json:"id`
	Name		string		`json:"name"`
	Description	string		`json:description`
	Price		string		`json:price`
 }

 type catalogService struct{
	repository Repository
 }

 func NewService(r Repository) Service{
	return &catalogService{r}
 }

 func (s *catalogService) PostProduct(ctx context.Context, name, description string, price float64) (*Product, error){

 }

 func (s *catalogService) GetProduct(ctx context.Context, id string) (*Product, error){

 }

 func (s *catalogService) GetProducts(ctx context.Context, skip uint64, take uint64) ([]Product, error){

 }

 func (s *catalogService) GetProducstByIDs(ctx context.Context, ids []string)([]Product, error){

 }

 func (s *catalogService) SearchProducts(ctx context.Context, query string, skip uint64, take uint64)	([]Product, error){
	
 }
