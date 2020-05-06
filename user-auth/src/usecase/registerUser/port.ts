import { User } from "../../entity/user";

export interface CreateUser {
    createUser(user: User): Promise<User>;
}

export interface GenerateJSONWebToken {
    generateJSONWebTOken(payload: any): string;
}