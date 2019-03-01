package safe

import "io"

// Close gracefully closes the Closer and calls the cleaners if an error occurred.
//
//  func handler(rw http.ResponseWriter, req *http.Request) {
//
//  	defer Close(req.Body, func(err error) { log.Println(err) })
//
//  	var data map[string]interface{}
//  	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
//  		rw.WriteHeader(http.StatusBadRequest)
//  		return
//  	}
//
//  	...
//  }
//
func Close(closer io.Closer, cleaners ...func(error)) {
	if err := closer.Close(); err != nil {
		for _, clean := range cleaners {
			clean(err)
		}
	}
}
