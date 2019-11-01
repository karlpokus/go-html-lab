# go-html-lab
How feasable is go for serving web pages and rendering html? We'll start from an old [blog post](https://golang.org/doc/articles/wiki/) and work our way from there.

Unrelated - WTF can you not send a cookie along with a redirect?!
- Awkward solution 1: send token in query string of the redirect path. This creates a problem for auth middleware.
- Awkward solution 2: /login POSTs to /view where a valid user will get html + token cookie set

I picked 2 for now.

# api
All routes that require a token redirect to `/login` if the token missing or invalid.

###### /
Displays the home page.

###### /login
Displays a login form that POSTs to `/view`.

###### /view
Requires valid user credentials. Displays something on success. Otherwise returns 401.

###### /view/:page
Requires a token. Displays a page. Redirects to `/edit/:page` if the page is not found.

###### /edit/:page
Requires a token. Displays an editing form that POSTs to `/save/:page`.

###### /save/:page
Requires a token. Redirects to `/view/:page` on success.

# todos
- [x] add httprouter
- [ ] avoid disk io w loadPage on each call
- [x] auth token
- [ ] user handling
- [ ] [jwt](https://github.com/dgrijalva/jwt-go)
- [ ] [bcrypt](https://gowebexamples.com/password-hashing/)
- [ ] prefix api call endpoints
- [ ] logout
- [ ] tests?
- [ ] check out template syntax here -> https://github.com/golang/blog/blob/master/template/root.tmpl

# license
MIT
