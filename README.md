# Rest_API for instagram clone

### Tools used
- Golang
- mongoDB

## File Structure

- /config 
    - db connection and returns collection based on parameter
- /controller
    - handles the requests and responses
- /models
    - defining the structure of models
- main
    - main file of app


## ROUTERS BY net/http(based on methods switch case is used)

```go
func UserHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method{
		case http.MethodPost:
			CreateUser(w,r)
		case http.MethodGet:
			id:= r.URL.Query().Get("id")
			if(id==""){
				GetUsers(w,r)
			}else{
				GetUserById(w,r)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
```

## Encryption of password (using bycrypt) 
-  when user signsIn we can use bycrypt.CompareHashed to verify password
```go
func HashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("failed to hash password: %w", err)
    }
    return string(hashedPassword), nil
}
```

## Pagination 
- done using limit function of mongo
- limit is passed as a query param to the url


## EndPoints

### for user
```js
    GET USERS = domain/users/        
    GET USER BY ID = domain/users/id (id as param)
    CREATE USER = domain/users/       (POST)

    GET USERS WITH LIMIT = domain/users/?limit=1 
```

### for posts
```js
    GET POSTS = domain/posts/        
    GET POST BY ID = domain/posts/id (id as param)
    CREATE POST = domain/posts/       (POST)

    GET POSTS WITH LIMIT = domain/posts/?limit=1 
```

### posts by user
```js
    GET POST BY USER = domain/posts/users/id (id as param)

    GET POST BY USERS WITH LIMIT = domain/posts/users/?limit=1 

```


## Responses of END Points

### GET USERS
### http://localhost:8081/users/

```yaml
[
  {
    "_id": "6161ca67ad7cdad5b05ac2e4",
    "name": "Karthik",
    "email": "karthik@gmail.com"
  },
  {
    "_id": "6161caacad7cdad5b05ac2e5",
    "name": "Arjun",
    "email": "Arjun@gmail.com"
  }
]
```

### GET USERS using limit=1
### http://localhost:8081/users/?limit=1
```yaml
[
  {
    "_id": "6161ca67ad7cdad5b05ac2e4",
    "name": "Karthik",
    "email": "karthik@gmail.com"
  }
]
```

## GET USER BY ID
### http://localhost:8081/users/?id=6161ca67ad7cdad5b05ac2e4
```yaml
{
  "_id": "6161ca67ad7cdad5b05ac2e4",
  "name": "Karthik",
  "email": "karthik@gmail.com",
}
```

## GET POSTS
### http://localhost:8081/posts/
```yaml
[
  {
    "_id": "6161cb92ad7cdad5b05ac2e6",
    "UID": "6161ca67ad7cdad5b05ac2e4",
    "caption": "test1",
    "image_url": "http://test.img",
    "created_at": "time.Date(2021, time.October, 9, 22, 34, 18, 211456000, time.Local)"
  },
  {
    "_id": "6161cb9ead7cdad5b05ac2e7",
    "UID": "6161ca67ad7cdad5b05ac2e4",
    "caption": "test2",
    "image_url": "http://test2.img",
    "created_at": "time.Date(2021, time.October, 9, 22, 34, 30, 362991000, time.Local)"
  },
  {
    "_id": "6161cbc5ad7cdad5b05ac2e8",
    "UID": "6161caacad7cdad5b05ac2e5",
    "caption": "test3",
    "image_url": "http://test3.img",
    "created_at": "time.Date(2021, time.October, 9, 22, 35, 9, 536000000, time.Local)"
  }
]
```

### GET POST WITH LIMIT=2
### http://localhost:8081/posts/?limit=2
```yaml
[
  {
    "_id": "6161cb92ad7cdad5b05ac2e6",
    "UID": "6161ca67ad7cdad5b05ac2e4",
    "caption": "test1",
    "image_url": "http://test.img",
    "created_at": "time.Date(2021, time.October, 9, 22, 34, 18, 211456000, time.Local)"
  },
  {
    "_id": "6161cb9ead7cdad5b05ac2e7",
    "UID": "6161ca67ad7cdad5b05ac2e4",
    "caption": "test2",
    "image_url": "http://test2.img",
    "created_at": "time.Date(2021, time.October, 9, 22, 34, 30, 362991000, time.Local)"
  }
]
```

### GET POST BY ID
### http://localhost:8081/posts/users/?id=6161ca67ad7cdad5b05ac2e4
```yaml
{
  "_id": "6161cb92ad7cdad5b05ac2e6",
  "UID": "6161ca67ad7cdad5b05ac2e4",
  "caption": "test1",
  "image_url": "http://test.img",
  "created_at": "time.Date(2021, time.October, 9, 22, 34, 18, 211456000, time.Local)"
}
```

### GET POST BY USER
### http://localhost:8081/posts/users/?id=6161ca67ad7cdad5b05ac2e4
```yaml
[
  {
    "_id": "6161cb92ad7cdad5b05ac2e6",
    "UID": "6161ca67ad7cdad5b05ac2e4",
    "caption": "test1",
    "image_url": "http://test.img",
    "created_at": "time.Date(2021, time.October, 9, 22, 34, 18, 211456000, time.Local)"
  },
  {
    "_id": "6161cb9ead7cdad5b05ac2e7",
    "UID": "6161ca67ad7cdad5b05ac2e4",
    "caption": "test2",
    "image_url": "http://test2.img",
    "created_at": "time.Date(2021, time.October, 9, 22, 34, 30, 362991000, time.Local)"
  }
]
```