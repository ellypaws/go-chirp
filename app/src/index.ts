import express from 'express';
import dotenv from 'dotenv';
import authRoutes from './routes/auth';
import tweetRoutes from './routes/tweet';
// import followRoutes from './routes/follow';
import {jwtMiddleware} from './middleware/auth';

dotenv.config();

const app = express();
const port = process.env.PORT || 8080;

app.use(express.json());
app.use('/auth', authRoutes);
app.use('/tweet', jwtMiddleware, tweetRoutes);
// app.use('/follow', jwtMiddleware, followRoutes);

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
