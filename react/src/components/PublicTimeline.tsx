'use client'

import { useState, useEffect } from 'react'
import Tweet from './Tweet'

interface TweetData {
  id: number
  user_id: number
  content: string
  created_at: string
}

export default function PublicTimeline() {
  const [tweets, setTweets] = useState<TweetData[]>([])

  useEffect(() => {
    const fetchTweets = async () => {
      try {
        const response = await fetch('http://localhost:8080/api/v1/tweets')
        const data = await response.json()
        setTweets(data)
      } catch (error) {
        console.error('Error fetching tweets:', error)
      }
    }

    fetchTweets()
  }, [])

  return (
    <div>
      <h2 className="text-xl font-semibold mb-4">Public Timeline</h2>
      <div className="space-y-4">
        {tweets.map((tweet) => (
          <Tweet key={tweet.id} tweet={tweet} />
        ))}
      </div>
    </div>
  )
}

