import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata = {
  title: 'Chirp',
  description: 'A simple Twitter clone',
}

export default function RootLayout({
                                     children,
                                   }: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
    <body className={inter.className}>
    <div className="min-h-screen bg-gray-100">
      {children}
    </div>
    </body>
    </html>
  )
}

