import { Request, Response, NextFunction} from "express";
import { Role } from "../../../entity/user";


export const adminOnly = (req: Request, res: Response, next:  NextFunction): any => {
    const role = req.headers['x-role']
    if (role == null || role != Role.Admin) {
        return res.status(401).json({
            message: "Unauthorized",
        });
    } 
    next();
}

export const authenticated = (req: Request, res: Response, next:  NextFunction): any => {
    const userID = req.headers['x-userid']
    if (userID == null) {
        return res.status(401).json({
            message: "Unauthorized",
        });
    } 
    next();
}