import type {Request, Response, NextFunction} from 'express';
import jwt, {type JwtPayload, type Secret } from 'jsonwebtoken';
import crypto from 'crypto';

// Define the JWT secret key
const JWT_SECRET: Secret = process.env.JWT_SECRET || crypto.randomBytes(32).toString('hex');

// Define the Claims type to match the payload structure
export interface Claims extends JwtPayload {
  userId: number;
  username: string;
  email: string;
  expiresAt: number;
}

// Middleware to check for JWT and validate it
export const jwtMiddleware = (req: Request, res: Response, next: NextFunction): void => {
  const authHeader = req.headers.authorization;

  if (!authHeader) {
    res.status(401).json({ error: 'Authorization header missing' });
    return;
  }

  const token = authHeader.startsWith('Bearer ') ? authHeader.slice(7) : authHeader;

  try {
    const claims = jwt.verify(token, JWT_SECRET) as Claims;

    if (jwtExpired(claims)) {
      res.status(401).json({ error: 'Token expired' });
      return;
    }

    // Add claims to the request context
    req.user = claims;
    next();
  } catch (err) {
    res.status(401).json({ error: 'Invalid token', debug: err});
  }
};

// Function to check if the JWT is expired
const jwtExpired = (claims: Claims): boolean => {
  return claims.expiresAt < Math.floor(Date.now() / 1000);
};

// Extend the Express Request object to include the user property
declare global {
  namespace Express {
    interface Request {
      user?: Claims;
    }
  }
}
