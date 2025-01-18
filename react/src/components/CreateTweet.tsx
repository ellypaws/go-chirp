'use client'

import { useState } from 'react'
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"

interface CreateTweetProps {
  onNewTweet: (tweet: any) => void
}

export default function CreateTweet({ onNewTweet }: CreateTweetProps) {
  const [content, setContent] = useState('')

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()

    try {
      const response = await fetch('http://localhost:8080/api/v1/tweet', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${localStorage.getItem('token')}`,
        },
        body: JSON.stringify({ content: content }),
      })

      if (!response.ok) {
        throw new Error(`'Failed to create tweet: ${response.statusText}: ${await response.text()}`)
      }

      const newTweet = await response.json()
      onNewTweet(newTweet)
      setContent('')
    } catch (error) {
      console.error('Error creating tweet:', error)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="space-y-4">
      <Textarea
        placeholder="What's happening?"
        value={content}
        onChange={(e) => setContent(e.target.value)}
        rows={3}
        maxLength={280}
        required
      />
      <Button type="submit">Tweet</Button>
    </form>
  )
}

