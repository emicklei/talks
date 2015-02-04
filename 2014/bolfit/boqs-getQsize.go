import "github.com/emicklei/go-restful"

//START OMIT
func (r QueueResource) getQueueSize(request *restful.Request, response *restful.Response) {
	request.SetAttribute("metric", "getQueueSize")
	name := request.PathParameter("name")
	size, err := r.dao.QueueSize(name)
	if err != nil {
		handleServerError(err, request, response)
		return
	}
	response.WriteEntity(size)
}

//END OMIT