# golunch
Not feeling like deciding where to eat? Let a random lunch picker do it for you! This is a tool that, given a keyword and optional location, uses Yelp's API to search for nearby restaurants and picks a random one.

## Installation
`go get github.com/tracehelms/golunch`
`go install`

You will need to get API keys from [Yelp](https://www.yelp.com/developers). Once you have them, create a file called `yelp_keys.json`. Structure it like this:

```
{
  "ConsumerKey": "",
  "ConsumerSecret": "",
  "AccessToken": "",
  "AccessTokenSecret": ""
}
```

And fill in the values with your API keys from Yelp.

### Known Issue
You will have to have this file in whatever directory you want to run the command in. If you want to run it from your home directory, make this file in the `~/` directory. I'm still working out how to let this be set in one place and have the command work from anywhere.


## Usage
From the terminal, run:
`$ golunch -location=zip query`

The default location is 80304, Boulder CO.

### Examples
```
$ golunch -location=80301 mexican
Time for grub! Your destination:
Name:      Sancho's Authentic Mexican Restaurant
Rating:    4.5
Location:  [2850 Iris Ave Ste H Boulder, CO 80301]
URL:       http://www.yelp.com/biz/sanchos-authentic-mexican-restaurant-boulder
```

```
$ golunch thai
Time for grub! Your destination:
Name:      Terra Thai
Rating:    4.5
Location:  [1121 Broadway Ste 103 Boulder, CO 80302]
URL:       http://www.yelp.com/biz/terra-thai-boulder
```
