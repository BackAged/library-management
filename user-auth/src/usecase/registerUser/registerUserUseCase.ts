import { CreateUser, GenerateJSONWebToken } from "./port";
import { User } from "../../entity/user";

// kinda like DTO
export interface RegisterUserUseCaseRequest {
    name: string,
    age: number,
    gender: string,
    email: string,
    phone?: string,
    password: string,
}

// kinda like DTO
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
    private createUserRepo: CreateUser;
    private jsonWebTokenGenerator: GenerateJSONWebToken;

    constructor(createUserRepo: CreateUser, jsonWebTokenGenerator: GenerateJSONWebToken) {
        this.createUserRepo = createUserRepo;
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
        // TODO:-> duplicate email checking
        const user = await this.createUserRepo.createUser(
            User.NewUser(userData),
        );

        const token = await this.jsonWebTokenGenerator.generateJSONWebTOken(user);

        return this.toCreateUserUseCaseResponse(user, token);
    }
}

export const newRegisterUserUseCase = (createUserRepo: CreateUser, jsonWebTokenGenerator: GenerateJSONWebToken) => {
    return new RegisterUserUseCase(createUserRepo, jsonWebTokenGenerator);
}