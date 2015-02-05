ws.Route(ws.GET("/{name}/messages").To(m.leaseOrPeekMessages).
	// docs
	Doc("Lease a set of messages from the queue").
	Operation("leaseOrPeekMessages").
	Param(nameParam).
	
	Param(ws.QueryParameter("peek",
	"if true then only have a look at the messages. It does not set the visiblity timeout and you do not get a messageHandle. (optional,default=false)").DataType("boolean")).
	
	Param(ws.QueryParameter("limit",
	"limit on the number of messages fetched. (optional, default = 10)").DataType("int")).
	
	Param(ws.QueryParameter("visibilityTimeout",
	"number of seconds the messages are visible for processing; they are hidden to other consumers. (optional, default = 300)")).
	
	Writes([]model.StoredMessage{}).
	Do(Returns200, Returns400, Returns500))