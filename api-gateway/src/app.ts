import Express from "express";
import httpProxy from "http-proxy";
import morgan from "morgan";
import config from "./configuration";
import { USER_AUTH_URLS } from "./service-proxy/user-auth"
import { isAuthenticated } from "./middleware/authenticated";
import { getUserFromToken } from "./service/authService";

const app = Express()

const proxy = httpProxy.createProxyServer();

app.use(morgan('combined'));

app.get('/', (req, res) => res.send('Hello World!'))

app.post("/api/v1/auth/login", (req, res) => {
    proxy.web(req, res, { target: USER_AUTH_URLS.login, prependPath: false});
})

app.post("/api/v1/auth/register", (req, res) => {
    proxy.web(req, res, { target: USER_AUTH_URLS.register, prependPath: false});
})

app.get("/api/v1/user", isAuthenticated, async (req, res) : Promise<any> => {
    try {
        const authHeader = req.headers['authorization']
        const user = await getUserFromToken(authHeader as string);
        console.log(user);

        req.headers["X-ROLE"] = user.role;
        proxy.web(req, res, { target: USER_AUTH_URLS.listUser, prependPath: false});
    } catch(e) {
        console.log(e);
        return res.status(401).send({
            message: e.message ? e.message : "unauthorized",
        });
    }
    
})

app.get("/api/v1/user/{userID}", isAuthenticated, (req, res) => {
    proxy.web(req, res, { target: USER_AUTH_URLS.getUser(req.params.userID), prependPath: false});
})

export default app;