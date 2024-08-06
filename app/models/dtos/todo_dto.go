package dtos

type TodoDto struct {
	Id    int64
	Title string
	Tags  []*TagDto
}
