package ports

type CreateConf[T any] struct {
	Item T
}

type CreateManyConf[T any] struct {
	Items []T
}

type GetConf struct {
	Token       string
	OnlyActives bool
}

type GetByKeyConf struct {
	Value       any
	Key         string
	OnlyActives bool
}

type ListConf struct {
	OnlyActives bool
	Page        int
	Size        int
	ClientToken string
}

type UpdateConf struct {
	Token   string
	Updates map[string]interface{}
}

type InactivateConf struct {
	Token string
}

type ExistConf struct {
	Value       any
	Key         string
	OnlyActives bool
}

type SearchConf struct {
	Search      string
	Fields      []string
	OnlyActives bool
	Size        int
	Page        int
}

type BasePort[T any] interface {
	Exists(conf ExistConf) (bool, error)
	Create(conf CreateConf[T]) (*T, error)
	CreateMany(conf CreateManyConf[T]) error
	Get(conf GetConf) (*T, error)
	GetByKey(conf GetByKeyConf) (*T, error)
	List(conf ListConf) ([]T, error)
	Search(SearchConf) ([]T, error)
	Update(conf UpdateConf) error
	Inactivate(conf InactivateConf) error
}
