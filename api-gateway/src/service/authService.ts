import axios from "axios";
import { USER_AUTH_URLS } from "../service-proxy/user-auth";

export const getUserFromToken = async (token: string) => {
    try {
        const response = await axios.get(USER_AUTH_URLS.verifyToken, {
            headers:{
                "authorization": token,
            }
        });

        return response.data;
    } catch (e) {
        // TODO:-> Custom application side error wrapper
        throw e;
    }   
}