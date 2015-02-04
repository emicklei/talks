func StatsdFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	now := time.Now()
	chain.ProcessFilter(req, resp)
	label := req.Attribute("metric")
	if label != nil {
		statsdClient.Timing(label.(string), time.Now().Sub(now).Nanoseconds()/1E6, 1.0) // diff in ms
	}
}