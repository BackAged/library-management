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
import { newGetUserController } from "./controller/getUserController";
import { newGetUserUseCase } from "../../usecase/getUser/getUserUseCase";
import { newListUserUseCase } from "../../usecase/listUser/listUserUseCase";
import { newListUserController } from "./controller/listUserController";


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
    const getUserUseCase = await newGetUserUseCase(userRepository);
    const listUserUseCase = await newListUserUseCase(userRepository);

    // initializing controllers
    const registerUserController = await newRegisterUserController(registerUserUseCase);
    const loginUserController = await newLoginUserController(loginUserUseCase);

    const getUserController = await newGetUserController(getUserUseCase);
    const listUserController = await newListUserController(listUserUseCase);

    //initialize routers
    const authRouter = Router();
    authRouter.post("/register", registerUserController.registerUser);
    authRouter.post("/login", loginUserController.loginUser);
    app.use("/api/v1/auth", authRouter);

    const userRouter = Router();
    userRouter.get("/:user_id", getUserController.getUser);
    userRouter.get("/", listUserController.listUser);
    app.use("/api/v1/user", userRouter);

})();

export default app;

