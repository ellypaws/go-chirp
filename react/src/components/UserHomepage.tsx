'use client'

import { useState, useEffect } from 'react'
import Tweet from './Tweet'
import CreateTweet from './CreateTweet'

interface TweetData {
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

interface UserHomepageProps {
  userId: number
}

export default function UserHomepage({ userId }: UserHomepageProps) {
  const [tweets, setTweets] = useState<TweetData[]>([])

  useEffect(() => {
    const fetchTweets = async () => {
      try {
        const response = await fetch(`http://localhost:8080/api/v1/user/${userId}/tweets`)
        if (!response.ok) {
          console.error('Error fetching tweets:', response.statusText)
          return
        }
        const data = await response.json()
        setTweets(data.tweets ? [...data.tweets].reverse() : [])
      } catch (error) {
        console.error('Error fetching tweets:', error)
      }
    }

    fetchTweets()
  }, [userId])

  const handleNewTweet = (newTweet: TweetData) => {
    setTweets([newTweet, ...tweets])
  }

  return (
    <div className="flex flex-col h-[calc(100vh-4rem)]">
      <div className="flex-none p-4">
        <h2 className="text-xl font-semibold mb-4">Your Homepage</h2>
        <CreateTweet onNewTweet={handleNewTweet} />
      </div>
      <div className="flex-1 overflow-y-auto p-4">
        <div className="space-y-4">
          {tweets.map((tweet) => (
            <Tweet key={tweet.ID} tweet={tweet} />
          ))}
        </div>
      </div>
    </div>
  )
}

