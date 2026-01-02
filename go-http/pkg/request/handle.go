package request


func HandleBody[T any](w http.ResponseWriter, req *http.Request) (interface{}, error) (*T, error) {{
	var payload LoginRequest
        body, err := Decode(req.Body)
		if err != nil {
			response.JsonResponse(w, err.Error(), 402)
			return nil, err
		}
		err = IsValid(body)
		if err != nil {
			response.JsonResponse(w, err.Error(), 402)
			return nil, err
		}
		return body, nil
}