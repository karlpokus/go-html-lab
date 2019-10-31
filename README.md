# go-html-lab
How feasable is go for serving web pages and rendering html? We'll start from an old [blog post](https://golang.org/doc/articles/wiki/) and work our way from there.

# api
- http://localhost:8080/view/<page>
	if NotFound
		redirects to /edit/<page>
- http://localhost:8080/edit/<page>
	on submit
		call /save/<page>
			which redirects to /view/<page>
- http://localhost:8080/save/<page>
		redirects to /view/<page>

# todos
- [x] add httprouter
- [ ] only admin can edit
- [ ] check out template syntax here -> https://github.com/golang/blog/blob/master/template/root.tmpl

# license
MIT
