package script

type Runner interface {
	Arg(i interface{}) error
	Ret() ([]byte, error)
	Run() error
}
