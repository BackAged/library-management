import { Request, Response } from "express";
import { object, string } from "@hapi/joi";
import { 
    GetUserUseCase, GetUserUseCaseResponse
} from "../../../usecase/getUser/getUserUseCase";
import { UserNotFound } from "../../../usecase/getUser/errors";
import { 
    JsonWebTokenManager 
} from "../../../infrastucture/json-web-token/json-web-token";
import { TokenExpiredError, JsonWebTokenError, NotBeforeError } from "jsonwebtoken";

export class VerifyTokenController {
    private getUserUseCase: GetUserUseCase;
    private jsonWebTokenManager: JsonWebTokenManager;

    constructor(getUserUseCase: GetUserUseCase, jsonWebTokenManager: JsonWebTokenManager) {
        this.getUserUseCase = getUserUseCase;
        this.jsonWebTokenManager = jsonWebTokenManager;

        this.verifyToken = this.verifyToken.bind(this);
    }

    private seralize(response: GetUserUseCaseResponse) {
        // TODO:-> if serialization needed
        return response;
    }

    public async verifyToken(req: Request, res: Response) {
        const schema = object().keys({
            authorization: string().required(),
        });

        const { error } = schema.validate(req.headers, { abortEarly: false, allowUnknown: true});
        if (error) {
            return res.status(400).send({
                message: "authorization header is required",
                errors: error.details
            });
        }

        try {
            const authHeader = req.headers['authorization'];
            const token = authHeader && authHeader.split(' ')[1];
            console.log(token);
            const {id, email} = this.jsonWebTokenManager.generatePayloadFromToken(token as string);

            const user = await this.getUserUseCase.execute(id);
            
            return res.status(200).send(this.seralize(user));
        } catch(e) {
            console.log(e)
            if (e instanceof TokenExpiredError || 
                e instanceof JsonWebTokenError ||
                e instanceof NotBeforeError
            ) {
                return res.status(400).send({message: e.message});
            }

            if (e instanceof UserNotFound) {
                return res.status(400).send({message: e.message});
            }

            return res.status(500).send({message: "Something went Wrong"});
        }
       
    }
}

export const newVerifyTokenController = (getUserUseCase: GetUserUseCase, jsonWebTokenManager: JsonWebTokenManager): 
VerifyTokenController => {
    return new VerifyTokenController(getUserUseCase, jsonWebTokenManager);
}