package mock

type RetrieverPoster struct {
	Contents string
}

func (r RetrieverPoster) Get(url string) string {
	return "Get:" + r.Contents
}

func (r RetrieverPoster) Post(url string) string {
	return "Post:" + r.Contents
}
