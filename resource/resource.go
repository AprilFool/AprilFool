package resource

type Resource interface {
	Name() string
}

type GetSupported interface {
	Resource
	Get(map[string][]string) (interface{}, error)
}

type PostSupported interface {
	Resource
	Post(map[string][]string) (interface{}, error)
}

type PutSupported interface {
	Resource
	Put(map[string][]string) (interface{}, error)
}

type DeleteSupported interface {
	Resource
	Delete(map[string][]string) (interface{}, error)
}

type HeadSupported interface {
	Resource
	Head(map[string][]string) (interface{}, error)
}

type PatchSupported interface {
	Resource
	Patch(map[string][]string) (interface{}, error)
}
