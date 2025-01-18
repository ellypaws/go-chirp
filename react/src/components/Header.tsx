import { Button } from "@/components/ui/button"

interface HeaderProps {
  isLoggedIn: boolean
  onLoginClick: () => void
  onLogout: () => void
  username?: string
}

export default function Header({ isLoggedIn, onLoginClick, onLogout, username }: HeaderProps) {
  return (
    <header className="py-4 border-b">
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-blue-500">Chirp</h1>
        <div>
          {isLoggedIn ? (
            <div className="flex items-center gap-4">
              <span>Welcome, {username}!</span>
              <Button onClick={onLogout} variant="outline">Logout</Button>
            </div>
          ) : (
            <Button onClick={onLoginClick}>Login</Button>
          )}
        </div>
      </div>
    </header>
  )
}

