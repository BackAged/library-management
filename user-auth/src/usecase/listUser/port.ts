import { User } from "../../entity/user";

export interface ListUser {
    listUser(skip: number, limit: number): Promise<User[]>;
}