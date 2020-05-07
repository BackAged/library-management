import { GenerateJSONWebToken } from "../../usecase/registerUser/port";
import jsonwebtoken from "jsonwebtoken";

// hard coded for now
const jwtKey = "blahblah";
const jwtExpirySeconds = 3000;

export class JsonWebTokenManager implements GenerateJSONWebToken {
    generateJSONWebTOken(payload: any): string {
        return jsonwebtoken.sign(payload, jwtKey, {
            algorithm: "HS256",
		    expiresIn: jwtExpirySeconds,
        })
    }
}


export const newJsonWebTokenManager = () => {
    return new JsonWebTokenManager();
}