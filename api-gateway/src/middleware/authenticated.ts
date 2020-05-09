import { Request, Response, NextFunction} from "express";
import { getUserFromToken } from "../service/authService";


export const Authenticated = async(req: Request, res: Response, next:  NextFunction): Promise<any> => {
    const authHeader = req.headers['authorization']
    const token = authHeader && authHeader.split(' ')[1];
    
    if (!token) {
        return res.status(401).json({
            message: "authorized token not found",
        });
    }

    try {
        const user = await getUserFromToken(authHeader as string);
        console.log(user)
        req.headers["X-USERID"] = user.id;
    } catch(e) {
        console.log(e);
        return res.status(401).send({
            message: e.message ? e.message : "unauthorized",
        });
    }

    next();
}