### A distributed library-management app

*api-gateway - typescript nodejs
    - all request goes through this
    - and it's resposible for orchestrating authentication and authorization

*user-auth - typescript nodejs
    - user-auth service maintains users and authentication and authorization

*book - golang 
    - book service maintains books and authors

*library - golang
    - library service maintains book loans


###### Request flow => client -> api-gateway -> user-auth(authentication & authorization) -> proxy to other service
###### TODO => 
    1.Improve excel export by using worker process to generate excel and generate at timeinterval ex-hourly, 
    save the file and serve that saved file, NOT ON REQUEST TRANSACTION
