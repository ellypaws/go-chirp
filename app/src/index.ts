import express from 'express';
import dotenv from 'dotenv';
import authRoutes from './routes/auth';
import tweetRoutes from './routes/tweet';
// import followRoutes from './routes/follow';
import index from './routes';
import {jwtMiddleware} from './middleware/auth';
import {migrate} from "./db/db.ts";

dotenv.config();

const app = express();
const port = process.env.PORT || 8080;

app.use(express.json());
app.use('/', index)
app.use('/api/v1', authRoutes);
app.use('/api/v1', jwtMiddleware, tweetRoutes);
// app.use('/follow', jwtMiddleware, followRoutes);

migrate().then(r => console.log(r))
  .catch(e => console.error(e))
  .finally(() => console.log('Migration complete'));

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
