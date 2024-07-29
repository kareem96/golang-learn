package simple

import "errors"

type SimpleReposiotry struct {
	Error bool	
}

type SimpleService struct {
	*SimpleReposiotry
}


func NewSimpleReposiotry(isError bool) *SimpleReposiotry {
	return &SimpleReposiotry{
		Error: isError,
	}
}

func NewSimpleService(simpleReposiotry *SimpleReposiotry) (*SimpleService, error) {
	if simpleReposiotry.Error{
		return nil, errors.New("Failed create service")
	}else{
		return &SimpleService{SimpleReposiotry: simpleReposiotry}, nil
	}
}

