# gobahnhof

* Author: [https://github.com/benpolzin](Ben Polzin)
* Github Repo: https://github.com/benpolzin/gobahnhof

Thanks to:
* https://golang.org/doc/articles/wiki/
* http://astaxie.gitbooks.io/build-web-application-with-golang/
* http://gohugo.io/templates/go-templates/
* https://golang.org/doc/effective_go.html#web_server

I am using this project to [http://www.golang-book.com/](learn Go) and general software development concepts.

## Goals for gobahnhof
1. Standalone web app that engineers can run on a laptop and capture PRT files from Cisco Jabber and Cisco endpoints that support PRT.
2. Page to view captured PRT submissions and search/sort those submissions.
3. Capture manually uploaded files (PRT or other).
4. Optional database backend for more permanent installations.

## Usage
### Manual Upload
*URL: ./go/upload/
### Jabber PRT
*URL: ./jabberprt/
*Example: Use http://prt.widgets.co/jabberprt/ in the Cisco PSC field for Problem Report Server URL
### Phone PRT
*URL: ./phoneprt/
*Example: Use http://prt.widgets.co/phoneprt/ in the Cisco PSC field for Customer Support Upload URL

## Version 0.1
* Basic web server implemented
* Can upload a file manually
* Can capture PRTs from Jabber
* Can capture PRTs from phones