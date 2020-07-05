package page

type DuoUIpage struct {
	Name    string
	Command func()
	Data    interface{}
}

type DuoUIpages *map[string]*DuoUIpage
