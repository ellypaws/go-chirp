'use client'

import { useState, useEffect } from 'react'
import { BrowserRouter as Router, Routes, Route, Navigate, useParams } from 'react-router-dom'
import Header from './components/Header'
import PublicTimeline from './components/PublicTimeline'
import UserHomepage from './components/UserHomepage'
import LoginModal from './components/LoginModal'

interface User {
  ID: number
  username: string
  email: string
}

interface LoginData {
  user: User
  token: string
}

function UserProfile() {
  const { username } = useParams<{ username: string }>()
  return <UserHomepage username={username || ''} />
}

export default function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [user, setUser] = useState<User | null>(null)
  const [showLoginModal, setShowLoginModal] = useState(false)

  useEffect(() => {
    const token = localStorage.getItem('token')
    if (token) {
      // Verify token with backend
      fetch('http://localhost:8080/api/v1/verify', {
        method: 'POST',
        headers: {
          'Authorization': `Bearer ${token}`,
          'Content-Type': 'application/json'
        }
      })
      .then(response => {
        if (response.ok) {
          return response.json()
        }
        throw new Error('Invalid token')
      })
      .then((data: LoginData) => {
        setUser(data.user)
        setIsLoggedIn(true)
      })
      .catch(() => {
        localStorage.removeItem('token')
        setUser(null)
        setIsLoggedIn(false)
      })
    }
  }, [])

  const handleLogin = (data: LoginData) => {
    localStorage.setItem('token', data.token)
    setUser(data.user)
    setIsLoggedIn(true)
    setShowLoginModal(false)
  }

  const handleRegister = (data: LoginData) => {
    localStorage.setItem('token', data.token)
    setUser(data.user)
    setIsLoggedIn(true)
    setShowLoginModal(false)
  }

  const handleLogout = () => {
    localStorage.removeItem('token')
    setUser(null)
    setIsLoggedIn(false)
  }

  return (
    <Router>
      <div className="container mx-auto px-4">
        <Header
          isLoggedIn={isLoggedIn}
          user={user}
          onLoginClick={() => setShowLoginModal(true)}
          onLogout={handleLogout}
        />
        
        <main className="mt-8">
          <Routes>
            <Route path="/" element={<PublicTimeline />} />
            <Route path="/user/:username" element={<UserProfile />} />
          </Routes>
        </main>

        {showLoginModal && (
          <LoginModal
            onClose={() => setShowLoginModal(false)}
            onLogin={handleLogin}
            onRegister={handleRegister}
          />
        )}
      </div>
    </Router>
  )
}

