package foo

import "fmt"

// Service ...
type Service struct {
	props Props
}

// Props ...
type Props struct {
	MyPrivateKey string
	Name         string
}

// New ...
func New(props Props) *Service {
	return &Service{
		props: props,
	}
}

// Props ...
func (s *Service) Props() Props {
	return s.props
}

// UsePrivateKey ...
func (s *Service) UsePrivateKey() {
	fmt.Println(len(s.props.MyPrivateKey))
}
