import { CreateUser, RegisterUserContext } from "../../usecase/registerUser/port";
import { GetUser } from "../../usecase/getUser/getUser";
import { DeleteUser } from "../../usecase/deleteUser/deleteUser";
import { User } from "../../entity/user";
import { DBInterface } from "../database/db";
import { ObjectID } from "mongodb";

export class UserRepository implements RegisterUserContext, GetUser, DeleteUser {
    private db: DBInterface;
    private collection: string;
    
    public constructor(db: DBInterface, collectionName: string) {
        this.db = db;
        this.collection = collectionName;
    }

    private toPersistence(user: User) {
        //TODO:-> datastorage format data validation could be done here
        return {
            name: user.name,
            age: user.age,
            gender: String(user.gender),
            email: user.email,
            phone: user.phone,
            password: user.password,
        }
    }

    private toModel(user: any) {
        return User.FromUser(user);
    }

    public async createUser(user: User): Promise<User> {
        const ID = await this.db.create(this.collection, this.toPersistence(user));
        user.ID = ID as unknown as string;
        return user;
    }

    public async getUser(userID: string): Promise<User | null> {
        const dbUser = await this.db.findOne(this.collection, {_id: new ObjectID(userID)});
        return this.toModel(dbUser);
    }

    public async getUserByEmail(email: string): Promise<User | null> {
        const dbUser =  await this.db.findOne(this.collection, {email});
        return this.toModel(dbUser);
    }

    public async deleteUser(userID: string): Promise<void> {
        return await this.db.deleteOne(this.collection, {_id: new ObjectID(userID)});
    }
    
}

export const newUserRepository = (db: DBInterface, collectionName: string) => {
    return new UserRepository(db, collectionName);
}