package proxy

type iServer interface {
	handleRequest(string, string) (int, string)
}
