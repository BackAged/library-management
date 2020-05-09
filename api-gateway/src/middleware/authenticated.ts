import { Request, Response, NextFunction} from "express";

export const Authenticated = (req: Request, res: Response, next:  NextFunction): any => {
    const authHeader = req.headers['authorization']
    const token = authHeader && authHeader.split(' ')[1];
    
    if (!token) {
        return res.status(401).json({
            message: "authorized token not found",
        });
    } 

    next();
}