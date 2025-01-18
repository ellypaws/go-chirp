import {validateRequest} from "zod-express-middleware";
import {z} from "zod";


export interface Tweet {
  id?: number;
  userId: number;
  content: string;
  createdAt?: string;
}

export const TweetSchema = {
  params: z.object({
    id: z.number()
  }),
  body: z.object({
    content: z.string().min(1)
  })
}

export const validateTweet = validateRequest(TweetSchema);