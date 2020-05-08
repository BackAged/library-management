import { Request, Response } from "express";
import { object, string, number} from "@hapi/joi";
import { 
    LoginUserUseCase, LoginUserUseCaseRequest, LoginUserUseCaseResponse 
} from "../../../usecase/loginUser/loginUserUseCase";
import { Gender } from "../../../entity/user";
import { PasswardMisMatch, UserNotFound } from "../../../usecase/loginUser/errors";

export class LoginUserController {
    private loginUserUseCase: LoginUserUseCase;

    constructor(loginUserUseCase: LoginUserUseCase) {
        this.loginUserUseCase = loginUserUseCase;
        this.loginUser = this.loginUser.bind(this);
    }

    private seralize(response: LoginUserUseCaseResponse) {
        return response;
    }

    public async loginUser(req: Request, res: Response) {
        const schema = object().keys({
            email: string().required(),
            password: string().required(),
        });

        const { error } = schema.validate(req.body, { abortEarly: false });
        if (error) {
            return res.status(400).send({
                message: "Please fill up with valid data in all the required fields.",
                errors: error.details
            });
        }

        try {
            const loginCred: LoginUserUseCaseRequest = {
                email: req.body.email,
                password: req.body.password,
            }

            const response = await this.loginUserUseCase.execute(loginCred);

            return res.status(200).send(this.seralize(response));
        } catch(e) {
            console.log(e)
            if (e instanceof UserNotFound) {
                return res.status(400).send({message: e.message});
            }
            if (e instanceof PasswardMisMatch) {
                return res.status(400).send({message: e.message});
            }

            return res.status(500).send({message: "Something went Wrong"});
        }
       
    }
}

export const newLoginUserController = (loginUserUseCase: LoginUserUseCase) => {
    return new LoginUserController(loginUserUseCase);
}