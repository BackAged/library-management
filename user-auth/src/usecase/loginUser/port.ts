import { User } from "../../entity/user";

export interface GetUserByEmail {
    getUserByEmail(email: string): Promise<User | null>;
}