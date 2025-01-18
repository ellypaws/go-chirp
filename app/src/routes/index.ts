import express, {type Request, type Response} from 'express';

const router = express.Router();

router.get('/index.html', (req: Request, res: Response) => {
  res.sendFile('index.html', {root: './public'});
});

export default router;