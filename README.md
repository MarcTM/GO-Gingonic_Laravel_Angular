<h1>2DAW PROJECT WITH GO, LARAVEL AND ANGULAR</h1>

This web page works with 2 backends (Go and Laravel) and the frontend (Angular11).  
Poject developed by *Marc Torres Martínez*.
<br>
With this project I was trying to do a page where a community of persons can share their own recipes, follow people, favorite other recipes, etc; and at the same time a shop, so users can buy the meals/tuppers that it offers.
---

## Getting started

If you want to run this repository locally:

- Clone this repository to a folder
- Install mySQL and run it. Also create the database that you are going to use for the application.
- Search for the files with database connections, redis, etc, and change them.
- Install angular packages, and go and laravel dependences (npm install, go get, etc). Also migrate laravel tables with php artisan migrate.
- Now you can execute the frontend with **npm start**, the go backend with **go run main.go** inside each microservice, and laravel with **php artisan serve**.

You can also run the application in docker:
 - sudo docker-compose up --build


---

## Features

| Page | Features |
| - | - |
| Home | The home shows a carousel, the categories and the user with more recipes |
| Recipes | Lists all the user's recipes paginated and ordered by time of creation |
| Meals | Lists the shop paginated |
| Recipes Editor | You can create a recipe, only if logged id |
| Meals Editor | You can create a product, only if logged in and if you are an admin |
| Profile | Lists a profile and his recipes |
| Account | Lists your profile and your recipes. You can update here your profile |


<br>

| Services | Features |
| - | - |
| Register | You can register manualy |
| Login | You can login manually |
| Favorites | You can favorite recipes |
| Follow | You can follow users |

<br>

| Technical Features |  |
| - | - |
| Authentication | You need to be signed in (authenticated) to be allowed to do some actions |
| Go Microservices | Go backend is divided into two microservices (users and recipes) |
| Admin features | A user that is an admin can do other things, like create a category and a meal |

---

### Technologies used in this project

* Go: Golang and Gorm
* Angular 11
* Laravel
* Docker
* mySQL

### Other technologies used

* JWT 
* Toaster
* Docker Compose

---

## Backend Application Structure

- `go` - There are features up to all the users: see recipes, see profiles, follow user(if logged), favorite recipe(if logged), create recipe(if logged), login, register.
- `laravel` - There are features related to things that the admins can controll: create catgory (only admin), create meal (only admin), show carousel, show meals.

## Frontend Application Structure

- `angular` - The frontend works with angular and has some features like the AuthGuard or AdminGuard so you cannot access some routes that require it. You store two bearers, one from Go to know that you are logged in, and the other from Laravel to know that you are an admin.dule.
---
