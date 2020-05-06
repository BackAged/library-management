import express, { Router } from "express";
import bodyParser, { json } from "body-parser";
import morgan from "morgan";

import { initializeDBConnection } from "../../infrastucture/database/mongo";
import { newJsonWebTokenManager } from "../../infrastucture/json-web-token/json-web-token";
import { newUserRepository } from "../../infrastucture/repository/userRepository";
import { newRegisterUserUseCase } from "../../usecase/registerUser/registerUserUseCase";
import { newLoginUserUseCase } from "../../usecase/loginUser/loginUserUseCase";
import { newLoginUserController } from "./controller/loginUserController";
import { newRegisterUserController } from "./controller/registerUserController";

import config from "../../configuration";


const app = express();

// registering app level middleware
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

// bootstrapping the application
(async () => {
    // initialize logger

    // initializing db connection
    const db = await initializeDBConnection(config.MONGO.MONGO_HOST, config.MONGO.MONGO_DB);

    // initializing repos
    const userRepository = await newUserRepository(db, "user");

    // initializing usecases
    const jsonWebTokenManager = await newJsonWebTokenManager();
    const registerUserUseCase = await newRegisterUserUseCase(userRepository, jsonWebTokenManager);
    const loginUserUseCase = await newLoginUserUseCase(userRepository, jsonWebTokenManager);

    // initializing controllers
    const registerUserController = await newRegisterUserController(registerUserUseCase);
    const loginUserController = await newLoginUserController(loginUserUseCase);

    //initialize routers
    const authRouter = Router();
    authRouter.post("/register", registerUserController.registerUser);
    authRouter.post("/login", loginUserController.LoginUser);

    app.use("/api/v1/auth", authRouter);

})();

export default app;

