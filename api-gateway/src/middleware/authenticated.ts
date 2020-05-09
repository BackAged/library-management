import { Request, Response, NextFunction} from "express";


export const isAuthenticated = (req: Request, res: Response, next:  NextFunction): any => {
    const authHeader = req.headers['authorization']
    const token = authHeader && authHeader.split(' ')[1];
    console.log(authHeader);
    if (token == null) {
        return res.status(401).json({
            message: "authorized token not found",
        });
    } 
    next();
}