import { Metadata } from 'next';
import * as React from 'react';

import '@/styles/globals.css';
// !STARTERCONF This is for demo purposes, remove @/styles/colors.css import immediately
import '@/styles/colors.css';
import { Button } from 'antd';

import { siteConfig } from '@/constant/config';
import Logo from '~/svg/Logo.svg';
// !STARTERCONF Change these default meta
// !STARTERCONF Look at @/constant/config to change them
export const metadata: Metadata = {
  metadataBase: new URL(siteConfig.url),
  title: {
    default: siteConfig.title,
    template: `%s | ${siteConfig.title}`,
  },
  description: siteConfig.description,
  robots: { index: true, follow: true },
  // !STARTERCONF this is the default favicon, you can generate your own from https://realfavicongenerator.net/
  // ! copy to /favicon folder
  icons: {
    icon: '/favicon/favicon.ico',
    shortcut: '/favicon/favicon-16x16.png',
    apple: '/favicon/apple-touch-icon.png',
  },
  manifest: `/favicon/site.webmanifest`,
  openGraph: {
    url: siteConfig.url,
    title: siteConfig.title,
    description: siteConfig.description,
    siteName: siteConfig.title,
    images: [`${siteConfig.url}/images/og.jpg`],
    type: 'website',
    locale: 'en_US',
  },
  twitter: {
    card: 'summary_large_image',
    title: siteConfig.title,
    description: siteConfig.description,
    images: [`${siteConfig.url}/images/og.jpg`],
    // creator: '@th_clarence',
  },
  // authors: [
  //   {
  //     name: 'Theodorus Clarence',
  //     url: 'https://theodorusclarence.com',
  //   },
  // ],
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html>
      <body className='m-0 p-0'>
        <header className='flex sticky bg-white top-0 z-50 flex-row items-center border border-gray-200 border-solid text-lg justify-between py-4 px-6 lg:px-48'>
          <div className='flex flex-row items-center border-1 space-x-2'>
            <div className='font-semibold'>Keygate</div>
          </div>
          <nav className='flex flex-row text-sm font-semibold'>
            <div className='flex flex-row space-x-4 items-center'>
              <Button className='text-[#6B7589]' type='text'>
                Contact
              </Button>
              <Button type='text'>Sign up</Button>
            </div>
          </nav>
        </header>
        {children}
      </body>
    </html>
  );
}
