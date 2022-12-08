<div id="top"></div>


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://i.imgur.com/p45abdy.png" alt="Logo" width="180">
  </a>

<h3 align="center">Go Lang - Base</h3>
</div>


<p align="right">(<a href="#top">back to top</a>)</p>

## Packages

- https://github.com/gin-gonic/gin
- golang.org/x/crypto/bcrypt
- https://github.com/golang-migrate
- https://gorm.io/
- https://github.com/golang-jwt/jwt

## Commands

- Start Server
  ```
  go run <project_name>
  
  e.g: go run base
  ```

- Create Migration File

    ```
    migrate create -ext sql -dir DB/migrations -seq create_users_table
    ```

- Migrate Up
   ```
   go run DB/migrate.go up
   ```
- Migrate Down
   ```
   go run DB/migrate.go down
   ```

- Seeds
   ```
   go run DB/seeder.go
   ```

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D

[Vue-url]: https://vuejs.org/

[Golang-url]: https://go.dev/

[Golang]: https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white