import { User } from "../../entity/user";
import { ListUser } from "./port";

export interface ListUserUseCaseResponse {
    id: string,
    name: string,
    age: number,
    gender: string,
    email: string,
    phone?: string,
}

export class ListUserUseCase {
    private userRepo: ListUser

    constructor(getUserRepo: ListUser) {
        this.userRepo = getUserRepo;
    }

    private toGetUserUseCaseResponse(users: User[]): ListUserUseCaseResponse[] {
        return users.map(user => ({
            id: user.ID as string,
            name: user.name,
            age: user.age,
            gender: String(user.gender),
            email: user.email,
            phone: user.phone,
        }));
    }

    public async execute(skip: number, limit: number) {
        const users = await this.userRepo.listUser(skip, limit);
        return this.toGetUserUseCaseResponse(users);
    }
}

export const newListUserUseCase = (userRepo: ListUser) => {
    return new ListUserUseCase(userRepo);
}