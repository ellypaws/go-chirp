'use client'

import { useState } from 'react'
import Header from './components/Header'
import PublicTimeline from './components/PublicTimeline'
import UserHomepage from './components/UserHomepage'
import LoginModal from './components/LoginModal'

export default function Home() {
  const [isLoggedIn, setIsLoggedIn] = useState(false)
  const [user, setUser] = useState(null)
  const [showLoginModal, setShowLoginModal] = useState(false)

  const handleLogin = (data) => {
    localStorage.setItem('token', data.token)
    setUser(data.user)
    setIsLoggedIn(true)
    setShowLoginModal(false)
  }

  const handleLogout = () => {
    setUser(null)
    setIsLoggedIn(false)
  }

  return (
    <div className="container mx-auto px-4">
      <Header 
        isLoggedIn={isLoggedIn} 
        onLoginClick={() => setShowLoginModal(true)}
        onLogout={handleLogout}
        username={user?.username}
      />
      <main className="mt-8">
        {isLoggedIn ? (
          <UserHomepage userId={user.id} />
        ) : (
          <PublicTimeline />
        )}
      </main>
      {showLoginModal && (
        <LoginModal onClose={() => setShowLoginModal(false)} onLogin={handleLogin} />
      )}
    </div>
  )
}

