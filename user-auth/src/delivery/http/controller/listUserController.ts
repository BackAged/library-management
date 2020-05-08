import { Request, Response } from "express";
import { 
    ListUserUseCase, ListUserUseCaseResponse
} from "../../../usecase/listUser/listUserUseCase";
import { object, number } from "@hapi/joi";
import { skipLimitParser } from "../utils/queryParser";

export class ListUserController {
    private listUserUseCase: ListUserUseCase;

    constructor(listUserUseCase: ListUserUseCase) {
        this.listUserUseCase = listUserUseCase;
        this.listUser = this.listUser.bind(this);
    }

    private seralize(response: ListUserUseCaseResponse[]) {
        return response;
    }

    public async listUser(req: Request, res: Response) {
        const schema = object().keys({
            skip: number,
            limit: number,
        });

        const { error } = schema.validate(req.query, { abortEarly: false });
        if (error) {
            return res.status(400).send({
                message: "Please fill up with valid data in all the required fields.",
                errors: error.details
            });
        }

        try {
            const {skip, limit} = skipLimitParser(req.query);
            const response = await this.listUserUseCase.execute(skip, limit);

            return res.status(200).send(this.seralize(response));
        } catch(e) {
            return res.status(500).send({message: "Something went Wrong"});
        }
    }
}

export const newListUserController = (listUserUseCase: ListUserUseCase): ListUserController => {
    return new ListUserController(listUserUseCase);
}