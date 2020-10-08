# golang-snippetbox
A web app which lets people paste and share snippets of text, similar to Pastebin or GitHub Gists. 

## Key learning
1. How to use `gomod` to manage dependencies in Go. Which is a big change. In the past I always needed to work within an `src` folder under the `GOPATH` directory but with `gomod` you explicitly work in anywhere but the `GOPATH` it seems.
2. How to add new routes to a Go server. First create a handler function then register the handler function into the servemux mapping it to the route you want it to route from.
3. Distinguish between fixed path and subtree URL patterns.
4. How to customize HTTP headers when responding to client.
5. How to get handler to accept query string parameters from the user and validate/sanitize and use it.
6. Project structure best practices based on [Peter Bourgon's Go Best Practices.](https://peter.bourgon.org/go-best-practices-2016/#repository-structure)
7. Learn HTML templating and template composition (quite a bit like with React components) and how to read, parse and execute template sets in the relevant handler. So now we can return HTML to the user not just plaintext and JSON.
