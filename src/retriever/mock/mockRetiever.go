package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string{
	return "Get:" + r.Contents
}

func (r Retriever) Post(url string) string{
	return "Post:" + r.Contents
}
