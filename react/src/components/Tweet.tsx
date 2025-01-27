import { formatDistanceToNow } from 'date-fns'

interface User {
  ID: number
  username: string
  email: string
  picture?: string
}

interface Tweet {
  ID: number
  user_id: number
  user?: User
  content: string
  created_at?: string
  CreatedAt: string
  like_count: number
  replies_count: number
  parent_id?: number
}

interface TweetProps {
  tweet: Tweet
}

export default function Tweet({ tweet }: TweetProps) {
  const username = tweet.user?.username || 'Anonymous'
  
  return (
    <div className="bg-[#16181c] p-4 relative border border-transparent rounded-xl hover:bg-[#1d1f23] transition-colors duration-200 before:absolute before:inset-0 before:rounded-xl before:border before:border-[#2f3336] before:content-[''] before:pointer-events-none before:bg-gradient-to-b before:from-[#ffffff10] before:to-[#ffffff08]">
      <div className="flex items-start justify-between">
        <div className="flex items-center space-x-3">
          <div className="w-10 h-10 rounded-full bg-[#2f3336] flex items-center justify-center">
            <span className="text-gray-200 font-medium">
              {username[0].toUpperCase()}
            </span>
          </div>
          <div className="flex items-center space-x-2">
            <span className="font-bold text-gray-200 hover:text-gray-300">
              {username}
            </span>
            <span className="text-sm text-gray-500">·</span>
            <span className="text-sm text-gray-500">
              {formatDistanceToNow(new Date(tweet.created_at || tweet.CreatedAt))} ago
            </span>
          </div>
        </div>
        {tweet.parent_id && (
          <div className="text-sm text-gray-500">
            Reply to Tweet #{tweet.parent_id}
          </div>
        )}
      </div>
      
      <p className="mt-3 text-gray-200 whitespace-pre-wrap">{tweet.content}</p>
      
      <div className="mt-4 flex items-center space-x-6 text-sm text-gray-500">
        <div className="flex items-center space-x-2 group cursor-pointer">
          <div className="p-2 rounded-full group-hover:bg-[#f9132830] group-hover:text-[#f91328] transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
            </svg>
          </div>
          <span>{tweet.like_count}</span>
        </div>
        <div className="flex items-center space-x-2 group cursor-pointer">
          <div className="p-2 rounded-full group-hover:bg-[#1d9bf030] group-hover:text-[#1d9bf0] transition-colors">
            <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={1.5} d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
            </svg>
          </div>
          <span>{tweet.replies_count}</span>
        </div>
      </div>
    </div>
  )
}

