# GRUUUB!
Not feeling like deciding where to eat? Let gruuub do it for you! This is a tool that, given a keyword and location, uses Yelp's API to search for nearby restaurants and picks a random one.

Find it live at http://www.gruuub.com.


## Contributing
### Installation
- First, fork the repo
- `go get github.com/your_github_username/gruuub`
- Change into the project directory: `cd $GOPATH/src/your_github_username/gruuub/`
- You may have to change the import of `github.com/tracehelms/gruuub/yelp` in `main.go` to your GitHub username.
- `go get`
- `go build`

You will need to get API keys from [Yelp](https://www.yelp.com/developers). Once you have them, create a file called `yelp_keys.json` in your project directory. Structure it like this:

```
{
  "ConsumerKey": "",
  "ConsumerSecret": "",
  "AccessToken": "",
  "AccessTokenSecret": ""
}
```

And fill in the values with your API keys from Yelp.

### Running
From the project directory, run:
- `./gruuub`

Now browse to http://localhost:8080.

When you make changes to the Go code, you'll have to close the server `Cmd + C`, run `go build`, and start the server again. HTML / CSS changes should be reflected without restarting the server.

### Pull Requests
Make your changes, commit, and then make a pull request on GitHub.


## Notes
### Presentation
This project was originally the result of a talk I gave at the [Boulder Gophers Meetup](http://www.meetup.com/Boulder-Gophers). The slides can be found [here](http://tracehelms.com/talk_consuming_apis_with_go). It was originally a command line app.

### Sources
- https://github.com/JustinBeckwith/go-yelp
