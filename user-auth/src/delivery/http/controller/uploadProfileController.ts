import { Request, Response } from "express";
import { object, string } from "@hapi/joi";
import multer from "multer";
import { 
    UpdateUserUseCase, UpdateUserUseCaseResponse
} from "../../../usecase/updateUser/updateUserUseCase";
import { UserNotFound } from "../../../usecase/updateUser/errors";

export class UploadProfilePicController {
    private updateUserUseCase: UpdateUserUseCase;

    constructor(updateUserUseCase: UpdateUserUseCase) {
        this.updateUserUseCase = updateUserUseCase;
        this.uploadProfilePic = this.uploadProfilePic.bind(this);
    }

    private seralize(response: UpdateUserUseCaseResponse) {
        return {data: response};
    }

    public async uploadProfilePic(req: Request, res: Response) {
        try {
            const userID = req.headers["x-userid"];
            const profilePic = req.file.path

            const response = await this.updateUserUseCase.execute(userID as string, profilePic);
            
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

export const newUploadProfilePicController = (updateUserUseCase: UpdateUserUseCase): 
UploadProfilePicController => {
    return new UploadProfilePicController(updateUserUseCase);
}