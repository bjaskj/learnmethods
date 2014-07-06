// base handler to compare request string with CompareWith, will not call next if matched
package handlers

import "fmt"

type CompareHandler struct {
	CompareWith string
}

func (h *CompareHandler) Process(request string) (response string, next bool) {
	if request == h.CompareWith {
		return fmt.Sprintf("Input '%s' matches '%s'", request, h.CompareWith), false
	}

	return fmt.Sprintf("Input '%s' does not match '%s'", request, h.CompareWith), true
}
