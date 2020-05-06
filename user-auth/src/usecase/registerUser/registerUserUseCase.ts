import { GenerateJSONWebToken, RegisterUserContext } from "./port";
import { User } from "../../entity/user";
import { EmailAlreadyExist } from "./errors";

export interface RegisterUserUseCaseRequest {
    name: string,
    age: number,
    gender: string,
    email: string,
    phone?: string,
    password: string,
}

export interface RegisterUserUseCaseResponse {
    user: {
        name: string,
        age: number,
        gender: string,
        email: string,
        phone?: string,
    },
    token : string,
}

export class RegisterUserUseCase {
    private userRepo: RegisterUserContext;
    private jsonWebTokenGenerator: GenerateJSONWebToken;

    constructor(userRepo: RegisterUserContext, jsonWebTokenGenerator: GenerateJSONWebToken) {
        this.userRepo = userRepo;
        this.jsonWebTokenGenerator = jsonWebTokenGenerator;
    }

    private toCreateUserUseCaseResponse(user: User, token: string): RegisterUserUseCaseResponse {
        return {
            user: {
                name: user.name,
                age: user.age,
                gender: String(user.gender),
                email: user.email,
                phone: user.phone,
            },
            token,
        }
    }

    public async execute(userData: RegisterUserUseCaseRequest) {
        const existingUser = await this.userRepo.getUserByEmail(userData.email);
        if (existingUser) {
            console.log(existingUser);
            throw new EmailAlreadyExist("This email already exists");
        }

        const user = await this.userRepo.createUser(
            User.NewUser(userData),
        );

        const token = await this.jsonWebTokenGenerator.generateJSONWebTOken({id: user.ID, email: user.email});

        return this.toCreateUserUseCaseResponse(user, token);
    }
}

export const newRegisterUserUseCase = (userRepo: RegisterUserContext, jsonWebTokenGenerator: GenerateJSONWebToken) => {
    return new RegisterUserUseCase(userRepo, jsonWebTokenGenerator);
}