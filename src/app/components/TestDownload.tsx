'use client'

import { Button } from "@/components/ui/button";
export default function TestDownload() {
  const handleDownload = () => {
    const link = document.createElement('a')
    const blob = new Blob(['Hello, world!'], { type: 'text/plain' })
    link.href = URL.createObjectURL(blob)
    link.download = 'test.txt'
    link.click()
    document.body.removeChild(link)
  }
  return <Button onClick={handleDownload}>Test Download</Button>
}
