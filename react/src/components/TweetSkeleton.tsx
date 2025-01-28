import { Skeleton } from "./ui/skeleton"

export default function TweetSkeleton() {
  return (
    <div className="bg-[#16181c] p-4 relative border border-transparent rounded-xl">
      <div className="flex items-start justify-between">
        <div className="flex items-center space-x-3">
          <Skeleton className="w-10 h-10 rounded-full" />
          <div className="flex items-center space-x-2">
            <Skeleton className="h-4 w-24" />
            <span className="text-sm text-gray-500">·</span>
            <Skeleton className="h-4 w-20" />
          </div>
        </div>
      </div>
      
      <div className="mt-3">
        <Skeleton className="h-4 w-full" />
        <Skeleton className="h-4 w-3/4 mt-2" />
      </div>
      
      <div className="mt-4 flex items-center space-x-6">
        <div className="flex items-center space-x-2">
          <Skeleton className="h-8 w-8 rounded-full" />
          <Skeleton className="h-4 w-8" />
        </div>
        <div className="flex items-center space-x-2">
          <Skeleton className="h-8 w-8 rounded-full" />
          <Skeleton className="h-4 w-8" />
        </div>
      </div>
    </div>
  )
} 