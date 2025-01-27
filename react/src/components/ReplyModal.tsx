'use client'

import { Dialog, DialogContent, DialogHeader, DialogTitle } from "@/components/ui/dialog"
import { useState } from 'react'
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"

interface Tweet {
  ID: number
  user_id: number
  content: string
  created_at?: string
  CreatedAt: string
  like_count: number
  replies_count: number
  parent_id?: number
  user?: {
    username: string
  }
}

interface ReplyModalProps {
  isOpen: boolean
  onClose: () => void
  parentTweet: Tweet
  onNewReply: (tweet: any) => void
}

export default function ReplyModal({ isOpen, onClose, parentTweet, onNewReply }: ReplyModalProps) {
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
        body: JSON.stringify({ 
          content: content,
          parent_id: parentTweet.ID
        }),
      })

      if (!response.ok) {
        throw new Error(`Failed to create reply: ${response.statusText}`)
      }

      const newTweet = await response.json()
      onNewReply(newTweet)
      setContent('')
      onClose()
    } catch (error) {
      console.error('Error creating reply:', error)
    }
  }

  return (
    <Dialog open={isOpen} onOpenChange={onClose}>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Reply to @{parentTweet.user?.username}</DialogTitle>
        </DialogHeader>
        <div className="mt-4">
          <div className="text-sm text-gray-500 mb-4">
            Replying to Tweet: {parentTweet.content}
          </div>
          <form onSubmit={handleSubmit} className="space-y-4">
            <Textarea
              placeholder="Tweet your reply"
              value={content}
              onChange={(e) => setContent(e.target.value)}
              rows={3}
              maxLength={280}
              required
            />
            <div className="flex justify-end space-x-2">
              <Button variant="outline" onClick={onClose}>Cancel</Button>
              <Button type="submit">Reply</Button>
            </div>
          </form>
        </div>
      </DialogContent>
    </Dialog>
  )
} 