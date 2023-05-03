# GoCommunityAPI

GoCommunityAPI is a backend API service for a community website. It provides various endpoints to manage users, articles, and comments. This service can be built and deployed using Docker Compose.

## API Endpoints

### **User**

#### Retrieve User Information

`GET /api/user/:email`

Retrieve a user's information by their email address.

#### Create User

`POST /api/user/`

Create a new user.

#### User Login

`POST /api/user/login`

Authenticate a user and generate an access token.

### **Article**

#### Fetch Article Page

`GET /api/article`

Retrieve a list of articles by page.

#### Retrieve Article Information

`GET /api/article/:id`

Retrieve an article's information by its ID.

#### Upload Article

`POST /api/article/`

Upload a new article.

#### Update Article

`PUT /api/article/:id`

Update an existing article by its ID.

#### Delete Article

`DELETE /api/article/:id`

Delete an existing article by its ID.

### **Comment**

#### Fetch Comment Page

`GET /api/comment/articleId/:articleId`

Retrieve a list of all comments for a specific article.

#### Upload Comment

`POST /api/comment/`

Upload a new comment for a specific article.

#### Delete Comment

`DELETE /api/comment/:id`

Delete an existing comment by its ID.

## Getting Started

To build and run this project, you will need to have Docker and Docker Compose installed on your machine.

1. Clone this repository:

```
git clone <repository-url>
```

2. Navigate to the project directory:

```
cd <project-directory>
```

3. Start the Docker Compose services:  
This will create the API service and a MYSQL database[root:12345@tcp(communitymysql:3306)].  
```
docker-compose up
```

4. Access the API endpoints via `http://localhost:3000`.

## Conclusion

This README.md file provides an overview of the GoCommunityAPI service and its API endpoints. By following the Getting Started instructions, you can quickly build and run this service on your local machine.