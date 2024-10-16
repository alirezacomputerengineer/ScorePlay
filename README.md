# ScorePlay Tecnical Task

## Description

Based on Technical Task needs, This service exposes a REST API to perform 4 main functionalities : **Create a Tag**, **List All Tags**, **Create a Media** and **Search Medias by tag**. 

## Assumptions

- **Simplicity**: I avoid complex solution, to make solution easy to undersatnd and usable for all, I design my proposed solution simple. code is totally easy to understand, commented and components are well-named. also using swagger make it more documentable.
- **Generality**: I used **Gin** framework to develop the solution, I could develop this REST API with out extra package, which make the solution a bit complex, or using **Echo** or **Fiber** framework instead of **Gin**, but **Gin** is more global and at this simple context, there is not meaningful difference between frameworks.
- **Productivity**: As mentioned in Technical Task description, application must be production-ready, therefor I avoid persist data using Databases or using Docker, it is simple to run and test. in **Improvements** section I proposed some improvements for market ready product.

## Solution Design

Based on assumptions above, I proposed architecture follows a common and effective pattern for organizing a Go-based web application using the Gin framework. This architecture is centered around the **Separation of Concerns (SoC)** principle, which improves modularity, readability, and maintainability. Each folder in this structure has a specific responsibility:
- `controllers/`: This handles the application logic, responding to HTTP requests, processing them, and sending responses back to the client.
- `models/`: This defines the application's data structures (e.g., `Tag` and `Media`) and serves as the link between business logic and data layer.
- `routes/`: This centralizes route definitions, making it easier to configure and expand routing as the application grows.
- `main.go`: Acts as the entry point for the application. It sets up the server, routes, and middleware (like Swagger) in one place without cluttering it with the business logic or models.

**Advantages of this Architecture** are:

***a. Modularity***: Each part of your application is decoupled from others. The `controllers` focus only on handling requests, `models` are responsible for the data representation, and `routes` are for wiring things up. This makes it easier to maintain, test, and scale the application. If you need to modify or add a new feature (like adding a new API), you know exactly where to look and modify, without worrying about unintended side effects in other parts of the system.

***b. Readability and Maintainability***: By organizing the application into clear directories (`controllers/`, `models/`, and `routes/`), developers unfamiliar with the project can easily understand its structure. The application is easy to navigate and find what you're looking for, which becomes crucial as the codebase grows.

***c. Scalability***: As your project grows and you add more controllers, models, and routes, this structure naturally accommodates expansion. Each new feature can be implemented by adding a new controller and route. You don't need to refactor the entire codebase for adding new features.

***d. Ease of Testing***: Each layer of the application can be tested independently, using either *Unit tests* or *Integreation Test*.

***e. Extensibility***: As the API grows, this architecture can be extended to support additional concerns like *Middlewares*(Custom middleware can be added in the `main.go` or in separate files to handle cross-cutting concerns, somethings like logging, authentication, rate-limiting, etc.) or *Versioning*(API versioning can be easily introduced by organizing controllers/routes into subdirectories like `v1/`, `v2/`, etc.)

Also because of ***Independent components***, ***Loose coupling*** and ***Easy integration*** the proposed structure will work well in **MicroServices Arcitecture**.

### API Structure

API files are structured as this:

```text

myapp/
│   main.go
│   go.mod
└───controllers/
│       tagController.go
│       mediaController.go
└───models/
│       tag.go
│       media.go
└───routes/
        routes.go

``` 

### Dependencies

- [gin](https://gin-gonic.com/) - Gin is a HTTP web framework written in Go (Golang).
- [gin-swagger](https://pkg.go.dev/github.com/swaggo/gin-swagger) - Gin middleware to automatically generate RESTful API documentation with Swagger 2.0.

### Initialize The Project

Download or clone code base(I named the project `myapp`), navigate to the main folder and run these commands.

```bash
go mod init myapp
go get -u github.com/gin-gonic/gin
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/swaggo/swag/cmd/swag
```

Also for Swagger document generation:

```bash
swag init
```

Visit [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html) in your browser to see the Swagger-generated API documentation.


Run the app using this:

```bash
go run main.go
```
## Testing

I developed some test cases as follow:

1. **TestCreateTag**:
- Tests creating a tag with valid input data.
- Ensures that the response contains the correct tag name and a generated ID.

2. **TestCreateTagInvalid**:
- Tests creating a tag with invalid data (missing name).
- Ensures that the request fails with a 400 Bad Request.

3. **TestListTags**:
- Simulates listing all tags.
- Verifies that the API returns all existing tags.

4. **TestCreateMedia**:
- Tests media creation with valid data, including file upload.
- Verifies that the response contains a valid file URL and that the media data is stored correctly.

5. **TestCreateMediaMissingFields**:
- Tests creating media with missing required fields (like name or file).
- Ensures the request fails with a 400 Bad Request.

6. **TestSearchMedia**:
- Tests searching media by tag.
- Ensures the API returns the correct media items associated with a specific tag.

To run all test cases:

```bash
go test ./...
```

## Future Improvements
As I mentioned above, I developed the solution **Easy to Undersatnd** and **Production-Ready**.
Here is some improvments to turn this solution **Ready to Market**:  
1. **Persist data using a database**: To make the solution simple and production ready, I avoid persist data. at this solution the data is stored in memory. To make the data persistent across restarts and usable in a production environment, integrate a relational database like **PostgreSQL**, **MySQL**, or NoSQLs Likes **MongoDB**. to do that, I proposed using **GORM**(ORM for GO). Integrate GORM to interact with the database more easily. This allows you to define relationships, run migrations, and handle CRUD operations seamlessly. it is so easy, Define models with GORM, Replace in-memory operations with GORM database queries and Implement migrations using `db.AutoMigrate(&models.Tag{}, &models.Media{})` for example, we can store media in database:

```go
func CreateMedia(c *gin.Context) {
    var newMedia models.Media
    if err := c.ShouldBind(&newMedia); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Store in database
    if err := db.Create(&newMedia).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create media"})
        return
    }
    c.JSON(http.StatusCreated, newMedia)
}
```

2. **Add Security Features**: Secure the API using **JWT authentication** and **input validation**.
Using **JWT** we can secure our API endpoints and restrict access to certain routes. to activate that, easily Implement user authentication (login), Issue JWT tokens on successful login and Protect routes by verifying JWT in middleware. something like that:
```go
func JWTMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")
        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()
            return
        }

        // Validate token (this is a basic example, you'll want to expand on this)
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("secret_key"), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Next()
    }
}
```
Also use libraries like **go-playground/validator** to validate the input data can prevent lots of attacks to the API. here is an example:
```go
validator := validator.New()
err := validator.Struct(newTag)
if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
    return
}
```
And do not forget **Rate Limiting** strategies, one of good libraries is **golang.org/x/time/rate** or using **Nginx** as reverse proxy.

3. **Use background processing for file uploads** : If you are dealing with large media files, consider offloading file processing (e.g., storing media files, generating thumbnails) to background jobs. I can use **Go channels**, **goroutines**, or utelizing distributed messaging system like **Kafka** or **RabbitMQ** . also use a dedicated background job system **Goraft/Work & Redis** can be a solution.

4. **Add some extra helpful fuctionalities**: Add **update** `e.g. : PUT /tags/:id` and **delete** `e.g. : DELETE /tags/:id` functionalities for tags and media can extend the usefulness. something like these functions:
```go
func UpdateTag(c *gin.Context) {
    id := c.Param("id")
    var tag models.Tag
    if err := db.First(&tag, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
        return
    }

    if err := c.ShouldBindJSON(&tag); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    db.Save(&tag)
    c.JSON(http.StatusOK, tag)
}


func DeleteTag(c *gin.Context) {
    id := c.Param("id")
    if err := db.Delete(&models.Tag{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete tag"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Tag deleted"})
}
```

for listing endpoints `GET /tags` & `GET /media`, we can add **Pagination** and **Filtering**, for example:

```go
func ListTags(c *gin.Context) {
    var tags []models.Tag
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

    offset := (page - 1) * limit
    db.Offset(offset).Limit(limit).Find(&tags)

    c.JSON(http.StatusOK, tags)
}
```
Also we can perform **Filtering Media by Multiple Tags** (`GET /media?tags=tag1,tag2`), easily modify the search functionality to allow filtering by multiple tags, using SQL `IN` queries.

4. **Cache frequently accessed data**: If your media list or search grows, add caching to reduce the load on your database. You can use Redis or an in-memory caching system like Go's **sync.Map**. Example using **Redis**:
```go
import "github.com/go-redis/redis/v8"

// Initialize Redis client
rdb := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
})

// Example caching media search results
func SearchMedia(c *gin.Context) {
    tag := c.Query("tag")
    cachedData, err := rdb.Get(context.Background(), tag).Result()
    
    if err == redis.Nil {
        // Cache miss, perform database query
        var media []models.Media
        db.Where("tags LIKE ?", "%"+tag+"%").Find(&media)
        
        // Store result in Redis
        rdb.Set(context.Background(), tag, media, time.Hour)

        c.JSON(http.StatusOK, media)
    } else {
        // Cache hit, return cached data
        var cachedMedia []models.Media
        json.Unmarshal([]byte(cachedData), &cachedMedia)
        c.JSON(http.StatusOK, cachedMedia)
    }
}
```

5. **Add structured logging for better monitoring and debugging**: To monitor and debug the application more efficient, we can implement structured logging. we can use logging libraries like **logrus** or **zap** to add logs to your API:
```go
import log "github.com/sirupsen/logrus"

log.Info("Creating media record")
log.Error("Failed to upload file")
```
6. **Deploy your API with Docker and Kubernetes for scalability**: and last but not the least, Once our API is feature-complete and secure, consider deploying it using Docker and Kubernetes for scalability. Use a CI/CD pipeline for continuous deployment and testing.
