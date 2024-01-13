# server-sent-events
This repository is used to support [this](https://www.prakharsrivastav.com/posts/server-sent-events/) blog post. The interested readers can clone the repository and follow below steps to run the application. The code base has comments at the relevant places to help readers understand the logic.


## Running application
```shell
# from the project root run
go run main.go
```

This starts up an http server on port 8080 . 
- Open the browser and, navigate to [http://localhost:8080](http://localhost:8080).
- Click on the big blue button.
- The server should start streaming the response to browser.
- Browser will be updated the realtime response.


If you have any questions, please create an issue and I will be more than happy to answer the queries.