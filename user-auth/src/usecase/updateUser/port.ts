import { User } from "../../entity/user";

export interface UpdateUser {
    getUser(userID: string): Promise<User | null>;
    updateUser(userID: string, user: User): Promise<void>;
}