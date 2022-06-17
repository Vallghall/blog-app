# BlogApp

## API Reference

### Auth API:

- POST "/auth/sign-up"; JSON-body example:
  ```
    {
      "name": "Roman",
      "surname": "Gusev",
      "father_name": "Mikhailovich",
      "nickname": "Vallghall",
      "password": "123456",
      "confirm_password": "123456"
    }
  ```
- f
- f

### Blog API:

- POST "/blog/:user_id/:num"; JSON-body example:
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