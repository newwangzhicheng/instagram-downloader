'use client'

import { Input } from '@/components/ui/input'
import { Button } from '@/components/ui/button'
import { useEffect, useState } from 'react'
import axios, { AxiosError } from 'axios'
import { ResourceInfo } from '@/types'
import { Loader2 } from 'lucide-react'
import { useToast } from '@/hooks/use-toast'
import { isValidIgUrl } from '@/lib/utils'
import { usePathname, useSearchParams, useRouter } from 'next/navigation'
import { POST_URL_PARAMS } from '@/lib/constant'

export default function IgForm({
  onGetData
}: {
  onGetData?: (res: ResourceInfo[]) => void
}) {
  const [postUrl, setPostUrl] = useState('')
  const [loading, setLoading] = useState(false)
  const searchParams = useSearchParams()
  const pathname = usePathname()
  const router = useRouter()
  const { toast } = useToast()

  useEffect(() => {
    const url = searchParams.get(POST_URL_PARAMS)
    if (url) {
      const decoded = decodeURIComponent(url)
      setPostUrl(decoded)
      getIgInfo(decoded)
    }
  }, [])

  const updateSearchParams = (url: string) => {
    const newSearchParams = new URLSearchParams(searchParams.toString())
    newSearchParams.set(POST_URL_PARAMS, encodeURIComponent(url))
    router.push(`${pathname}?${newSearchParams.toString()}`)
  }

  const getIgInfo = async (url: string = postUrl) => {
    try {
      setLoading(true)
      if (!isValidIgUrl(url)) {
        toast({
          title: 'Error',
          description: 'Not a valid Instagram link',
          duration: 1500
        })
        return
      }
      const res = await axios({
        url: `/api/ig?${POST_URL_PARAMS}=${encodeURIComponent(url)}`,
        method: 'get'
      })
      if (res.status !== 200) {
        toast({
          title: 'Error',
          description: res.data.message,
          duration: 1500
        })
        return
      }
      updateSearchParams(url)
      onGetData?.(res.data.data)
    } catch (e) {
      console.log(e)
      toast({
        title: 'Error',
        description:
          ((e as AxiosError)?.response?.data as any)?.message ??
          (e as AxiosError).message,
        duration: 1500
      })
    } finally {
      setLoading(false)
    }
  }

  const onClear = () => {
    setPostUrl('')
  }

  const onPaste = async () => {
    try {
      const text = await navigator.clipboard.readText()
      setPostUrl(text)
    } catch {
      toast({
        title: 'Error',
        description: 'Please allow clipboard access.',
        duration: 1500
      })
    }
  }
  return (
    <>
      <div className="flex items-center space-x-2">
        <Input
          type="text"
          placeholder="Paste Instagram Link here"
          value={postUrl}
          onChange={(e) => setPostUrl(e.target.value)}
        />
        {postUrl && <Button onClick={onClear}>Clear</Button>}
        {!postUrl && <Button onClick={onPaste}>Paste</Button>}
      </div>
      <Button className="mt-4" onClick={() => getIgInfo()} disabled={loading}>
        {loading && <Loader2 className="animate-spin" />}
        Download
      </Button>
    </>
  )
}
