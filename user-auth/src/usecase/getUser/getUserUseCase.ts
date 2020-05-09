import { User } from "../../entity/user";
import { GetUser } from "./port";
import { UserNotFound } from "./errors";

export interface GetUserUseCaseResponse {
    name: string,
    age: number,
    gender: string,
    email: string,
    phone?: string,
    role: string,
}

export class GetUserUseCase {
    private userRepo: GetUser

    constructor(getUserRepo: GetUser) {
        this.userRepo = getUserRepo;
    }

    private toGetUserUseCaseResponse(user: User): GetUserUseCaseResponse {
        return {
            name: user.name,
            age: user.age,
            gender: String(user.gender),
            email: user.email,
            phone: user.phone,
            role: user.role,
        }
    }

    public async execute(userID: string) {
        const user = await this.userRepo.getUser(userID);
        if (!user) {
            throw new UserNotFound("No user exist with this id");
        }

        return this.toGetUserUseCaseResponse(user);
    }
}

export const newGetUserUseCase = (getUserRepo: GetUser) => {
    return new GetUserUseCase(getUserRepo);
}