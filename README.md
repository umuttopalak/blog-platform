# Blog Platform

## Description
A blog platform where users can create, edit, and delete blog posts. Users can also comment on posts and like them.

## Features
- User authentication (sign up, log in, log out)
- Create, edit, delete blog posts
- Comment on posts
- Like posts
- User profiles
- Admin roles and permissions
- Categories for posts

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/blog-platform.git
    ```
2. Navigate to the project directory:
    ```sh
    cd blog-platform
    ```
3. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Start the development server:
    ```sh
    go run main.go
    ```
2. Open your browser and navigate to `http://localhost:8080`

## API Endpoints

### User Routes
- `POST /users/register` - Register a new user
- `POST /users/login` - Login a user

### Post Routes
- `POST /posts` - Create a new post
- `GET /posts` - Get all posts
- `GET /posts/:post_id` - Get a specific post
- `PUT /posts/:post_id` - Update a specific post
- `DELETE /posts/:post_id` - Delete a specific post

### Comment Routes
- `GET /comments/user` - Get comments by the logged-in user
- `POST /comments/:post_id` - Create a comment on a specific post
- `GET /comments/post/:post_id` - Get comments on a specific post
- `PUT /comments/:comment_id` - Update a specific comment
- `DELETE /comments/:comment_id` - Delete a specific comment

### Reaction Routes
- `POST /reactions` - Add a reaction to a post or comment
- `GET /reactions/post/:post_id` - Get reactions on a specific post
- `GET /reactions/comment/:comment_id` - Get reactions on a specific comment
- `DELETE /reactions/:reaction_id` - Remove a specific reaction

### Category Routes
- `GET /category` - Get all categories
- `POST /category` - Create a new category
- `GET /category/:category_id` - Get a specific category

### Admin Routes
- `POST /admin/role/add` - Add a role to a user
- `POST /admin/role/create` - Create a new role
- `DELETE /admin/role/remove/:role_id` - Remove a specific role
- `POST /admin/role/remove-from-user` - Remove a role from a user

## Contributing
1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Make your changes
4. Commit your changes (`git commit -m 'Add some feature'`)
5. Push to the branch (`git push origin feature-branch`)
6. Open a pull request

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.