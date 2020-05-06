import express, { Router } from "express";
import bodyParser, { json } from "body-parser";
import morgan from "morgan";

import { initializeDBConnection } from "../../infrastucture/database/mongo";
import { newJsonWebTokenManager } from "../../infrastucture/json-web-token/json-web-token";
import { newUserRepository } from "../../infrastucture/repository/userRepository";
import { newRegisterUserUseCase } from "../../usecase/registerUser/registerUserUseCase";
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

    // initializing controllers
    const registerUserController = await newRegisterUserController(registerUserUseCase);

    //initialize routers
    const authRouter = Router();
    authRouter.post("/auth/register", registerUserController.registerUser);
    auth

    app.use("/api/v1", authRouter);

})();

export default app;

