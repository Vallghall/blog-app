# BlogApp

## Blog server API Reference

### Auth API:

- POST "/auth/sign-up" - registration endpoint; JSON-body example:
  ```
    {
      "name": "Ivan",
      "surname": "Ivanov",
      "father_name": "Ivanovich", // optional
      "nickname": "BearRider",
      "password": "123456",
      "confirm_password": "123456"
    }
  ```
  in response it returns new user's `user_id`.
- POST "/auth/sign-in" - login endpoint; JSON-body example:
  ```
  {
    "nickname": "BearRider",
    "password": "123456"
  }
  ```
  in response it sets refresh-token-cookie and returns `user_id` and `token`.
- POST "/auth/refresh" - token pair refreshment endpoint. If refresh token is valid
and unexpired, and the expired access token is also valid, it returns a new token pair.
- POST "/auth/logout" - logout endpoint. It resets the refresh-token stored in cookie.

### Blog API:

- POST "/blog/:user_id" - endpoint for creating blog posts;  JSON-body example:
  ```
    {
      "title": "Initial",
      "content": "Here's my post, bugaga",
      "hashtags": [
          "starters",
          "initials",
          "letsGOooo"
      ]
  }
  ```
  - GET "/blog/:user_id/:num" - endpoint for receiving the `:num` amount of user's
  posts.
  - PUT "/blog/:post_id" - endpoint for editing blog posts;  JSON-body example:
  ```
    {
      "title": "Initial",
      "content": "Here's my post, bugaga",
      "hashtags": [
          "starters",
          "initials",
          "letsGOooo"
      ]
  }
  ```
  - DELETE "/blog/:post_id" - endpoint for deleting blog posts.
  - GET "/blog/post/:post_id" - endpoint for reading an exact blog post.