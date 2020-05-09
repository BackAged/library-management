import { GenerateJSONWebToken } from "../../usecase/registerUser/port";
import jsonwebtoken from "jsonwebtoken";

interface PayLoad {
    id: string;
    email: string;
}

// hard coded for now
const jwtKey = "blahblah";
const jwtExpirySeconds = 3000;

export class JsonWebTokenManager implements GenerateJSONWebToken {
    generateJSONWebTOken(payload: PayLoad): string {
        return jsonwebtoken.sign(payload, jwtKey, {
            algorithm: "HS256",
		    expiresIn: jwtExpirySeconds,
        })
    }

    generatePayloadFromToken(token: string): PayLoad {
        const decoded = jsonwebtoken.verify(token, jwtKey);
        return decoded as PayLoad;
    }
}


export const newJsonWebTokenManager = () => {
    return new JsonWebTokenManager();
}