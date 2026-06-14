package resty

type Request struct {
	// ... existing fields
}

func (r *Request) isMultipart() bool {
	return r.isMultipart || len(r.Files) > 0 || len(r.FormData) > 0
}