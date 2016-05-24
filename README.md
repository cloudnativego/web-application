[![wercker status](https://app.wercker.com/status/8937b0ba294453fedb0616fcc66201bb/m "wercker status")](https://app.wercker.com/project/bykey/8937b0ba294453fedb0616fcc66201bb)

# Web Application
A simple illustration of a web application that hosts static content as well as facilitates some
JavaScript AJAX style consumption of a service hosted in Go.

| Resource | Method | Description |
|---|---|---|
| / | GET | Home page (index.html) |
| /api/test | GET | Simple REST API that can be consumed from JavaScript |
| /cookies/write | GET | Writes cookies to the browser |
| /cookies/read | GET | Reads previously written cookies |
