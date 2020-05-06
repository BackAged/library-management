import { GetUserByEmail } from "./port";
import { User } from "../../entity/user";
import { GenerateJSONWebToken } from "../registerUser/port";
import { UserNotFound, PasswardMisMatch } from "./errors";


export interface LoginUserUseCaseRequest {
    email: string,
    password: string,
}

export interface LoginUserUseCaseResponse {
    user: {
        name: string,
        age: number,
        gender: string,
        email: string,
        phone?: string,
    },
    token : string,
}

export class LoginUserUseCase {
    private userRepo: GetUserByEmail;
    private jsonWebTokenGenerator: GenerateJSONWebToken;

    constructor(userRepo: GetUserByEmail, jsonWebTokenGenerator: GenerateJSONWebToken) {
        this.userRepo = userRepo;
        this.jsonWebTokenGenerator = jsonWebTokenGenerator;
    }

    private toCreateUserUseCaseResponse(user: User, token: string): LoginUserUseCaseResponse {
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

    public async execute(loginCred: LoginUserUseCaseRequest) {
        const user = await this.userRepo.getUserByEmail(loginCred.email);
        if (!user) {
            throw new UserNotFound("No user found with this email");
        }

        if (!user.matchPassword(loginCred.password)) {
            throw new PasswardMisMatch("Password didn't match");
        }

        const token = await this.jsonWebTokenGenerator.generateJSONWebTOken(user);

        return this.toCreateUserUseCaseResponse(user, token);
    }
}

export const newLoginUserUseCase = (userRepo: GetUserByEmail, jsonWebTokenGenerator: GenerateJSONWebToken) => {
    return new LoginUserUseCase(userRepo, jsonWebTokenGenerator);
}