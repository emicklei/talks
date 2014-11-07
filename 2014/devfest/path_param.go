	func (r *Request) PathParameter(name string) string { // HL
		return r.pathParameters[name]
	}