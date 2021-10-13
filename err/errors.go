package err

type UnauthorizedError struct {
	Msg string
}

func (p *UnauthorizedError) Error() string {
	return p.Msg
}
