import { Request, Response, NextFunction} from "express";
import { Role } from "../../../entity/user";


export const adminOnly = (req: Request, res: Response, next:  NextFunction): any => {
    const role = req.headers['x-role']
    console.log(req.headers);
    if (role == null || role != Role.Admin) {
        return res.status(401).json({
            message: "Unauthorized",
        });
    } 
    next();
}