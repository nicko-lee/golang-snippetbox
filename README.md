# golang-snippetbox
A web app which lets people paste and share snippets of text, similar to Pastebin or GitHub Gists. 

## Key learning
1. How to use `gomod` to manage dependencies in Go. Which is a big change. In the past I always needed to work within an `src` folder under the `GOPATH` directory but with `gomod` you explicitly work in anywhere but the `GOPATH` it seems.
2. How to add new routes to a Go server. First create a handler function then register the handler function into the servemux mapping it to the route you want it to route from.
3. Distinguish between fixed path and subtree URL patterns.
