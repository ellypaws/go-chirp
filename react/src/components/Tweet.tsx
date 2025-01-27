import { formatDistanceToNow } from 'date-fns'

interface TweetProps {
  tweet: {
    ID: number
    user_id: number
    content: string
    created_at?: string
    CreatedAt: string
  }
}

export default function Tweet({ tweet }: TweetProps) {
  return (
    <div className="bg-white p-4 rounded-lg shadow">
      <p className="text-gray-800">{tweet.content}</p>
      <div className="mt-2 text-sm text-gray-500">
        User ID: {tweet.user_id} • {formatDistanceToNow(new Date(tweet.created_at || tweet.CreatedAt))} ago
      </div>
    </div>
  )
}

