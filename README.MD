


# Updated schema.graphqls?

```
go run github.com/99designs/gqlgen generate
```

# Disclaimer

This is the clean architecture approach used by the Fullcycle college, but I don't really feel like it is "clean", 
at the start of the project, there's so much code involved just for a single use case (adding order), things could be simpler.
This might feel like a maze for people who are just starting, so from my perspective, this goes against what the book
clean architecture preaches.

# Challenge

We get a new order and we have to calculate the price and the interest

# About the App
This application is running at the port 8080
To run the application first start up the dependencies with a ``docker-compose up -d
And then ``go run .
To make a http request, you can either use the documents inside the api folder, or a tool like postman
The graphql playground is located at localhost:8080/graphql
