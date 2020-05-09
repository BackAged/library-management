import { Request, Response, NextFunction} from "express";
import { getUserFromToken } from "../service/authService";

export const AdminOnly = async (req: Request, res: Response, next:  NextFunction): Promise<any> => {
    const authHeader = req.headers['authorization']    
    if (!authHeader) {
        return res.status(401).json({
            message: "authorized token not found",
        });
    }

    try {
        const user = await getUserFromToken(authHeader as string);
    
        req.headers["X-ROLE"] = user.role;
    } catch(e) {
        console.log(e);
        return res.status(401).send({
            message: e.message ? e.message : "unauthorized",
        });
    }

    next();
}