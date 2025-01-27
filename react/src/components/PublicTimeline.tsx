'use client'

import { useState, useEffect } from 'react'
import Tweet from './Tweet'

interface TweetData {
  ID: number
  user_id: number
  content: string
  created_at?: string
  CreatedAt: string
}

export default function PublicTimeline() {
  const [tweets, setTweets] = useState<TweetData[]>([])

  useEffect(() => {
    const fetchTweets = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/v1/tweets')
        const data = await response.json()
        setTweets(data.reverse())
      } catch (error) {
        console.error('Error fetching tweets:', error)
      }
    }

    fetchTweets()
  }, [])

  return (
    <div className="flex flex-col h-[calc(100vh-4rem)]">
      <div className="flex-none p-4">
        <h2 className="text-xl font-semibold mb-4">Public Timeline</h2>
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

