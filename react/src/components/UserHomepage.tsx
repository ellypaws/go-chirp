'use client'

import { useState, useEffect } from 'react'
import Tweet from './Tweet'
import CreateTweet from './CreateTweet'

interface TweetData {
  id: number
  user_id: number
  content: string
  created_at: string
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
        const data = await response.json()
        setTweets(data)
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
    <div>
      <h2 className="text-xl font-semibold mb-4">Your Homepage</h2>
      <CreateTweet onNewTweet={handleNewTweet} />
      <div className="space-y-4 mt-8">
        {tweets.map((tweet) => (
          <Tweet key={tweet.id} tweet={tweet} />
        ))}
      </div>
    </div>
  )
}

