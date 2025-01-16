import express, {type Request, type Response} from 'express';
import {createTweet} from "../db/queries.ts";
import {type Claims, jwtMiddleware} from "../middleware/auth.ts";

const router = express.Router();

router.post('/new', jwtMiddleware, async (req: Request, res: Response) => {
  const user: Claims | undefined = req.user;

  if (!user) {
    res.status(401).json({error: 'Unauthorized'});
    return;
  }

  try {
    const tweet = await createTweet({
      id: 0,
      userId: user.userId,
      content: req.body.content,
      createdAt: new Date().toISOString()
    });
    res.status(201).json(tweet);
  } catch (error) {
    res.status(500).json({error: 'Internal server error'});
  }
});

export default router;