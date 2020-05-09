import { User } from "../../entity/user";
import { UpdateUser } from "./port";
import { UserNotFound } from "./errors";

export interface UpdateUserUseCaseResponse {
    id: string,
    name: string,
    age: number,
    gender: string,
    email: string,
    phone?: string,
    role: string,
    profilePic?: string,
}

export class UpdateUserUseCase {
    private userRepo: UpdateUser

    constructor(userRepo: UpdateUser) {
        this.userRepo = userRepo;
    }

    private toGetUserUseCaseResponse(user: User): UpdateUserUseCaseResponse {
        return {
            id: user.ID as string,
            name: user.name,
            age: user.age,
            gender: String(user.gender),
            email: user.email,
            phone: user.phone,
            role: user.role,
            profilePic: user.profilePic,
        }
    }

    public async execute(userID: string, profilePic: string) {
        const user = await this.userRepo.getUser(userID);
        if (!user) {
            throw new UserNotFound("No user exist with this id");
        }

        user.profilePic = profilePic;
        await this.userRepo.updateUser(userID, user);
        return this.toGetUserUseCaseResponse(user);
    }
}

export const newUpdateUserUseCase = (userRepo: UpdateUser) => {
    return new UpdateUserUseCase(userRepo);
}