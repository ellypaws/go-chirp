import request from 'supertest';
import express from 'express';
import authRoutes from '../src/routes/auth';
import { createUser, getUserByUsername } from '../src/db/queries';
import jwt from 'jsonwebtoken';
import bcrypt from 'bcryptjs';

const app = express();
app.use(express.json());
app.use('/api/v1', authRoutes);

jest.mock('../src/db/queries');
jest.mock('bcryptjs');
jest.mock('jsonwebtoken');

describe('Auth Routes', () => {  describe('POST /signup', () => {
    it('should create a new user successfully', async () => {
      (createUser as jest.Mock).mockResolvedValue({
        id: 1,
        username: 'testuser',
        email: 'testuser@example.com',
        password: 'hashedpassword'
      });

      const response = await request(app)
        .post('/api/v1/signup')
        .send({
          username: 'testuser',
          email: 'testuser@example.com',
          password: 'password123'
        });

      expect(response.status).toBe(201);
      expect(response.body.message).toBe('User created successfully');
    });

    it('should return 500 if there is an error', async () => {
      (createUser as jest.Mock).mockRejectedValue(new Error('Database error'));

      const response = await request(app)
        .post('/api/v1/signup')
        .send({
          username: 'testuser',
          email: 'testuser@example.com',
          password: 'password123'
        });

      expect(response.status).toBe(500);
      expect(response.body.error).toBe('Internal server error');
    });
  });

  describe('POST /login', () => {
    it('should login successfully with valid credentials', async () => {
      (getUserByUsername as jest.Mock).mockResolvedValue({
        id: 1,
        username: 'testuser',
        email: 'testuser@example.com',
        password: 'hashedpassword'
      });
      (bcrypt.compare as jest.Mock).mockResolvedValue(true);
      (jwt.sign as jest.Mock).mockReturnValue('validtoken');

      const response = await request(app)
        .post('/api/v1/login')
        .send({
          username: 'testuser',
          password: 'password123'
        });

      expect(response.status).toBe(200);
      expect(response.body.token).toBe('validtoken');
    });

    it('should return 401 for invalid credentials', async () => {
      (getUserByUsername as jest.Mock).mockResolvedValue(null);

      const response = await request(app)
        .post('/api/v1/login')
        .send({
          username: 'testuser',
          password: 'wrongpassword'
        });

      expect(response.status).toBe(401);
      expect(response.body.error).toBe('Invalid credentials');
    });

    it('should return 500 if there is an error', async () => {
      (getUserByUsername as jest.Mock).mockRejectedValue(new Error('Database error'));

      const response = await request(app)
        .post('/api/v1/login')
        .send({
          username: 'testuser',
          password: 'password123'
        });

      expect(response.status).toBe(500);
      expect(response.body.error).toBe('Internal server error');
    });
  });
});
