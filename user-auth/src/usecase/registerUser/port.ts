import { User } from "../../entity/user";
import { GetUserByEmail } from "../loginUser/port";

export interface CreateUser {
    createUser(user: User): Promise<User>;
}

export interface RegisterUserContext extends CreateUser, GetUserByEmail {

}

export interface GenerateJSONWebToken {
    generateJSONWebTOken(payload: any): string;
}