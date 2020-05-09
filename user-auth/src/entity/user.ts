import bcrypt from "bcrypt";

export enum Gender {
    Male = "Male",
    Female = "Female",
}

export enum Role {
    Admin = "Admin",
    Member = "Member"
}

// UserProps holds User domain data
interface UserProps {
    ID?: string;
    name: string;
    age: number;
    gender: string; 
    email:string; 
    phone?: string;
    password: string;
    role?: Role;
    profilePic?: string;
}


export class User {
    private _ID?: string;
    private _name: string;
    private _age: number;
    private _gender: Gender
    private _email: string;
    private _phone?: string;
    private _password: string;
    private _role: Role;
    private _profilePic?: string;

    private constructor(userData: UserProps) {
        this._name = userData.name;
        this._age = userData.age;
        this._gender = userData.gender as Gender;
        this._email = userData.email;
        this._phone = userData.phone;
        this._password = userData.password;
        this._ID = userData.ID;
        this._role = userData.role ? userData.role : Role.Member;
        this._profilePic = userData.profilePic;
    }

    private static isValid(userData: UserProps): string[] {
        const errors: string[]= []

        if (!userData.name || userData.name.length < 2 || userData.name.length > 20){
            errors.push("Invalid Name");
        }

        if (!userData.age || userData.age <= 0 || userData.age > 150){
            errors.push("Invalid age");
        }

        if (userData.gender !== Gender.Male && userData.gender !== Gender.Female) {
            errors.push("Invalid gender");
        }

        if (!userData.email) { //TODO-> email validation
            errors.push("Invalid email");
        }

        if (!userData.password || userData.password.length <= 2 || userData.password.length >= 10) {
            errors.push("Invalid password");
        }

        if (userData.role != Role.Member && userData.role != Role.Admin) {
            errors.push("Invalid Role");
        }
        
        console.log(errors);
        return errors;
    }

    public static FromUser(userData: UserProps): User {
        return new User(userData);
    }

    public static NewUser(userData: UserProps): User {
        const errors = this.isValid(userData);
        if (errors.length !== 0) {
            throw new Error("Invalid User")
        }
    
        userData.password = this.hasPassword(userData.password);

        return new User(userData);
    }

    set ID(ID: string | undefined) {
        this._ID = ID;
    }

    get ID(): string | undefined {
        return this._ID;
    }

    get name(): string {
        return this._name;
    }

    get age(): number {
        return this._age;
    }

    get gender(): Gender {
        return this._gender;
    }

    get email(): string {
        return this._email;
    }

    get phone(): string | undefined {
        return this._phone;
    }

    get password(): string {
        return this._password;
    }

    get role(): Role | string {
        return this._role;
    }

    get profilePic(): string | undefined {
        return this._profilePic;
    }

    set profilePic(profilePic: string | undefined) {
        this._profilePic = profilePic;
    }

    private static hasPassword(plainPassword: string): string{
        return bcrypt.hashSync(plainPassword, 10);
    }

    public async matchPassword(plainPassword: string): Promise<boolean>{
        return await bcrypt.compare(plainPassword, this.password);
    }
}

