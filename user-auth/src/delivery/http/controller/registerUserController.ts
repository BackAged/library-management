import { Request, Response } from "express";
import { object, string, number} from "@hapi/joi";
import { 
    RegisterUserUseCase, RegisterUserUseCaseRequest, RegisterUserUseCaseResponse 
} from "../../../usecase/registerUser/registerUserUseCase";
import { Gender } from "../../../entity/user";

export class RegisterUserController {
    private registerUserUseCase: RegisterUserUseCase;

    constructor(registerUserUseCase: RegisterUserUseCase) {
        this.registerUserUseCase = registerUserUseCase;
        this.registerUser = this.registerUser.bind(this);
    }

    private seralize(response: RegisterUserUseCaseResponse) {
        return response;
    }

    public async registerUser(req: Request, res: Response) {
        const schema = object().keys({
            name: string().required(),
            age: number().positive().required(),
            gender: string().valid(Gender.Male, Gender.Female).required(),
            email: string().required(),
            phone: string(),
            password: string().required(),
        });

        const { error } = schema.validate(req.body, { abortEarly: false });
        if (error) {
            return res.status(400).send({message: "Please fill up with valid data in all the required fields."});
        }

        try {
            const userData: RegisterUserUseCaseRequest = {
                name: req.body.name,
                age: req.body.age,
                gender: req.body.gender,
                email: req.body.email,
                phone: req.body.phone,
                password: req.body.password,
            }

            const response = await this.registerUserUseCase.execute(userData);

            return res.status(200).send(this.seralize(response));
        } catch(e) {
            console.log(e)
            return res.status(500).send({message: "Something went Wrong"});
        }
       
    }
}

export const newRegisterUserController = (registerUserUseCase: RegisterUserUseCase) => {
    return new RegisterUserController(registerUserUseCase);
}