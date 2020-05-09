import { Request, Response } from "express";
import { object, string } from "@hapi/joi";
import { 
    GetUserUseCase, GetUserUseCaseResponse
} from "../../../usecase/getUser/getUserUseCase";
import { UserNotFound } from "../../../usecase/getUser/errors";

export class GetUserController {
    private getUserUseCase: GetUserUseCase;

    constructor(getUserUseCase: GetUserUseCase) {
        this.getUserUseCase = getUserUseCase;
        this.getUser = this.getUser.bind(this);
    }

    private seralize(response: GetUserUseCaseResponse) {
        return {data: response};
    }

    public async getUser(req: Request, res: Response) {
        const schema = object().keys({
            user_id: string().required(),
        });

        const { error } = schema.validate(req.params, { abortEarly: false });
        if (error) {
            return res.status(400).send({
                message: "Please fill up with valid data in all the required fields.",
                errors: error.details
            });
        }

        try {
            const response = await this.getUserUseCase.execute(req.params.user_id);
            
            return res.status(200).send(this.seralize(response));
        } catch(e) {
            console.log(e)
            if (e instanceof UserNotFound) {
                return res.status(400).send({message: e.message});
            }

            return res.status(500).send({message: "Something went Wrong"});
        }
       
    }
}

export const newGetUserController = (getUserUseCase: GetUserUseCase): GetUserController => {
    return new GetUserController(getUserUseCase);
}