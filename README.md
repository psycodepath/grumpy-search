# Grumpy Search

## About
This is a proof of concept for a claim i did on [mastodon](https://mastodon.social/@psycodepath/109796986115923657). 
As i always wanted to try an easy to deploy search for my projects

### Usage
1. build the index ```bash  make index``` (*hint make sure there is no previous index called grumpy.bleve here*)
2. run the api ```bash make run  ```
3. run a search query [http://localhost:1337/?search=monkey](http://localhost:1337/?search=monkey]) and have a look at the json response


### Caution
This is not production ready code. There is no tls setup for the api nor is this fully tested. It is just a proof 
of concept. 

### License
Copyright 2023 Steffen Frosch

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
