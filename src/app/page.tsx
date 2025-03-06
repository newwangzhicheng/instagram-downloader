import Form from '@/app/components/Form'
import { Github } from 'lucide-react'
import { Toaster } from '@/components/ui/toaster'
import { Suspense } from 'react'
import TestDownload from '@/app/components/TestDownload'
export default function Home() {
  return (
    <div>
      <header className="shadow-md shadow-rose-700">
        <ul className="w-full flex justify-between py-4 px-8 bg-dark ">
          <li>
            <h1 className="font-bold">Insta J</h1>
          </li>
          <li>
            <a
              href="https://github.com/newwangzhicheng/instagram-downloader"
              target="_blank"
            >
              <Github className="text-white" />
            </a>
          </li>
        </ul>
      </header>
      <main className="flex flex-col p-8">
        <Suspense>
          <Form />
        </Suspense>
        <TestDownload />
      </main>
      <Toaster />
    </div>
  )
}
