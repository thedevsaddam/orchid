Orchid
=================
Golang MVC style boilerplate using gin-gonic framework with gorm ORM.
![Orchid](public/orchid.jpg)
### Installation
* Go to your `$GOPATH/src` and clone the directory using `git clone https://github.com/thedevsaddam/orchid.git`
or [download the zip file](https://github.com/thedevsaddam/orchid/archive/master.zip)

* Install dependency manager `govendor` using the command below
    ```
    go get -u github.com/kardianos/govendor
    ```

* Go to the `$GOPATH/src/orchid` and install dependencies using `govendor sync` command

* Copy `.env.example` to `.env` and set your configurations.

* Run `go build` to build binary file and to start the application use `./orchid`


### Todo
- [ ] Session
- [ ] OAuth2
- [ ] Fixing inconstant codes
- [ ] Request validation
- [ ] Add some helper function
- [ ] Security
- [ ] Find out performance issues
- [ ] Benchmarking


### Credits
* Routing and templating [gin-gonic](https://gin-gonic.github.io/gin)
* Object-relational mapping [gorm](https://github.com/jinzhu/gorm)
* Dependency management package [govendor](https://github.com/kardianos/govendor)
* Environment management  package [godotenv](https://github.com/joho/godotenv)

### Contribution
As it is the very begining of *Orchid* boilerplate, we need a good amount for contribution from you.
If you are interested to contribute please feel free to send pull request.

### License
The **Orchid** is a open-source software licensed under the [MIT License](LICENSE.md).
