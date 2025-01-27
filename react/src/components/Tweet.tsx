import { formatDistanceToNow } from 'date-fns'

interface TweetProps {
  tweet: {
    ID: number
    user_id: number
    content: string
    created_at?: string
    CreatedAt: string
    like_count: number
    replies_count: number
    parent_id?: number
    username?: string
  }
}

export default function Tweet({ tweet }: TweetProps) {
  return (
    <div className="bg-white p-6 rounded-lg shadow-sm border border-gray-100 hover:border-blue-100 transition-colors">
      <div className="flex items-start justify-between">
        <div className="flex items-center space-x-3">
          <div className="w-10 h-10 rounded-full bg-blue-100 flex items-center justify-center">
            <span className="text-blue-600 font-medium">
              {tweet.username ? tweet.username[0].toUpperCase() : 'U'}
            </span>
          </div>
          <div>
            <p className="font-medium text-gray-900">
              {tweet.username || `User ${tweet.user_id}`}
            </p>
            <p className="text-sm text-gray-500">
              {formatDistanceToNow(new Date(tweet.created_at || tweet.CreatedAt))} ago
            </p>
          </div>
        </div>
        {tweet.parent_id && (
          <div className="text-sm text-gray-500">
            Reply to Tweet #{tweet.parent_id}
          </div>
        )}
      </div>
      
      <p className="mt-4 text-gray-800 whitespace-pre-wrap">{tweet.content}</p>
      
      <div className="mt-4 flex items-center space-x-6 text-sm text-gray-500">
        <div className="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
          </svg>
          <span>{tweet.like_count} likes</span>
        </div>
        <div className="flex items-center space-x-2">
          <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
          </svg>
          <span>{tweet.replies_count} replies</span>
        </div>
      </div>
    </div>
  )
}

