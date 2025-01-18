import request from 'supertest';
import express from 'express';
import tweetRoutes from '../src/routes/tweet';
import { jwtMiddleware } from '../src/middleware/auth';
import { createTweet } from '../src/db/queries';
import type {Request, Response} from "express";
import { Tweet } from "../src/models/tweet";

jest.mock('../src/db/queries');
jest.mock('../src/middleware/auth', () => ({
  jwtMiddleware: (req: Request, res: Response, next: () => void) => {
    if (!req.headers.authorization) {
      next();
    }
    req.user = {
      userId: 1,
      username: 'testuser',
      email: 'testuser@example.com',
      expiresAt: Math.floor(Date.now() / 1000) + 3600,
    };
    next();
  },
}));

const app = express();
app.use(express.json());
app.use('/api/v1', jwtMiddleware, tweetRoutes);

describe('Tweet Routes', () => {
  let token: string;

  beforeAll(() => {
    token = 'Bearer valid-token';
  });

  afterEach(() => {
    jest.clearAllMocks();
  });

  describe('POST /tweet', () => {
    it('should create a new tweet', async () => {
      const mockTweet: Tweet = {
        userId: 1,
        content: 'This is a test tweet',
      };

      (createTweet as jest.Mock).mockResolvedValue(mockTweet);

      const response = await request(app)
        .post('/api/v1/tweet')
        .set('Authorization', token)
        .send({ content: 'This is a test tweet' });

      expect(response.status).toBe(201);
      expect(response.body).toEqual(mockTweet);
      expect(createTweet).toHaveBeenCalledWith({
        userId: 1,
        content: 'This is a test tweet',
      });
    });

    it('should return 401 if no token is provided', async () => {
      const response = await request(app)
        .post('/api/v1/tweet')
        .send({ content: 'This is a test tweet' });

      expect(response.status).toBe(401);
      expect(response.body).toEqual({ error: 'Unauthorized' });
    });

    it('should return 500 if there is an internal server error', async () => {
      (createTweet as jest.Mock).mockRejectedValue(new Error('Internal server error'));

      const response = await request(app)
        .post('/api/v1/tweet')
        .set('Authorization', token)
        .send({ content: 'This is a test tweet' });

      expect(response.status).toBe(500);
      expect(response.body).toEqual({ error: 'Internal server error' });
    });
  });
});
