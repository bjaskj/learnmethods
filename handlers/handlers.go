package handlers

type Handler interface {
	Process(request string) (response string, next bool)
}
