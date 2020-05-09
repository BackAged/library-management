import configuration from "../configuration";

const BaseURL = configuration.USER_AUTH_URL;

export const USER_AUTH_URLS = {
    login : `${BaseURL}/api/v1/auth/login`,
    register: `${BaseURL}/api/v1/auth/register`,
    verifyToken: `${BaseURL}/api/v1/auth/verify-token`,
    listUser: `${BaseURL}/api/v1/user`,
    getUser: (param: string) => `${BaseURL}/api/v1/user/${param}`,
    uploadPorfilePic: `${BaseURL}/api/v1/user/upload-profile-pic`,
}