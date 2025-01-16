import express, {type Request, type Response} from 'express';
import bcrypt from 'bcryptjs';
import jwt from 'jsonwebtoken';
import {Pool} from 'pg';
import type {User} from '../models/user';
import {createUser, getUserByUsername} from "../db/queries.ts";

const router = express.Router();

router.post('/signup', async (req: Request, res: Response) => {
  const {username, email, password} = req.body;

  try {
    const hashedPassword = await bcrypt.hash(password, 10);
    const newUser: User = {id: 0, username, email, password: hashedPassword};

    const createdUser = await createUser(newUser.username, newUser.email, newUser.password);

    res.status(201).json({message: 'User created successfully'});
  } catch (error) {
    res.status(500).json({error: 'Internal server error'});
  }
});

router.post('/login', async (req: Request, res: Response)=> {
  const {username, password} = req.body;

  try {
    const user = await getUserByUsername(username);
    
    if (!user) {
      res.status(401).json({error: 'Invalid credentials'});
    }
    
    const passwordValid = user?.password && await bcrypt.compare(password, user.password) || false;

    if (!passwordValid) {
      res.status(401).json({error: 'Invalid credentials'});
    }

    const token = jwt.sign({
      userId: user?.id,
      username: user?.username,
      email: user?.email
    }, process.env.JWT_SECRET || 'secret', {expiresIn: '24h'});

    res.status(200).json({token});
  } catch (error) {
    res.status(500).json({error: 'Internal server error'});
  }
});

export default router;
