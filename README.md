### Functionality
This is a CLI program uses the GitHub API
It needs to perform few basic functions:

1. Fetch a user's or orgs profile info:  
   GET `curl https://api.github.com/users/matthausen`
2.  List all the Github Repos given a username:  
    GET `curl https://api.github.com/users/matthausen/repos`
3. Authenticate `https://api.github.com/user/repos?page=1&per_page=1000` + Header `Authorization: token <token>`
4. View private repos
5. Create a repo:  
    POST `curl \
   -X POST \
   -H "Authorization: token <access_token>" \
   https://api.github.com/user/repos \
   -d '{"name":"test"}'`
   
5. Delete a repo
    a. create a token with delete_scope: `curl -v -u username -X POST https://api.github.com/authorizations -d '{"scopes":["delete_repo"], "note":"token with delete repo scope"}'`
    b. delete the repo `curl -X DELETE -H 'Authorization: token xxx' https://api.github.com/repos/:owner/:repo
   `

### Initialize a cobra CLI:
Initialise a normal project:
- `go mod init github/...`
  
Initialise a cobra project with package name: 
- `cobra init --pkg-name <pkg-name>`
  
Add commands:
- `cobra add <command name>`
  
Compile and add to terminal commands:
- `go install <pkg-name>`