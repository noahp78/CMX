#This project is in 'semi-active' development. 
And currently should not be used for production.

# CMX
Content Manager eXtended


### What is CMX?
CMX is a web application framework written in go, made to quickly develop new and exciting web ideas.
At it's core it's a very simple HTTP/Web server, but thanks to a whole host of core libraries it offers 
features ranging from Authentication to ORM support.


### How does it work?
CMX is designed to be extensible, secure, and modular.
We have a few modules that ship with default CMX.
- core
    - core/data
        - Provide persistant storage to your application
    - core/pages
        - Templating Engine
        - More Spec needs to be defined
    - core/util
        - Utils used in the entire application
    
- addons/accounts
    - Provides login / user functionality
- addons/dashboard
    -  Default dashboard that ships with CMX.
- core/pages
    - Basic page system, allows you to define models, and dynamicly render pages.

You can configure the addons that are enabled for your  setup of CMX by updating 
```
addons/main.go
```
It makes use of the following, excellent Open Source Libraries

- https://github.com/gin-gonic/gin
- https://github.com/aymerick/raymond
